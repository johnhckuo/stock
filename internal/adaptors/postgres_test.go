package adaptors

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var psql DB
var psql_timeId, psql_timeId2, psql_timeId3 int64

func TestConnection_PSQL(t *testing.T) {
	var err error
	psql, err = NewPostgres()
	if err != nil {
		t.Errorf("Error when connecting to DB: %v", err)
		return
	}
	t.Logf("DB connected")
}

func TestCreateUser_PSQL(t *testing.T) {
	var err error
	userId, err = psql.CreateUser("john")
	assert.Nil(t, err)
	fmt.Printf("User created %v \n", userId)
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

func TestGetTimeslot5_PSQL(t *testing.T) {
	timeslots, err := psql.GetTimeslots(userId, 0, 0)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(timeslots), "should have none timeslots")
}
