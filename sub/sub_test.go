// Copyright 2023 Matthew P. Dargan.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/matthewdargan/mqtt-clients-showcase/internal/room"
)

type MockMQTTMessage struct {
	payload []byte
}

func (m *MockMQTTMessage) Duplicate() bool   { return false }
func (m *MockMQTTMessage) Qos() byte         { return 0 }
func (m *MockMQTTMessage) Retained() bool    { return false }
func (m *MockMQTTMessage) Topic() string     { return "" }
func (m *MockMQTTMessage) MessageID() uint16 { return 0 }
func (m *MockMQTTMessage) Ack()              {}
func (m *MockMQTTMessage) Payload() []byte {
	return m.payload
}

func TestHandleMessage(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		data     room.SensorData
		expected string
	}{
		{data: room.SensorData{Temperature: 50}, expected: "Current Temperature: 50 degrees - Fan Status: Off\n"},
		{data: room.SensorData{Temperature: 55}, expected: "Current Temperature: 55 degrees - Fan Status: Off\n"},
		{data: room.SensorData{Temperature: 60}, expected: "Current Temperature: 60 degrees - Fan Status: Off\n"},
		{data: room.SensorData{Temperature: 65}, expected: "Current Temperature: 65 degrees - Fan Status: Off\n"},
		{data: room.SensorData{Temperature: 70}, expected: "Current Temperature: 70 degrees - Fan Status: Off\n"},
		{data: room.SensorData{Temperature: 75}, expected: "Current Temperature: 75 degrees - Fan Status: On\n"},
		{data: room.SensorData{Temperature: 80}, expected: "Current Temperature: 80 degrees - Fan Status: On\n"},
		{data: room.SensorData{Temperature: 85}, expected: "Current Temperature: 85 degrees - Fan Status: On\n"},
		{data: room.SensorData{Temperature: 90}, expected: "Current Temperature: 90 degrees - Fan Status: On\n"},
		{data: room.SensorData{Temperature: 95}, expected: "Current Temperature: 95 degrees - Fan Status: On\n"},
		{data: room.SensorData{Temperature: 100}, expected: "Current Temperature: 100 degrees - Fan Status: On\n"},
	}

	for _, testCase := range testCases {
		payload, err := json.Marshal(testCase.data)
		if err != nil {
			t.Fatalf("Error marshalling JSON: %v", err)
		}
		mockMsg := &MockMQTTMessage{payload: payload}

		// Capture the log output
		var buf bytes.Buffer
		log.SetOutput(&buf)
		handleMessage(nil, mockMsg)
		log.SetOutput(os.Stdout)

		got := buf.String()[20:] // Remove date and time from log before comparison
		if got != testCase.expected {
			t.Errorf("For temperature %d, got %s, want %s", testCase.data.Temperature, got, testCase.expected)
		}
	}
}
