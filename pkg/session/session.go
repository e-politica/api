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

// Session struct -> uuid:user_id
func NewSession(ctx context.Context, userId string) (token string, err error) {
	token = uuid.NewString()
	err = client.Set(
		ctx,
		token,
		userId,
		time.Until(time.Now().Add(config.RedisSessionDurationHour)),
	).Err()
	return
}

// Session struct -> uuid:user_id
func GetSession(ctx context.Context, token string) (string, error) {
	return client.Get(
		ctx,
		token,
	).Result()
}
