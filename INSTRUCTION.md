# Dummy Timetable

## Design
### Capacity
Assuming this dummy timetable will have a Read/Write ratio of 10:1, meaning this system will expect more read request, and number of write requests(add timeslots) is comparatively low.

### Api interface
- GetTimeslots(userId, before, after int64) ([]*models.Timeslot, error)
- AddTimeslot(userId, startAt, endAt int64) (*models.Timeslot, error)
- DeleteTimeslot(userId, timeId int64) (bool, error)
- CreateUser(name string) (int64, error)

### DB Schema
```
User{
    id int64(index)
    name string
}

TimeSlot{
    id int64
    user_id int64(index)
    start_at int64(index)
    end_at int64(index)
}
```
### Choice of DB type
NA, because RDBMS is being chosen in advance

### Simple architecture
```
(client) -write request-> (api server) –update in DB-> (PostgreSQL)
(client) -read request-> (api server) -read from DB-> (PostgreSQL)
```
~~we use memcache with write through cache strategy since we know that the number of read request will be more than write request, so sacrificing the performance of write for the sake of higher reading speed is reasonable~~

I've decided to abandon cache here since the key of cache will be userid & value will be all the timeslots that belong to this user.
if a user holds lots of time slots and update it frequently, system memory might get overloaded. and the expiration time of each cache will be quite long because we can't restrict user from adding timeslot only in certain time period. 
Setting clustered index at both start_time & end_time in SQL will give us a huge boost for reading data & ranged query as well.

## Command

```bash
# build this application as docker image
make docker

# run locally at port 8080
make run

# generate http handler based on swagger.yaml
make swagger

# show api docs in browser
make docs

# generate api specs in Markdown
make spec

# start testing and output coverage
make unit_test
```

## Setup

```bash
# build application as docker image first
make docker

# spin up necessary containers
docker-compose up

# after seeing this log: `Serving codementordev hello johnhckuo at http://[::]:8080`
# then you are good to go!
curl --location --request POST 'http://127.0.0.1:8080/users/john'

# clean up
docker-compose down
```

## Test

1. You can run `make unit_test` for handler testing
1. Import `test/Timetable.postman_collection.json` and test api endpoint in postman

## Powered by
- [Go Swagger](https://goswagger.io/)
- [Testify](https://github.com/stretchr/testify)
- [GoDotenv](https://pkg.go.dev/github.com/joho/godotenv?utm_source=godoc)
- [Sqlx](https://jmoiron.github.io/sqlx/)