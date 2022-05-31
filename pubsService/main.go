package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/nats-io/nats.go"
)

type Payload struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

const (
	StreamName     = "FORDB"
	StreamSubjects = "FORDB.*"
	SubjectName    = "FORDB.subject1"
)

func main() {
	nc, err := nats.Connect("localhost:4222")
	if err != nil {
		fmt.Printf("error connect, cause :  %+v \n", err)
	}

	js, err := nc.JetStream()
	if err != nil {
		fmt.Printf("error create jetStream ctx, cause :  %+v \n", err)
	}
	err = createStream(js)
	if err != nil {
		fmt.Printf("error create stream, cause :  %+v \n", err)
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
	var order Payload
	for i := 1; i <= 10; i++ {
		order = Payload{
			ID:   strconv.Itoa(i),
			Name: "Name-" + strconv.Itoa(i) + "from" + SubjectName,
			// Name: "Name.subj1-" + strconv.Itoa(i),
			// Name: "Name.subj2-" + strconv.Itoa(i),
		}
		orderJSON, _ := json.Marshal(order)
		_, err := js.Publish(SubjectName, orderJSON)
		if err != nil {
			return err
		}
		log.Printf("Order with OrderID:%d has been published to subject:%s \n", i, SubjectName)
	}
	return nil
}

// createStream creates a stream by using JetStreamContext
func createStream(js nats.JetStreamContext) error {
	// Check if the ORDERS stream already exists; if not, create it.
	stream, err := js.StreamInfo(StreamName)
	if err != nil {
		log.Println(err)
	}
	if stream == nil {
		log.Printf("creating stream %q and subjects %q", StreamName, StreamSubjects)
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     StreamName,
			Subjects: []string{StreamSubjects},
		})
		if err != nil {
			return err
		}
	}
	return nil
}
