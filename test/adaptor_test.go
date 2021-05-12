package test

import(
    "time"
    "testing"
    "codementordev/hello-johnhckuo/internal/adaptors"
)

var mockDB adaptors.DB
var userId int64
var timeId int64

func TestCreateUser(t *testing.T){
    mockDB = adaptors.NewMock()
    var err error
    userId, err = mockDB.CreateUser("john")
    if err != nil{
        t.Errorf("Error when creating user %v", err)
    }
    t.Logf("User created %v", userId)
}

func TestCreateTimeslot(t *testing.T){
	in, _ := time.Parse(time.RFC822, "01 Jan 15 10:00 UTC")
	out, _ := time.Parse(time.RFC822, "01 Jan 16 10:00 UTC")

    slot, err := mockDB.AddTimeslot(userId, in.Unix(), out.Unix())
    if err != nil{
        t.Errorf("Error when creating timeslot %v", err)
    }
    t.Logf("Timeslot created %v", slot)
    timeId = slot.ID
}

// check get
func TestGetTimeslot1(t *testing.T){
    id, err := mockDB.GetTimeslots(userId, 0, 0)
    if err != nil{
        t.Errorf("Error when creating user %v", err)
    }
    t.Logf("User created %v", id)
}

//check before & after
func TestGetTimeslot2(t *testing.T){
    
    start, _ := time.Parse(time.RFC822, "01 Jan 15 10:00 UTC")
	end, _ := time.Parse(time.RFC822, "01 Jan 16 10:00 UTC")
    timeslots, err := mockDB.GetTimeslots(userId, start.Unix(), end.Unix())
    if err != nil{
        t.Errorf("Error when getting timeslot %v", err)
    }
    t.Logf("Timeslot retrieved %v", timeslots)
}

//check before
func TestGetTimeslot3(t *testing.T){
    
    before, _ := time.Parse(time.RFC822, "01 Jan 15 11:00 UTC")
    timeslots, err := mockDB.GetTimeslots(userId, before.Unix(), 0)
    if err != nil{
        t.Errorf("Error when getting timeslot %v", err)
    }
    t.Logf("Timeslot retrieved %v", timeslots)
}
//check after
func TestGetTimeslot4(t *testing.T){
    
    after, _ := time.Parse(time.RFC822, "01 Jan 15 9:00 UTC")
    timeslots, err := mockDB.GetTimeslots(userId, 0, after.Unix())
    if err != nil{
        t.Errorf("Error when getting timeslot %v", err)
    }
    t.Logf("Timeslot retrieved %v", timeslots)
}

func TestAddTimeslot(t *testing.T){
    in, _ := time.Parse(time.RFC822, "01 Jan 15 20:00 UTC")
	out, _ := time.Parse(time.RFC822, "01 Jan 17 10:00 UTC")

    slot, err := mockDB.AddTimeslot(userId, in.Unix(), out.Unix())
    if err != nil{
        t.Logf("Timeslot created failed as expected: %v", err)  
    }
    t.Errorf("Timeslot is still created even overlapped, %v", slot)
}    

func TestDeleteTimeslot(t *testing.T){
    success, err := mockDB.DeleteTimeslot(userId, timeId)
    if err != nil{
        t.Errorf("Error when deleting timeslot %v", err)
    }

    if success{
        t.Logf("timeslot deleted")
    }
}


func TestGetTimeslot5(t *testing.T){
    timeslots, err := mockDB.GetTimeslots(userId, 0, 0)
    if err != nil{
        t.Errorf("Error when getting timeslot %v", err)
    }
    t.Logf("Timeslot retrieved %v", timeslots)
}