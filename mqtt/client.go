package mqtt

import (
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("\n MQTT Connected")
	// MQTT subscribe to $SYS topics
	SubscribeToSysTopics(client, Topics)
	online.Set(1)
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("\n MQTT Connection lost: %v", err)
}

func Init() mqtt.Client {
	var MQTT_BROKER = os.Getenv("MQTT_BROKER")
	var MQTT_PORT = os.Getenv("MQTT_PORT")
	var MQTT_USERNAME = os.Getenv("MQTT_USERNAME")
	var MQTT_PASSWORD = os.Getenv("MQTT_PASSWORD")
	var MQTT_CLIENT_ID = os.Getenv("MQTT_CLIENT_ID")
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", MQTT_BROKER, MQTT_PORT))
	opts.SetClientID(MQTT_CLIENT_ID)
	opts.SetUsername(MQTT_USERNAME)
	opts.SetPassword(MQTT_PASSWORD)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}
