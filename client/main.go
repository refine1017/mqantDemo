package main

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/liangdas/armyant/work"
	"github.com/liangdas/mqant/log"
)

func main() {
	client := new(work.MqttWork)

	opts := client.GetDefaultOptions("ws://127.0.0.1:3653")
	opts.SetOnConnectHandler(func(client MQTT.Client) {
		log.Info("client connect")
	})
	opts.SetConnectionLostHandler(func(client MQTT.Client, err error) {
		log.Info("client disconnect")
	})

	if err := client.Connect(opts); err != nil {
		panic(err)
	}

	res, err := client.Request("Gate/HD_Hello", []byte(`{"name":"mqant"}`))
	if err != nil {
		panic(err)
	}

	log.Info("res: topic=%v, body=%s", res.Topic(), res.Payload())
}
