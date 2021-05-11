package models

type TimeSlot struct {
	Id      int64
	UserId  int64
	StartAt int64
	EndAt   int64
}

type User struct {
	Id   int64
	Name string
}
