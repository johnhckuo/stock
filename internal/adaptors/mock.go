package adaptors

import (
	"codementordev/hello-johnhckuo/models"
	"errors"
	"sync/atomic"
)

type Mock struct {
}

var data map[int64][]models.Timeslot
var users []models.User
var dataCounter int64
var userCounter int64

func NewMock() *Mock {
	data = make(map[int64][]models.Timeslot)
	return &Mock{}
}

func (m *Mock) GetTimeslots(userId, before, after int64) ([]models.Timeslot, error) {
	if before == 0 && after == 0 {
		return data[userId], nil
	}

	var candidates []models.Timeslot
	for _, slot := range data[userId] {
		var isCandidate bool = true
		if before != 0 && *slot.EndAt > before {
			isCandidate = false
		}

		if after != 0 && *slot.EndAt < after {
			isCandidate = false
		}

		if isCandidate {
			candidates = append(candidates, slot)
		}
	}
	return candidates, nil
}

func (m *Mock) AddTimeslot(userId, startAt, endAt int64) (models.Timeslot, error) {
	if startAt > endAt {
		return models.Timeslot{}, errors.New("end date cannot be ealier than start date")
	}

	for _, val := range users {
		if val.ID == userId {
			//(StartA <= EndB)  and  (EndA >= StartB)
			for _, slot := range data[userId] {
				if startAt < *slot.EndAt && endAt > *slot.StartAt {
					return models.Timeslot{}, errors.New("timestamp overlap")
				}
			}
			atomic.AddInt64(&dataCounter, 1)
			newTimeslot := models.Timeslot{UserID: &userId, EndAt: &endAt, StartAt: &startAt, ID: dataCounter}
			data[userId] = append(data[userId], newTimeslot)
			return newTimeslot, nil
		}
	}

	return models.Timeslot{}, errors.New("user id doesn't exists")
}

func (m *Mock) DeleteTimeslot(userId, timeId int64) (bool, error) {
	if _, exists := data[userId]; exists {
		for i := range data[userId] {
			if data[userId][i].ID == timeId {
				length := len(data[userId])
				data[userId][i], data[userId][length-1] = data[userId][length-1], data[userId][i]
				data[userId] = data[userId][:length-1]
				break
			}
		}
		return true, nil
	}
	return false, errors.New("Cannot find corresponding user")
}

func (m *Mock) CreateUser(name string) (int64, error) {
	atomic.AddInt64(&userCounter, 1)
	newUser := models.User{ID: userCounter, Name: &name}
	users = append(users, newUser)
	return userCounter, nil
}
