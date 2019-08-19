package game

import (
	"encoding/json"
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/log"
	"server/model"
)

type PlayerManager struct {
	sessions map[int]gate.Session
	players  map[int]*model.Player
}

func NewPlayerManager() *PlayerManager {
	return &PlayerManager{
		sessions: make(map[int]gate.Session),
		players:  make(map[int]*model.Player),
	}
}

func (pm *PlayerManager) AddPlayer(player *model.Player, session gate.Session) {
	pm.sessions[player.Id] = session
	pm.players[player.Id] = player
}

func (pm *PlayerManager) Broadcast(topic string, body []byte) {
	for _, session := range pm.sessions {
		session.Send(topic, body)
	}
}

func (m *Module) SavePlayer(player *model.Player) {
	buf, err := json.Marshal(player)
	if err != nil {
		log.Error("SavePlayer with err: %v", err)
		return
	}

	err = m.RpcInvokeNR("DB", "SavePlayerData", player.Username, string(buf))
	if err != nil {
		log.Error("SavePlayer with err: %v", err)
	}
}
