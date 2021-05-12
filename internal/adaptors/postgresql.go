package adaptors

import (
	"codementordev/hello-johnhckuo/models"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres() (*Postgres, error) {
	// establish connection
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWD"), os.Getenv("DB_NAME")),
	)
	if err != nil {
		return nil, err
	}
	// check connection
	if err = db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Successfully created connection to database")

	return &Postgres{db: db}, nil
}

func (m *Postgres) GetTimeslots(userId, before, after int64) ([]models.Timeslot, error) {

	return []models.Timeslot{}, nil
}

func (m *Postgres) AddTimeslot(userId, startAt, endAt int64) (models.Timeslot, error) {

	return models.Timeslot{}, nil
}

func (m *Postgres) DeleteTimeslot(userId, timeId int64) (bool, error) {

	return false, nil
}

func (m *Postgres) CreateUser(name string) (int64, error) {
	return 0, nil
}
