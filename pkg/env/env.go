package env

import (
	"os"
	"strconv"
)

func GetString(k, d string) string {
	if v, found := os.LookupEnv(k); found {
		return v
	}
	return d
}

func GetInt(k string, d int) int {
	if v, found := os.LookupEnv(k); found {
		if x, err := strconv.Atoi(v); err == nil {
			return x
		}
	}
	return d
}

func GetBool(k string, d bool) bool {
	if v, found := os.LookupEnv(k); found {
		if x, err := strconv.ParseBool(v); err == nil {
			return x
		}
	}
	return d
}

func Get[T any](k string, d T) T {
	v, found := os.LookupEnv(k)
	if !found {
		return d
	}

	var y any
	switch x := any(d).(type) {
	case string:
		y = any(v)
	case int:
		y = any(convInt(v, x))
	case bool:
		y = any(convBool(v, x))
	default:
		return d
	}
	return y.(T)
}

func convInt(v string, d int) int {
	if x, err := strconv.Atoi(v); err == nil {
		return x
	}
	return d
}

func convBool(v string, d bool) bool {
	if x, err := strconv.ParseBool(v); err == nil {
		return x
	}
	return d
}
