package gate

import (
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/log"
)

func (m *Module) handleHello(session gate.Session, req map[string]interface{}) (result string, err string) {
	if req["msg"] == nil {
		result = "nothing"
		return
	}

	msg := req["msg"].(string)

	log.Info("handleHello: %v", msg)

	return msg, ""
}
