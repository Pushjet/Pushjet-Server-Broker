package main

import (
	"encoding/json"
	"fmt"
	zmq "github.com/pebbe/zmq4"
	"log"
	"os"
	"os/signal"
	"syscall"
	"flag"
)

var SocketRelayName string
var SocketPubName   string

func main() {
	flag.StringVar(&SocketRelayName, "r", "ipc:///tmp/pushjet-relay.ipc", "Relay socket")
	flag.StringVar(&SocketPubName  , "p", "ipc:///tmp/pushjet-publisher.ipc", "Publish socket")
	flag.Parse()

	log.Println("Starting up the publishing server")

	context, _ := zmq.NewContext()

	socketRelay, err := context.NewSocket(zmq.PULL)
	if err != nil {
		log.Fatalf("Could not create a ZeroMQ socket: %s", err)
	}
	socketPub, err := context.NewSocket(zmq.PUB)
	if err != nil {
		log.Fatalf("Could not create a ZeroMQ socket: %s", err)
	}

	err = socketRelay.Bind(SocketRelayName)
	if err != nil {
		log.Fatalf("Could not create the ZeroMQ relay socket: %s", err)
	}
	err = socketPub.Bind(SocketPubName)
	if err != nil {
		log.Fatal("Could not create the ZeroMQ publisher socket: %s", err)
	}

	// Catch signals
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-signalChannel
		if sig == os.Interrupt || sig == syscall.SIGTERM {
			socketRelay.Unbind(SocketRelayName)
			socketPub.Unbind(SocketPubName)
			log.Fatal("Caught signal!")
		}
	}()

	log.Printf("Listening on '%s' and '%s'", SocketRelayName, SocketPubName)

	var apiMessageRaw string
	var apiMessage PushjetApiCall
	for { // loop forever
		apiMessage = PushjetApiCall{}
		apiMessageRaw, err = socketRelay.Recv(0)
		if err != nil {
			continue
		}

		log.Println("Parsing message... ")

		err = json.Unmarshal([]byte(apiMessageRaw), &apiMessage)
		if err != nil {
			log.Println("ERROR: Could not decode message sent by server. Skipping it; ", err)
			continue
		}

		if apiMessage.Message.Timestamp > 0 {
			log.Println("Sending out message for ", apiMessage.Message.Service.Public)
			socketPub.Send(fmt.Sprintf("%s %s", apiMessage.Message.Service.Public, apiMessageRaw), 0)
		}
		if apiMessage.Listen.Timestamp > 0 {
			log.Println("Sending out listen update for ", apiMessage.Listen.Uuid)
			socketPub.Send(fmt.Sprintf("%s %s", apiMessage.Listen.Uuid, apiMessageRaw), 0)
		}
	}
}
