package store

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/abdulloh76/invoice-dashboard/types"
	"github.com/go-redis/redis/v9"
)

type RedisCacheStore struct {
	expires time.Duration
	client  *redis.Client
}

var _ types.InvoiceCacheStore = (*RedisCacheStore)(nil)

func NewRedisCacheStore(address, password string, db int, expires time.Duration) *RedisCacheStore {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     address,
		DB:       db,
		Password: password,
	})

	return &RedisCacheStore{
		client:  redisClient,
		expires: expires,
	}
}

func (r *RedisCacheStore) Set(ctx context.Context, key string, invoice *types.InvoiceModel) {
	json, err := json.Marshal(invoice)
	if err != nil {
		panic(err)
	}

	r.client.Set(ctx, key, json, r.expires*time.Second)
}

func (r *RedisCacheStore) Get(ctx context.Context, key string) *types.InvoiceModel {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return nil
	}

	invoice := types.InvoiceModel{}
	err = json.Unmarshal([]byte(val), &invoice)
	if err != nil {
		panic(err)
	}

	return &invoice
}

func (r *RedisCacheStore) Delete(ctx context.Context, key string) error {
	confirmation, err := r.client.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	if confirmation == 0 {
		return errors.New("smth went wrong could not remove from redis")
	}
	return nil
}
