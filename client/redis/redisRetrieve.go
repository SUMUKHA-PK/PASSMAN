package redis

import (
	"github.com/gomodule/redigo/redis"
)

// VaultData is the data in format as stored in REDIS
type VaultData struct {
	VaultPwd string
	Vault    string
}

// Retrieve gets data from the DB
func Retrieve(username string) (VaultData, error) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		return VaultData{}, err
	}
	defer conn.Close()

	vault, err := redis.String(conn.Do("HGET", username, "vault"))
	if err != nil {
		return VaultData{}, err
	}

	vaultPwd, err := redis.String(conn.Do("HGET", username, "vaultPwd"))
	if err != nil {
		return VaultData{}, err
	}

	return VaultData{vaultPwd, vault}, nil
}
