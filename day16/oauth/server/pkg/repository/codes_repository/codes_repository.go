package codesrepository

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type Codes struct {
	redis redis.Client
}

func NewCodes(addr string) (Codes, error) {
	codes := Codes{
		redis: *redis.NewClient(&redis.Options{
			Network:  "tcp",
			Addr:     addr,
			Password: "",
			DB:       0,
		}),
	}

	_, err := codes.redis.Ping(context.Background()).Result()
	if err != nil {
		return codes, errors.New("Ошибка подключения к Redis")
	}

	return codes, nil
}

func (c Codes) RandomString(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:length]
}

func (c Codes) Add(id int, code string) error {
	str := strconv.Itoa(id)
	return c.redis.Set(context.Background(), str, code, time.Minute*30).Err()
}

func (c Codes) Get(id int) (string, error) {
	str := strconv.Itoa(id)
	return c.redis.Get(context.Background(), str).Result()
}
