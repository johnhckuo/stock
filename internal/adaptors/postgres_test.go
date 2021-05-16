package adaptors

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var psql DB
var psql_timeId, psql_timeId2, psql_timeId3 int64

func init() {
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_NAME", "timetable")
	os.Setenv("DB_USER", "user")
	os.Setenv("DB_PASSWD", "mysecretpassword")
	os.Setenv("DB_MODE", "PSQL")
}

func TestConnection_PSQL_failed(t *testing.T) {
	os.Setenv("DB_HOST", "256.0.0.2")
	_, err := NewPostgres()
	assert.NotNil(t, err)
}

func TestConnection_PSQL(t *testing.T) {
	var err error
	os.Setenv("DB_HOST", "127.0.0.1")
	psql, err = NewPostgres()
	assert.Nil(t, err)
	t.Logf("DB connected")
	if err != nil {
		t.Fatalf("No DB connection established: %v", err.Error())
	}
}

func TestCreateUser_PSQL_failed(t *testing.T) {
	var err error
	userId, err = psql.CreateUser("")
	assert.NotNil(t, err)
}

func TestCreateUser_PSQL(t *testing.T) {
	var err error
	userId, err = psql.CreateUser("john")
	assert.Nil(t, err)
}

func TestCreateTimeslot_PSQL_failed1(t *testing.T) {
	_, err := psql.AddTimeslot(userId, 0, 0)
	assert.NotNil(t, err)
}

func TestCreateTimeslot_PSQL_failed2(t *testing.T) {
	_, err := psql.AddTimeslot(userId, -1, -1)
	assert.NotNil(t, err)
}

func TestCreateTimeslot_PSQL(t *testing.T) {
	in, _ := time.Parse(time.RFC822, "01 Jan 15 10:00 UTC")
	out, _ := time.Parse(time.RFC822, "01 Jan 16 10:00 UTC")
	slot, err := psql.AddTimeslot(userId, in.Unix(), out.Unix())
	assert.Nil(t, err)
	assert.NotEqual(t, int64(-1), slot.ID, "should not be -1")
	psql_timeId = slot.ID
}

func TestCreateTimeslot2_PSQL(t *testing.T) {
	in, _ := time.Parse(time.RFC822, "01 Jan 16 10:00 UTC")
	out, _ := time.Parse(time.RFC822, "01 Jan 17 10:00 UTC")
	slot, err := psql.AddTimeslot(userId, in.Unix(), out.Unix())
	assert.Nil(t, err)
	assert.NotEqual(t, int64(-1), slot.ID, "should not be -1")
	psql_timeId2 = slot.ID
}

func TestCreateTimeslot3_PSQL(t *testing.T) {
	in, _ := time.Parse(time.RFC822, "01 Jan 17 10:00 UTC")
	out, _ := time.Parse(time.RFC822, "01 Jan 18 10:00 UTC")
	slot, err := psql.AddTimeslot(userId, in.Unix(), out.Unix())
	assert.Nil(t, err)
	assert.NotEqual(t, int64(-1), slot.ID, "should not be -1")
	psql_timeId3 = slot.ID
}

func TestCreateTimeslot4_PSQL(t *testing.T) {
	in, _ := time.Parse(time.RFC822, "01 Jan 18 9:59 UTC")
	out, _ := time.Parse(time.RFC822, "01 Jan 20 10:00 UTC")
	_, err := psql.AddTimeslot(userId, in.Unix(), out.Unix())
	assert.NotNil(t, err, "timeslot overlapped")
}

// check get
func TestGetTimeslot1_PSQL(t *testing.T) {
	timeslots, err := psql.GetTimeslots(userId, 0, 0)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(timeslots), "should have 3 timeslots")
}

//check before & after
func TestGetTimeslot2_PSQL(t *testing.T) {
	before, _ := time.Parse(time.RFC822, "01 Jan 17 11:00 UTC")
	after, _ := time.Parse(time.RFC822, "01 Jan 16 11:00 UTC")
	timeslots, err := psql.GetTimeslots(userId, before.Unix(), after.Unix())
	assert.Nil(t, err)
	assert.Equal(t, 1, len(timeslots), "should have 1 timeslots")
}

//check before
func TestGetTimeslot3_PSQL(t *testing.T) {
	before, _ := time.Parse(time.RFC822, "01 Jan 17 11:00 UTC")
	timeslots, err := psql.GetTimeslots(userId, before.Unix(), 0)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(timeslots), "should have 2 timeslots")
}

//check after
func TestGetTimeslot4_PSQL(t *testing.T) {
	after, _ := time.Parse(time.RFC822, "01 Jan 17 11:00 UTC")
	timeslots, err := psql.GetTimeslots(userId, 0, after.Unix())
	assert.Nil(t, err)
	assert.Equal(t, 1, len(timeslots), "should have 1 timeslots")
}

func TestAddTimeslot_PSQL(t *testing.T) {
	in, _ := time.Parse(time.RFC822, "01 Jan 15 20:00 UTC")
	out, _ := time.Parse(time.RFC822, "01 Jan 17 10:00 UTC")
	_, err := psql.AddTimeslot(userId, in.Unix(), out.Unix())
	assert.NotNil(t, err, "timeslot overlapped")
}

func TestDeleteTimeslot_PSQL(t *testing.T) {
	success, err := psql.DeleteTimeslot(userId, psql_timeId)
	assert.Nil(t, err)
	assert.Equal(t, true, success, "delete successfully")
}

func TestDeleteTimeslot_PSQL_failed(t *testing.T) {
	_, err := psql.DeleteTimeslot(-1, 01201240)
	assert.NotNil(t, err)
}

func TestGetTimeslot5_PSQL(t *testing.T) {
	timeslots, err := psql.GetTimeslots(userId, 0, 0)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(timeslots), "should have none timeslots")
}
