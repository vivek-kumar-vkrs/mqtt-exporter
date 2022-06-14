package mqtt

import (
	"strconv"

	"fmt"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/prometheus/client_golang/prometheus"
)

func subHandler(client mqtt.Client, message mqtt.Message, handlerName string, gauge prometheus.Gauge) {
	if handlerName == "uptime" {
		uptimeSubHandler(client, message, handlerName, gauge)
	} else if handlerName == "version" {
		versionSubHandler(client, message, handlerName, gauge)
	} else {
		t := string(message.Payload())

		s, err := strconv.ParseFloat(t, 64)
		if err == nil {
			fmt.Printf("\n subHandler %s Parsed float: %f", handlerName, s)
			gauge.Set(s)
		} else {
			fmt.Printf("\n Error subHandler %s", handlerName)
		}
	}

}

func uptimeSubHandler(client mqtt.Client, message mqtt.Message, handlerName string, gauge prometheus.Gauge) {
	str := string(message.Payload())
	strL := strings.Split(strings.TrimSpace(str), "seconds")
	s, err := strconv.ParseFloat(strings.TrimSpace(strL[0]), 8)
	if err == nil {
		fmt.Printf("\n subHandler %s Parsed float: %f", handlerName, s)
		gauge.Set(s)
	} else {
		fmt.Printf("\n Error subHandler %s", handlerName)
	}
}

func versionSubHandler(client mqtt.Client, message mqtt.Message, handlerName string, gauge prometheus.Gauge) {
	str := string(message.Payload())
	Version = strings.Split(strings.TrimSpace(str), "mosquitto version ")[1]
}

func SubscribeToSysTopics(client mqtt.Client, topics []Topic) {
	for i := 0; i < len(topics); i++ {
		topic := topics[i]
		if topic.Topic == "" {
			continue
		}
		token := client.Subscribe(topic.Topic, 1, func(client mqtt.Client, message mqtt.Message) {
			subHandler(client, message, topic.Name, topic.Gauge)
		})
		token.Wait()
		fmt.Printf("\n Subscribed to topic %s", topic.Topic)
	}
}

func ResetGauge() {
	for i := 0; i < len(Topics); i++ {
		topic := Topics[i]
		topic.Gauge.Set(0)
	}

	// mark broker offline
	online.Set(0)
}
