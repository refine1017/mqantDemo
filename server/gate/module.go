package gate

import (
	"fmt"
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/gate"
	basegate "github.com/liangdas/mqant/gate/base"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	"time"
)

type Module struct {
	basegate.Gate
}

func (m *Module) GetType() string {
	return "Gate"
}

func (m *Module) Version() string {
	return "1.0.0"
}

func (m *Module) OnInit(app module.App, settings *conf.ModuleSettings) {
	m.Gate.OnInit(m, app, settings, gate.Heartbeat(time.Second * 10))

	m.GetServer().RegisterGO("HD_Hello", m.handleHello)

	_ = m.Gate.SetSessionLearner(m)
	_ = m.Gate.SetStorageHandler(m)
}

func (m *Module) Connect(session gate.Session) {
	log.Info("session connect from %v-%v", session.GetNetwork(), session.GetIP())
}

func (m *Module) DisConnect(session gate.Session) {
	log.Info("session disconnect from %v-%v", session.GetNetwork(), session.GetIP())
}

// storage session
func (m *Module) Storage(session gate.Session) error {
	log.Info("session#%v storage", session.GetUserId())

	return nil
}

// delete session
func (m *Module) Delete(session gate.Session) error {
	log.Info("session%%v delete", session.GetUserId())

	return nil
}

// Query user info
func (m *Module) Query(userId string) ([]byte, error) {
	log.Info("Query session %v", userId)
	return nil, fmt.Errorf("no implement")
}

// Session heartbeat
func (m *Module) Heartbeat(session gate.Session) {
	log.Info("session#%v heartbeat", session.GetUserId())
}
