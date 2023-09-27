package main

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	server         = "tcp://broker.emqx.io:1883"
	username       = "emqx"
	password       = "public"
	disconnectWait = 250
	topic          = "go/living_room_temperature"
	delay          = 1 * time.Second
	minTemp        = 68
	maxTemp        = 73
)

func main() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(server)
	opts.SetUsername(username)
	opts.SetPassword(password)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}
	defer client.Disconnect(disconnectWait)

	for temp := minTemp; temp < maxTemp; temp++ {
		msg := fmt.Sprintf("%d", temp)
		token := client.Publish(topic, 0, false, msg)
		token.Wait()
		fmt.Printf("Sent %s to topic %s\n", msg, topic)
		time.Sleep(delay)
	}
}
