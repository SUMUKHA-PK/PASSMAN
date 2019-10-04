package redis

import "github.com/gomodule/redigo/redis"

// Retrieve gets the data from the REDIS db
func Retrieve(authPwd string) (string, error) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	vault, err := redis.String(conn.Do("HGET", "VAULT", authPwd))
	if err != nil {
		return "", err
	}
	return vault, nil
}
