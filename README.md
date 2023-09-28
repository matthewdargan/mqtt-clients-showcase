# MQTT Clients Showcase

[![Build Status](https://github.com/matthewdargan/mqtt-clients-showcase/actions/workflows/go-ci.yml/badge.svg?branch=main)](https://github.com/matthewdargan/mqtt-clients-showcase/actions/workflows/go-ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/matthewdargan/mqtt-clients-showcase)](https://goreportcard.com/report/github.com/matthewdargan/mqtt-clients-showcase)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](LICENSE)

The MQTT Clients Showcase simulates an MQTT publisher and subscriber for room sensor data.

## Installation

In order to use the MQTT Clients Showcase, install [Go](https://go.dev/doc/install) on your operating system.

## Running the MQTT Clients

The showcase includes two main components: `pub` and `sub`, which simulate the MQTT publisher and subscriber services, respectively.

### Running the MQTT Publisher (`pub`)

To run the MQTT publisher, run the following command in the project root directory:

```sh
go run pub/main.go
```

### Running the MQTT Subscriber (`sub`)

To run the MQTT subscriber, run the following command in the project root directory:

```sh
go run sub/main.go
```

## Testing

To run automated unit and integration tests, use the following command:

```sh
go test ./...
```
