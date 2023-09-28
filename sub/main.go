// Copyright 2023 Matthew P. Dargan.
// SPDX-License-Identifier: Apache-2.0

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

func handleMessage(client mqtt.Client, msg mqtt.Message) {
	var data room.SensorData
	if err := json.Unmarshal(msg.Payload(), &data); err != nil {
		log.Printf("Error unmarshalling JSON: %v\n", err)
		return
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
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		log.Fatalln(token.Error())
	}
	defer client.Disconnect(disconnectWait)

	token = client.Subscribe(topic, 0, handleMessage)
	if token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		return
	}
	log.Printf("Subscribed to topic: %s\n", topic)
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint
	log.Println("Disconnecting...")
}
