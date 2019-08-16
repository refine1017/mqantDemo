package db

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
)



type Player struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}

func (m *Module) ExistPlayer(username string) (bool, error) {
	return redis.Bool(m.redis.Do("HEXISTS", RedisPlayers, username))
}

func (m *Module) CreatePlayer(username string, nickname string) (string, error) {
	id, err := redis.Int(m.redis.Do("INCR", RedisPlayerAutoId))
	if err != nil {
		return "", err
	}

	player := &Player{}
	player.Id = id
	player.Username = username
	player.Nickname = nickname

	buf, err := json.Marshal(player)
	if err != nil {
		return "", err
	}

	playerData := string(buf)

	if _, err := m.redis.Do("HSET", RedisPlayers, username, playerData); err != nil {
		return "", err
	}

	return playerData, nil
}
