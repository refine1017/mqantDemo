package db

import (
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/utils"
)

const (
	RedisPlayers = "mqant_players"
	RedisPlayerAutoId = "mqant_player_auto_id"
)

func (m *Module) initRedis(settings *conf.ModuleSettings) {
	m.redisFactory = utils.GetRedisFactory()
	var redisPool = m.redisFactory.GetPool(settings.Redis.Uri)
	var err error
	m.redis, err = redisPool.Dial()
	if err != nil {
		log.Error("Connect redis with err: %v", err)
	} else {
		log.Info("Connect redis success")
	}
}