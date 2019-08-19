package game

import (
	"github.com/liangdas/mqant/gate"
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
