package game

import "github.com/liangdas/mqant/gate"

func (m *Module) handleTalk(session gate.Session, req map[string]interface{}) (result string, err string) {
	if session.Get("login") == "" {
		return "", "no login"
	}

	msg := req["msg"].(string)

	m.playerManager.Broadcast("Game/TalkNotify", []byte(msg))

	return "success", err
}
