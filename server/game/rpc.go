package game

import (
	"encoding/json"
	"fmt"
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/log"
	"server/model"
)

func (m *Module) rpcLoginPlayer(session gate.Session, playerData string) (result string, err string) {
	var player = &model.Player{}

	if jsonErr := json.Unmarshal([]byte(playerData), player); jsonErr != nil {
		err = fmt.Sprintf("%v", jsonErr)
		return
	}

	m.playerManager.AddPlayer(player, session)

	log.Info("rpcLoginPlayer %v success", player.Id)
	return
}
