package main

import (
	"github.com/liangdas/mqant"
	"github.com/liangdas/mqant/module"
	"github.com/liangdas/mqant/registry"
	"github.com/nats-io/go-nats"
	"server/db"
	"server/game"
	"server/gate"
	"server/login"
)

func main() {
	// connect nats
	nc, err := nats.Connect(nats.DefaultURL, nats.MaxReconnects(10000))
	if err != nil {
		panic(err)
	}

	// create app
	app := mqant.CreateApp(module.Nats(nc), module.Registry(registry.DefaultRegistry))

	// run app
	if err := app.Run(true,
		new(db.Module),
		new(game.Module),
		new(login.Module),
		new(gate.Module),
	); err != nil {
		panic(err)
	}
}
