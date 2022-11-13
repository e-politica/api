package session

import (
	"context"
	"encoding/json"
	"time"

	"github.com/e-politica/api/config"
	"github.com/e-politica/api/models/v1/user"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

const (
	accessTokenPrefix = "access-token:"
	userIdPrefix      = "user-id:"
	emailPrefix       = "email:"
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

	sessionJson, err := json.Marshal(session)
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

func NewEmailVerification(ctx context.Context, params user.RegisterDefaultParams) (id string, err error) {
	paramsJson, err := json.Marshal(params)
	if err != nil {
		return
	}

	id = uuid.NewString()

	err = client.Set(
		ctx,
		emailPrefix+id,
		string(paramsJson),
		time.Until(time.Now().Add(time.Minute*30)),
	).Err()

	return
}

func GetEmailVerification(ctx context.Context, id string) (params user.RegisterDefaultParams, err error) {
	paramsJson, err := client.Get(
		ctx,
		emailPrefix+id,
	).Result()
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(paramsJson), &params)
	return
}
