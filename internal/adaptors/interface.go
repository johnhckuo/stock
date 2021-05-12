package adaptors

import "codementordev/hello-johnhckuo/models"

type DB interface {
	GetTimeslots(userId, before, after int64) ([]models.Timeslot, error)
	AddTimeslot(userId, startAt, endAt int64) (models.Timeslot, error)
	DeleteTimeslot(userId, timeId int64) (bool, error)
	CreateUser(name string) (int64, error)
}
