package repository

import "service/entity"

type RoomRepository interface {
	Create(body *entity.Payload) error
}
