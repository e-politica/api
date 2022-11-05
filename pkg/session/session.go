package session

import (
	"context"
	"encoding/json"
	"time"

	"github.com/e-politica/api/config"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

const (
	accessTokenPrefix = "access-token:"
	userIdPrefix      = "user-id:"
)

var client = redis.NewClient(&redis.Options{
	Addr:     config.RedisAddr,
	Password: config.RedisPassword,
	DB:       config.RedisDB,
})

type Session struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	Expiration   time.Time `json:"expiration"`
}

func NewSession(ctx context.Context, userId string) (session Session, err error) {
	expiration := time.Now().Add(config.RedisSessionDurationHour)

	session = Session{
		AccessToken:  uuid.NewString(),
		RefreshToken: uuid.NewString(),
		Expiration:   expiration,
	}

	sessionJson, err := json.Marshal(&session)
	if err != nil {
		return
	}

	err = client.Set(
		ctx,
		userIdPrefix+userId,
		string(sessionJson),
		time.Until(expiration),
	).Err()
	if err != nil {
		return
	}

	err = client.Set(
		ctx,
		accessTokenPrefix+session.AccessToken,
		userId,
		time.Until(expiration),
	).Err()

	return
}

func GetSession(ctx context.Context, userId string) (session Session, err error) {
	rawSession, err := client.Get(
		ctx,
		userIdPrefix+userId,
	).Result()
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(rawSession), &session)
	return
}

func GetUserId(ctx context.Context, accessToken string) (userId string, err error) {
	return client.Get(
		ctx,
		accessTokenPrefix+accessToken,
	).Result()
}
