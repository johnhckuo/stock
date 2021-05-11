package handlers

import "codementordev/hello-johnhckuo/models"

// Retrieve time slots based on userId & endAt filter
func GetTimeSlots(userId string, endAt int64) []models.Timeslot {
	return []models.Timeslot{}
}

// Add new time slot and return time slot id
func AddTimeSlot(userId string, newTime models.Timeslot) models.Timeslot {
	return models.Timeslot{}
}

// Delete time slot according to given id
func DeleteTimeSlot(userId string, timeId int) bool {
	return true
}

// Create new user
func CreateUser(name string) (int64, error) {
	return 12345, nil
}
