package gate

import (
	"fmt"
	"github.com/liangdas/mqant/gate"
	"github.com/liangdas/mqant/log"
)

func (m *Module) handleHello(session gate.Session, req map[string]interface{}) (result string, err string) {
	if req["name"] == nil {
		result = "nothing"
		return
	}

	name := req["name"].(string)

	log.Info("handleHello: %v", name)

	return fmt.Sprintf("hello %v", name), ""
}
