# Dummy Timetable

## Design
### Capacity
Assuming this dummy timetable will have a Read/Write ratio of 10:1, meaning this system will expect more read request, and number of write requests(add timeslots) is comparatively low.

### Api interface
- func getTimeSlots(userId string, endAt int64) []TimeSlot
- func addTimeSlot(userId string, newTime TimeSlot) TimeSlot
- func deleteTimeSlot(userId string, timeId int) bool
- func createUser(name string) int64

### DB Schema

User{
    id int64
    name string
}

TimeSlot{
    id int64
    user_id int64(index)
    start_at int64(index)
    end_at int64(index)
}

### Choice of DB type
NA, because RDBMS is being chosen in advance

### Simple architecture
(client) -write request-> (api server) –update in DB-> (PostgreSQL)
(client) -read request-> (api server) -read from DB-> (PostgreSQL)

~~we use memcache with write through cache strategy since we know that the number of read request will be more than write request, so sacrificing the performance of write for the sake of higher reading speed is reasonable~~
I've decided to abandon cache here since the key of cache will be userid & value will be all the timeslots that belong to this user.
if a user holds lots of time slots and update it frequently, system memory might get overloaded. and the expiration time of each cache will be quite long because we can't restrict user from adding timeslot only in certain time period. 
Setting clustered index at both start_time & end_time in SQL will give us a huge boost for reading data & ranged query as well.

## Setup

## Run