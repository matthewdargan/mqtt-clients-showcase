// Copyright 2023 Matthew P. Dargan.
// SPDX-License-Identifier: Apache-2.0

// Pub is an MQTT publisher service that simulates sending room sensor data to a broker.
// The data is published to the specified MQTT topic at regular intervals.
package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/matthewdargan/mqtt-clients-showcase/internal/room"
)

const (
	server         = "tcp://broker.emqx.io:1883"
	username       = "emqx"
	password       = "public"
	disconnectWait = 250
	topic          = "go/living_room_temperature"
	delay          = 1 * time.Second
	minTemp        = 68
	maxTemp        = 74
)

func publishSensorData(client mqtt.Client, data room.SensorData) {
	msg, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v\n", err)
	}
	token := client.Publish(topic, 0, false, msg)
	if token.Wait() && token.Error() != nil {
		log.Fatalln(token.Error())
	}
	log.Printf(`Sent %s to topic "%s"`, msg, topic)
	time.Sleep(delay)
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

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	go func() {
		<-sigint
		log.Println("Disconnecting...")
		client.Disconnect(disconnectWait)
		os.Exit(0)
	}()

	for {
		for temp := minTemp; temp <= maxTemp; temp++ {
			publishSensorData(client, room.SensorData{Temperature: temp})
		}
		for temp := maxTemp - 1; temp > minTemp; temp-- {
			publishSensorData(client, room.SensorData{Temperature: temp})
		}
	}
}
