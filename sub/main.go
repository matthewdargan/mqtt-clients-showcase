package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/matthewdargan/mqtt-clients-showcase/internal/room"
)

const (
	server         = "tcp://broker.emqx.io:1883"
	username       = "emqx"
	password       = "public"
	disconnectWait = 250
	topic          = "go/living_room_temperature"
	fanOnThreshold = 71
)

func messageHandler(client mqtt.Client, msg mqtt.Message) {
	var data room.SensorData
	err := json.Unmarshal(msg.Payload(), &data)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v\n", err)
	}

	fanStatus := "Off"
	if data.Temperature >= fanOnThreshold {
		fanStatus = "On"
	}
	log.Printf("Current Temperature: %d degrees - Fan Status: %s\n", data.Temperature, fanStatus)
}

func main() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(server)
	opts.SetUsername(username)
	opts.SetPassword(password)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalln(token.Error())
	}
	defer client.Disconnect(disconnectWait)

	if token := client.Subscribe(topic, 0, messageHandler); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		return
	}
	log.Printf("Subscribed to topic: %s\n", topic)
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint
	log.Println("Disconnecting...")
}
