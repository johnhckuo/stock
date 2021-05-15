package adaptors

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var mockDB DB
var userId int64
var timeId, timeId2, timeId3 int64

func TestCreateUser(t *testing.T) {
	mockDB = NewMock()
	var err error
	userId, err = mockDB.CreateUser("john")
	assert.Nil(t, err)
}

func TestCreateTimeslot(t *testing.T) {
	in, _ := time.Parse(time.RFC822, "01 Jan 15 10:00 UTC")
	out, _ := time.Parse(time.RFC822, "01 Jan 16 10:00 UTC")
	slot, err := mockDB.AddTimeslot(userId, in.Unix(), out.Unix())
	assert.Nil(t, err)
	timeId = slot.ID
}

func TestCreateTimeslot2(t *testing.T) {
	in, _ := time.Parse(time.RFC822, "01 Jan 16 10:00 UTC")
	out, _ := time.Parse(time.RFC822, "01 Jan 17 10:00 UTC")
	slot, err := mockDB.AddTimeslot(userId, in.Unix(), out.Unix())
	assert.Nil(t, err)
	timeId2 = slot.ID
}

func TestCreateTimeslot3(t *testing.T) {
	in, _ := time.Parse(time.RFC822, "01 Jan 17 10:00 UTC")
	out, _ := time.Parse(time.RFC822, "01 Jan 18 10:00 UTC")
	slot, err := mockDB.AddTimeslot(userId, in.Unix(), out.Unix())
	assert.Nil(t, err)
	timeId3 = slot.ID
}

// check get
func TestGetTimeslot1(t *testing.T) {
	timeslots, err := mockDB.GetTimeslots(userId, 0, 0)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(timeslots), "should have 3 timeslots")
}

//check before & after
func TestGetTimeslot2(t *testing.T) {
	before, _ := time.Parse(time.RFC822, "01 Jan 17 11:00 UTC")
	after, _ := time.Parse(time.RFC822, "01 Jan 16 11:00 UTC")
	timeslots, err := mockDB.GetTimeslots(userId, before.Unix(), after.Unix())
	assert.Nil(t, err)
	assert.Equal(t, 1, len(timeslots), "should have 1 timeslots")
}

//check before
func TestGetTimeslot3(t *testing.T) {
	before, _ := time.Parse(time.RFC822, "01 Jan 17 11:00 UTC")
	timeslots, err := mockDB.GetTimeslots(userId, before.Unix(), 0)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(timeslots), "should have 2 timeslots")
}

//check after
func TestGetTimeslot4(t *testing.T) {
	after, _ := time.Parse(time.RFC822, "01 Jan 17 11:00 UTC")
	timeslots, err := mockDB.GetTimeslots(userId, 0, after.Unix())
	assert.Nil(t, err)
	assert.Equal(t, 1, len(timeslots), "should have 1 timeslots")
}

func TestAddTimeslot(t *testing.T) {
	in, _ := time.Parse(time.RFC822, "01 Jan 15 20:00 UTC")
	out, _ := time.Parse(time.RFC822, "01 Jan 17 10:00 UTC")
	_, err := mockDB.AddTimeslot(userId, in.Unix(), out.Unix())
	assert.NotNil(t, err)
}

func TestDeleteTimeslot(t *testing.T) {
	success, err := mockDB.DeleteTimeslot(userId, timeId)
	assert.Nil(t, err)
	assert.Equal(t, true, success, "delete successfully")
}

func TestGetTimeslot5(t *testing.T) {
	timeslots, err := mockDB.GetTimeslots(userId, 0, 0)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(timeslots), "should have none timeslots")
}
