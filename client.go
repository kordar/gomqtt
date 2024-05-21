package gomqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/kordar/gocfg"
	logger "github.com/kordar/gologger"
	"github.com/kordar/goutil"
)

var mqttclient mqtt.Client

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	logger.Infof("[mqtt] Received message: %s from topic: %s", msg.Payload(), msg.Topic())
}

func SetMessagePubHandler(f func(client mqtt.Client, msg mqtt.Message)) {
	messagePubHandler = f
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	logger.Infof("[mqtt] Connected")
}

func SetConnectHandler(f func(client mqtt.Client)) {
	connectHandler = f
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	logger.Infof("[mqtt] Connect lost: %v", err)
}

func SetConnectLostHandler(f func(client mqtt.Client, err error)) {
	connectLostHandler = f
}

func CreateMqttClient(broker string, port int, id string, username string, password string) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(id)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetAutoReconnect(true)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logger.Errorf("[mqtt], err = %v", token.Error())
		return nil
	}
	return client
}

func InitMqttClient() {
	var broker = gocfg.GetSectionValue("mqtt", "broker")
	if broker == "" {
		return
	}
	var port = gocfg.GetSectionValueInt("mqtt", "port")
	clientId := gocfg.GetSectionValue("mqtt", "id")
	if clientId == "" {
		clientId = goutil.UUID()
	}
	username := gocfg.GetSectionValue("mqtt", "username")
	password := gocfg.GetSectionValue("mqtt", "password")
	client := CreateMqttClient(broker, port, clientId, username, password)
	if client != nil {
		mqttclient = client
	}
}
