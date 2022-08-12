package session

import (
	"context"
	"time"

	"github.com/e-politica/api/config"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

var client = redis.NewClient(&redis.Options{
	Addr:     config.RedisAddr,
	Password: config.RedisPassword,
	DB:       config.RedisDB,
})

func NewSession(ctx context.Context, userId string) (token string, err error) {
	token = uuid.NewString()

	err = client.Set(
		ctx,
		token,
		userId,
		time.Until(time.Now().Add(config.RedisSessionDurationHour)),
	).Err()
	if err != nil {
		return
	}

	err = client.Set(
		ctx,
		userId,
		token,
		time.Until(time.Now().Add(config.RedisSessionDurationHour)),
	).Err()

	return
}

func GetSession(ctx context.Context, token string) (string, error) {
	return client.Get(
		ctx,
		token,
	).Result()
}

func GetSessionToken(ctx context.Context, userId string) (string, error) {
	return client.Get(
		ctx,
		userId,
	).Result()
}
