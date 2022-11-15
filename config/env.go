package config

import (
	"time"

	"github.com/e-politica/api/pkg/env"
)

var (
	ServerPort       = env.Get("SERVER_PORT", "8080")
	CommentDelayHour = env.Get[time.Duration]("COMMENT_DELAY_HOUR", 0)

	DbHost         = env.Get("DB_HOST", "localhost")
	DbName         = env.Get("DB_NAME", "epolitica")
	DbUser         = env.Get("DB_USER", "postgres")
	DbPassword     = env.Get("DB_PASSWORD", "")
	DbPort         = env.Get("DB_PORT", "5432")
	DbSslMode      = env.Get("DB_SSLMODE", "require")
	DbReconnectSec = env.Get("DB_RECONNECT_SEC", 10)
	DbWaitStart    = env.Get("DB_WAIT_START", true)

	RedisAddr                = env.Get("REDIS_ADDR", "localhost:6379")
	RedisPassword            = env.Get("REDIS_PASSWORD", "")
	RedisDB                  = env.Get("REDIS_DB", 0)
	RedisSessionDurationHour = env.Get[time.Duration]("REDIS_SESSION_DURATION_HOUR", 24)

	GoogleClientSecret = env.Get("GOOGLE_CLIENT_SECRET", "")
)
