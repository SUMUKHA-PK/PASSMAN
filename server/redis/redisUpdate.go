package redis

import (
	"github.com/gomodule/redigo/redis"
)

// Update updates the REDIS db with given data
func Update(authPwd string, vault []byte) error {
	conn, err := redis.Dial("tcp", "172.17.0.1:6379")
	if err != nil {
		return err
	}

	defer conn.Close()

	_, err = conn.Do("HMSET", "VAULT", authPwd, string(vault))
	if err != nil {
		return err
	}
	return nil
}
