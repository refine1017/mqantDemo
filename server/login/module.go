package login

import (
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	basemodule "github.com/liangdas/mqant/module/base"
)

type Module struct {
	basemodule.BaseModule
}

func (m *Module) GetType() string {
	return "Login"
}

func (m *Module) Version() string {
	return "1.0.0"
}

func (m *Module) OnInit(app module.App, settings *conf.ModuleSettings) {
	m.BaseModule.OnInit(m, app, settings)
}

func (m *Module) Run(closeSig chan bool) {

}

func (m *Module) OnDestroy() {
	if err := m.GetServer().OnDestroy(); err != nil {
		log.Warning("Module server destroy with err: %v", err)
	}
}