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

var client = NewRedis()

type Redis struct {
	m map[string]string
}

func NewRedis() *Redis {
	return &Redis{m: make(map[string]string)}
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	value, ok := r.m[key]
	if !ok {
		return "", redis.Nil
	}
	return value, nil
}

func (r *Redis) Set(ctx context.Context, key, value string, expiration time.Duration) error {
	r.m[key] = value
	return nil
}

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
	)
	if err != nil {
		return
	}

	err = client.Set(
		ctx,
		accessTokenPrefix+session.AccessToken,
		userId,
		time.Hour*config.RedisSessionDurationHour,
	)

	return
}

func GetSession(ctx context.Context, userId string) (session Session, err error) {
	rawSession, err := client.Get(
		ctx,
		userIdPrefix+userId,
	)
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
	)

	if err == redis.Nil {
		err = ErrSessionNotFound
	}
	return
}
