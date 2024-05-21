package gomqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	logger "github.com/kordar/gologger"
)

func Sub(topic string, qos byte, f func(client mqtt.Client, message mqtt.Message)) {
	token := mqttclient.Subscribe(topic, qos, f)
	token.Wait()
	logger.Infof("[mqtt] Subscribed to topic: %s", topic)
}
