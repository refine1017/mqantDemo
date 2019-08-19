package db

import (
	"github.com/gomodule/redigo/redis"
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	basemodule "github.com/liangdas/mqant/module/base"
	"github.com/liangdas/mqant/utils"
)

type Module struct {
	basemodule.BaseModule
	redisFactory *utils.RedisFactory
	redis        redis.Conn
}

func (m *Module) GetType() string {
	return "DB"
}

func (m *Module) Version() string {
	return "1.0.0"
}

func (m *Module) OnInit(app module.App, settings *conf.ModuleSettings) {
	m.BaseModule.OnInit(m, app, settings)

	m.initRedis(settings)

	m.GetServer().Register("FindPlayerByUsername", m.rpcFindPlayerByUsername)
	m.GetServer().Register("CreatePlayerData", m.rpcCreatePlayerData)
	m.GetServer().Register("LoadPlayerDataByUsername", m.rpcLoadPlayerDataByUsername)
	m.GetServer().Register("SavePlayerData", m.rpcSavePlayerData)
}

func (m *Module) Run(closeSig chan bool) {

}

func (m *Module) OnDestroy() {
	if err := m.GetServer().OnDestroy(); err != nil {
		log.Warning("Module server destroy with err: %v", err)
	}
}
