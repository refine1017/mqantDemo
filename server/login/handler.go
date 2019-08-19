package login

import (
	"encoding/json"
	"fmt"
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/log"
	"server/model"
)

func (m *Module) handleRegister(session gate.Session, req map[string]interface{}) (result string, err string) {
	username := req["username"].(string)
	nickname := req["nickname"].(string)

	var rpcResult interface{}
	rpcResult, err = m.RpcInvoke("DB", "FindPlayerByUsername", username)
	if err != "" {
		return
	}

	exists := rpcResult.(bool)
	if exists {
		return "", "account already exists"
	}

	rpcResult, err = m.RpcInvoke("DB", "CreatePlayerData", username, nickname)
	if err != "" {
		return
	}

	log.Info("handleRegister: username=%v, nickname=%v", username, nickname)
	return rpcResult.(string), err
}

func (m *Module) handleLogin(session gate.Session, req map[string]interface{}) (result string, err string) {
	username := req["username"].(string)

	var rpcResult interface{}
	rpcResult, err = m.RpcInvoke("DB", "LoadPlayerDataByUsername", username)
	if err != "" {
		return
	}

	var player = &model.Player{}

	if jsonErr := json.Unmarshal([]byte(rpcResult.(string)), player); jsonErr != nil {
		err = fmt.Sprintf("%v", jsonErr)
		return
	}

	if _, err = m.RpcInvoke("Game", "LoginPlayer", session, rpcResult); err != "" {
		return
	}

	if err = session.Bind(fmt.Sprintf("%v", player.Id)); err != "" {
		return
	}

	session.Set("login", "true")
	session.Push()

	log.Info("handleLogin: username=%v, id=%v", username, player.Id)
	return rpcResult.(string), err
}
