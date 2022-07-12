GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64

heroku-deploy:
  go build -o bin/main -v cmd/main.go
 