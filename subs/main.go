package main

import (
	"encoding/json"
	"log"
	"runtime"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/shijuvar/go-distsys/jsdemo/model"
)

type Order struct {
	OrderID    int
	CustomerID string
	Status     string
}

func main() {
	// Connect to NATS
	nc, _ := nats.Connect("localhost:4222")
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}
	// Create durable consumer monitor
	js.Subscribe("ORDERS.*", func(msg *nats.Msg) {
		var order model.Order
		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("monitor service subscribes from subject:%s\n", msg.Subject)
		log.Printf("OrderID:%d, CustomerID: %s, Status:%s\n", order.OrderID, order.CustomerID, order.Status)
	}, nats.ManualAck(), nats.AckWait(60*time.Second)) //Ack wait (1 menit)

	runtime.Goexit()

}
