package login

import "github.com/liangdas/mqant/gate"

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

	return rpcResult.(string), err
}

func (m *Module) handleLogin(session gate.Session, req map[string]interface{}) (result string, err string) {
	return "", ""
}