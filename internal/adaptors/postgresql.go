package adaptors

import (
	"codementordev/hello-johnhckuo/models"
	"errors"
	"fmt"
	"os"

	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sqlx.DB
}

func NewPostgres() (*Postgres, error) {

	// establish connection
	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWD"), os.Getenv("DB_NAME")),
	)
	if err != nil {
		return nil, err
	}
	// check connection
	if err = db.Ping(); err != nil {
		return nil, err
	}
	log.Printf("Successfully created connection to database")

	return &Postgres{db: db}, nil
}

func (m *Postgres) GetTimeslots(userId, before, after int64) ([]*models.Timeslot, error) {
	log.Printf("GetTimeslots user_id: %v, before: %v, after: %v", userId, before, after)
	timeslots := []*models.Timeslot{}
	mappedInput := map[string]interface{}{"userId": userId}

	statement := "SELECT * FROM timeslots WHERE UserId=:userId"
	if before != 0 {
		statement += " AND EndAt < :before"
		mappedInput["before"] = before
	}
	if after != 0 {
		statement += " AND EndAt > :after"
		mappedInput["after"] = after
	}
	statement += " ORDER BY StartAt ASC;"

	rows, err := m.db.NamedQuery(statement, mappedInput)
	if err != nil {
		return []*models.Timeslot{}, err
	}
	for rows.Next() {
		var temp models.Timeslot
		err := rows.StructScan(&temp)
		if err != nil {
			return []*models.Timeslot{}, err
		}
		timeslots = append(timeslots, &temp)
	}

	log.Printf("Timeslots successfully retrieved=> User_ID: %v, timeslots retrieved: %v", userId, len(timeslots))

	return timeslots, nil
}

func (m *Postgres) AddTimeslot(userId, startAt, endAt int64) (*models.Timeslot, error) {
	log.Printf("AddTimeslot user_id: %v, start_at: %v, end_at: %v", userId, startAt, endAt)
	if startAt == 0 || endAt == 0 {
		return &models.Timeslot{}, errors.New("Insufficient paramenters")
	}
	var newId int64
	err := m.db.Get(&newId, `SELECT add_timeslot(user_id := $1, start_at := $2, end_at := $3);`, userId, startAt, endAt)
	if err != nil {
		return &models.Timeslot{}, err
	}

	if newId == -1 {
		return &models.Timeslot{}, errors.New("timeslot overlapped")
	}

	log.Printf("Timeslot successfully added=> User_ID: %v, startAt: %v, EndAt: %v", userId, startAt, endAt)

	return &models.Timeslot{ID: newId, UserID: userId, StartAt: &startAt, EndAt: &endAt}, nil
}

func (m *Postgres) DeleteTimeslot(userId, timeId int64) (bool, error) {
	log.Printf("DeleteTimeslot user_id: %v, time_id: %v", userId, timeId)
	res, err := m.db.NamedExec(`DELETE FROM timeslots WHERE ID=:timeId AND UserID=:userId;`,
		map[string]interface{}{
			"timeId": timeId,
			"userId": userId,
		})
	if err != nil {
		return false, err
	}

	count, err := res.RowsAffected()
	if count == 0 || err != nil {
		return false, err
	}

	log.Printf("Timeslot successfully deleted=> User_ID: %v, Time_ID: %v", userId, timeId)

	return true, nil
}

func (m *Postgres) CreateUser(name string) (int64, error) {

	// tx := m.db.MustBegin()

	// res := tx.MustExec("INSERT INTO Users (Name) VALUES ($1) RETURNING ID", name)
	// count, err := res.RowsAffected()
	// if count == 0 || err != nil {
	// 	return -1, err
	// }
	// tx.Commit()

	log.Printf("CreateUser name: %v", name)

	var id int64
	err := m.db.QueryRowx(`INSERT INTO users (Name) VALUES ($1) RETURNING ID`, name).Scan(&id)
	if err != nil {
		return -1, err
	}

	log.Printf("User successfully created=> ID: %v, name: %v", id, name)

	return id, nil
}
