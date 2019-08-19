package db

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"server/model"
)

func (m *Module) ExistPlayer(username string) (bool, error) {
	return redis.Bool(m.redis.Do("HEXISTS", RedisPlayers, username))
}

func (m *Module) CreatePlayer(username string, nickname string) (string, error) {
	id, err := redis.Int(m.redis.Do("INCR", RedisPlayerAutoId))
	if err != nil {
		return "", err
	}

	player := &model.Player{}
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

func (m *Module) LoadPlayer(username string) (string, error) {
	return redis.String(m.redis.Do("HGET", RedisPlayers, username))
}

func (m *Module) SavePlayer(username string, player string) (bool, error) {
	return redis.Bool(m.redis.Do("HSET", RedisPlayers, username, player))
}
