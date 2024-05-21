package gomqtt

import (
	logger "github.com/kordar/gologger"
)

func Pub(topic string, message string, qos byte, retained bool) {
	token := mqttclient.Publish(topic, qos, retained, message)
	token.Wait()
	if token.Error() != nil {
		logger.Warnf("[publish mqtt] token err = %v", token.Error())
	}
}
