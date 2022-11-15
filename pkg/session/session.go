package session

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/e-politica/api/config"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

const (
	accessTokenPrefix = "access-token:"
	userIdPrefix      = "user-id:"
	emailPrefix       = "email:"
)

var ErrSessionNotFound = errors.New("session not found")

var client = redis.NewClient(&redis.Options{
	Addr:     config.RedisAddr,
	Password: config.RedisPassword,
	DB:       config.RedisDB,
})

type Session struct {
	AccessToken string    `json:"access_token"`
	UserId      string    `json:"user_id"`
	Expiration  time.Time `json:"expiration"`
}

func NewSession(ctx context.Context, userId string) (session Session, err error) {
	expiration := time.Now().Add(time.Hour * config.RedisSessionDurationHour)

	session = Session{
		AccessToken: uuid.NewString(),
		UserId:      userId,
		Expiration:  expiration,
	}

	sessionJson, err := json.Marshal(session)
	if err != nil {
		return
	}

	err = client.Set(
		ctx,
		userIdPrefix+userId,
		string(sessionJson),
		time.Hour*config.RedisSessionDurationHour,
	).Err()
	if err != nil {
		return
	}

	err = client.Set(
		ctx,
		accessTokenPrefix+session.AccessToken,
		userId,
		time.Hour*config.RedisSessionDurationHour,
	).Err()

	return
}

func GetSession(ctx context.Context, userId string) (session Session, err error) {
	rawSession, err := client.Get(
		ctx,
		userIdPrefix+userId,
	).Result()
	if err != nil {
		if err == redis.Nil {
			err = ErrSessionNotFound
		}
		return
	}

	err = json.Unmarshal([]byte(rawSession), &session)
	return
}

func GetUserId(ctx context.Context, accessToken string) (userId string, err error) {
	userId, err = client.Get(
		ctx,
		accessTokenPrefix+accessToken,
	).Result()

	if err == redis.Nil {
		err = ErrSessionNotFound
	}
	return
}
