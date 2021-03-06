package redis

import (
	"github.com/gomodule/redigo/redis"
)

// Retrieve gets the data from the REDIS db
func Retrieve(authPwd string) ([]byte, error) {
	conn, err := redis.Dial("tcp", "172.17.0.1:6379")
	if err != nil {
		return []byte{}, err
	}
	defer conn.Close()

	vault, err := redis.String(conn.Do("HGET", "VAULT", authPwd))
	if err != nil {
		return []byte{}, err
	}
	return []byte(vault), nil
}
