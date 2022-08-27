package store

import (
	"context"
	"encoding/json"
	"time"

	"github.com/abdulloh76/invoice-dashboard/types"
	"github.com/go-redis/redis/v9"
)

type RedisCacheStore struct {
	expires time.Duration
	client  *redis.Client
}

func NewRedisCacheStore(address string, db int, expires time.Duration) *RedisCacheStore {
	redisClient := redis.NewClient(&redis.Options{
		Addr: address,
		DB:   db,
	})

	return &RedisCacheStore{
		client:  redisClient,
		expires: expires,
	}
}

func (r *RedisCacheStore) Set(ctx context.Context, key string, post *types.InvoiceModel) {
	json, err := json.Marshal(post)
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

	post := types.InvoiceModel{}
	err = json.Unmarshal([]byte(val), &post)
	if err != nil {
		panic(err)
	}

	return &post
}
