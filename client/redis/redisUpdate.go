package redis

import (
	"github.com/gomodule/redigo/redis"
)

// Update adds data to the DB
func Update(username string, vaultPwd string, vault string) error {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		return err
	}

	defer conn.Close()

	_, err = conn.Do("HMSET", username, "vault", vault, "vaultPwd", vaultPwd)
	if err != nil {
		return err
	}
	return nil
}
