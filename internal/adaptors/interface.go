package adaptor

import "codementordev/hello-johnhckuo/internal/models"

type DB interface {
	GetTimeSlots(userId string, endAt int64) ([]models.TimeSlot, error)
	AddTimeSlot(userId string, newTime models.TimeSlot) (models.TimeSlot, error)
	DeleteTimeSlot(userId string, timeId int) (bool, error)
	CreateUser(name string) (int64, error)
}
