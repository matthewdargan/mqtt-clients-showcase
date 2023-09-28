// Copyright 2023 Matthew P. Dargan.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"encoding/json"
	"reflect"
	"testing"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/matthewdargan/mqtt-clients-showcase/internal/room"
)

func TestPublishSensorDataIntegration(t *testing.T) {
	t.Parallel()
	opts := mqtt.NewClientOptions().AddBroker(server)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		t.Fatalf("Failed to connect: %v", token.Error())
	}
	defer client.Disconnect(disconnectWait)

	for temp := minTemp; temp <= maxTemp; temp++ {
		data := room.SensorData{Temperature: temp}
		token := client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
			var receivedData room.SensorData
			if err := json.Unmarshal(msg.Payload(), &receivedData); err != nil {
				t.Fatalf("Error unmarshalling JSON: %v", err)
			}
			if !reflect.DeepEqual(receivedData, data) {
				t.Errorf("Received incorrect data, got %v, want %v", receivedData, data)
			}
		})
		if token.Wait() && token.Error() != nil {
			t.Fatalf("Failed to subscribe: %v", token.Error())
		}
		publishSensorData(client, data)
	}
}
