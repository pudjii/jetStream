package subscribes

import (
	"encoding/json"
	"log"
	"service/config"
	"service/entity"
	"service/repository"
	"time"

	"github.com/nats-io/nats.go"
)

type (
	Payload struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	subs struct {
		repo repository.RoomRepository
	}
)

func Listen(repo repository.RoomRepository) {
	nc, err := nats.Connect(config.NATSURL)
	if err != nil {
		panic(err)
	}
	js, err := nc.JetStream()
	if err != nil {
		panic(err)
	}
	//add stream
	_, err = js.AddStream(&nats.StreamConfig{
		Name:     config.StreamName,
		Subjects: []string{config.StreamSubjects},
	})
	if err != nil {
		panic(err)
	}

	subs := subs{repo}
	subs.Subs(js)
}

func (s *subs) Subs(sc nats.JetStreamContext) {
	var subs, err = sc.Subscribe(config.SubjectName, func(msg *nats.Msg) {
		var payload *Payload
		err := json.Unmarshal(msg.Data, &payload)
		if err != nil {
			panic(err)
		}
		err = s.repo.Create(&entity.Payload{
			ID:   payload.ID,
			Name: payload.Name,
		})
		if err != nil {
			log.Printf("Error create, cause: %+v \n", err)
			return
		}
		msg.Ack()
	}, nats.ManualAck(), nats.AckWait(60*time.Second))
	if err != nil {
		panic(err)
	}
	log.Println("Subs order :", subs.IsValid())
}
