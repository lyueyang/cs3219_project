package storage

import (
	"fmt"
	"github.com/go-redis/redis"
)

type RedisStorage struct {
	RedisConnect *redis.Client
}

func NewRedisStorage(redis *redis.Client) RedisStorage {
	out := RedisStorage{RedisConnect: redis}

	return out
}
func (s RedisStorage) GetAccounts() []Account {

	keys, err := s.RedisConnect.Do("KEYS", "*").Result()

	if err != nil {
		return nil
	}
	var out []Account
	for _, key := range keys.([]interface{}) {
		name := key.(string)

		desc, _ := s.RedisConnect.Get(name).Result()

		out = append(out, Account{
			Name:        name,
			Description: desc,
		})
	}

	return out
}

func (s RedisStorage) StoreAccount(account *Account) {
	err := s.RedisConnect.Set(account.Name, account.Description, 0).Err()

	if err != nil {
		fmt.Println(err)
	}
}
