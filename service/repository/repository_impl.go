package repository

import (
	"context"
	"log"
	"service/config"
	"service/entity"

	"github.com/arangodb/go-driver"
)

type roomRepository struct {
	DBLive driver.Database
}

func NewRoomRepository(conn *config.DBConn) RoomRepository {
	var db, err = conn.DBLive.Database(context.Background(), config.DBNAME)
	if err != nil {
		panic(err)
	}

	return &roomRepository{DBLive: db}
}

func (c *roomRepository) Create(body *entity.Payload) error {
	ctx := config.Ctx
	col, err := c.DBLive.Collection(ctx, "name")
	if err != nil {
		log.Printf("Error getting collection, cause: %v", err)
		return err
	}
	_, err = col.CreateDocument(ctx, body)
	if err != nil {
		log.Printf("Error when save data to collection, cause: %v", err)
		return err
	}

	return nil
}
