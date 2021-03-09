# Dummy Timetable
> You need to implement a timetable API follow the below requirements


This is an internal Restful API service, the client will send HTTP requests to this service, and it should respond with an expected response or understandable error message.


## Business rules
* `start_at` must greater than `now`.
* `end_at` must greater than `start_at`
* Timeslots from the same user are not allowed to overlap
i.e., 

```
existing slot: 100 to 200

slot0: 10 to 30 => ok
slot1: 50 to 100 => ok
slot2: 90 to 120 => no
slot3: 150 to 180 => no
slot4: 10 to 250 => no
slot from different user: 50 to 120 => ok
```

## System requirements
* The timeslot data should be store in RDBMS(i.e., PostgreSQL, MySQL ...)


## API Spec

### Query timetable by user_id
Query user's time slots. The result should order by `start_at`.
HTTP Method: `GET`
Path: `/users/:user_id/time-slots`

### Response Body
```json
[
  {
    "start_at": 1615189066,
    "end_at": 1615192666
  },
  {
    "start_at": 1615199866,
    "end_at": 1615203466
  },
  {
    "start_at": 1615208866,
    "end_at": 1615212466
  }
]
```

---

### Create time slot by user_id
Create a time slot given user_id, and it should follow the rule mentioned.
HTTP Method: `POST`
Path: `/users/:user_id/time-slots`

#### Request Body
`start_at` and `end_at` are Unix epoch(the number of seconds that have elapsed since January 1, 1970 (midnight UTC/GMT))

"`json
{
  "start_at": 1615189066,
  "end_at": 1615192666
}
```

---

### Delete time slot by user_id and time_slot_id
Delete a time slot
HTTP Method: `DELETE`
Path: `/users/:user_id/time-slots/:time_slot_id`

## Submission
Your submission should have clear instructions on how to run your code. Please put the instruction in a file named `INSTRUCTION.md`. You should provide a `docker-compose.yml` file, so the tester can run `docker-compose up -d' to test your application.


## Evaluation
Your submission will be evaluated by:
* Correctness of the API implementation.
* Code quality and architecture.
* Good programming practices such as clean code, clear comments, tests.
