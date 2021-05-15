package handlers

import (
	"codementordev/hello-johnhckuo/internal/adaptors"
	"codementordev/hello-johnhckuo/models"
	"log"
	"os"
)

type Handler struct {
	db adaptors.DB
}

var handler Handler

// we use postgres by default
func init() {
	var err error
	if currentEnv, exist := os.LookupEnv("DB_MODE"); exist && currentEnv == "MOCK" {
		log.Println("Using Mock DB as data storage")
		handler.db = adaptors.NewMock()
	} else {
		log.Println("Using PostgreSQL as data storage")
		handler.db, err = adaptors.NewPostgres()
		if err != nil {
			log.Fatalf("Failed initializing DB connection: %v", err.Error())
		}
	}
	log.Println("DB Connection established...")
}

// Retrieve time slots based on userId & endAt filter
func GetTimeslots(userId, before, after int64) ([]*models.Timeslot, error) {
	return handler.db.GetTimeslots(userId, before, after)
}

// Add new time slot and return time slot id
func AddTimeslot(userId int64, newTime *models.Timeslot) (*models.Timeslot, error) {
	return handler.db.AddTimeslot(userId, *newTime.StartAt, *newTime.EndAt)
}

// Delete time slot according to given id
func DeleteTimeslot(userId int64, timeId int64) (bool, error) {
	return handler.db.DeleteTimeslot(userId, timeId)
}

// Create new user
func CreateUser(name string) (int64, error) {
	return handler.db.CreateUser(name)
}
