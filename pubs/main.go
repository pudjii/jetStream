package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/nats-io/nats.go"
)

type Order struct {
	OrderID    int
	CustomerID string
	Status     string
}

const (
	streamName     = "TEST"
	streamSubjects = "TEST.*"
	subjectName    = "TEST.subject1"
)

func main() {

	// Connect to NATS
	nc, err := nats.Connect("localhost:4222")
	if err != nil {
		fmt.Printf("error connect coz :  %+v \n", err)
	}
	// Creates JetStreamContext
	js, err := nc.JetStream()
	if err != nil {
		fmt.Printf("error create jetStream ctx coz :  %+v \n", err)
	}
	// Creates stream (create chanel/stream)
	err = createStream(js)
	if err != nil {
		fmt.Printf("error create stream coz :  %+v \n", err)
	}
	// Create orders by publishing messages (Pubs)
	err = createOrder(js)
	if err != nil {
		fmt.Printf("error pubs coz :  %+v \n", err)
	}

}

// createOrder publishes stream of events
// with subject "ORDERS.created"
func createOrder(js nats.JetStreamContext) error {
	var order Order
	for i := 1; i <= 10; i++ {
		order = Order{
			OrderID:    i,
			CustomerID: "Cust-" + strconv.Itoa(i),
			Status:     "created",
		}
		orderJSON, _ := json.Marshal(order)
		_, err := js.Publish(subjectName, orderJSON)
		if err != nil {
			return err
		}
		log.Printf("Order with OrderID:%d has been published to subject:%s \n", i, subjectName)
	}
	return nil
}

// createStream creates a stream by using JetStreamContext
func createStream(js nats.JetStreamContext) error {
	// Check if the ORDERS stream already exists; if not, create it.
	stream, err := js.StreamInfo(streamName)
	if err != nil {
		log.Println(err)
	}
	if stream == nil {
		log.Printf("creating stream %q and subjects %q", streamName, streamSubjects)
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{streamSubjects},
		})
		if err != nil {
			return err
		}
	}
	return nil
}
