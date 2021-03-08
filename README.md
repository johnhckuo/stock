# Dummy Timetable
> You need to implement a timetable API follow the below requirements


This is an internal Restful API service, client will send HTTP request to this service and it should respond with expected response or understandable error message.

## Rules
* The timeslot data should be store in RDBMS(i.e., PostgreSQL, MySQL ...)
* `start_at` must greater than `now`
* `end_at` must greater than `start_at`
* Timeslots from same user are not allow to overlap
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


## API Spec


### Query timetable by user_id
Query user's time slots. The result should order by `start_at`.

API Path: `GET /users/:user_id/timeslots`

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


### Create timetable by user_id
Create a time slot given user_id, it should follow the rule mentioned.

API Path: `POST /users/:user_id/timeslots`

#### Request Body
`start_at` and `end_at` are Unix epoch(the number of seconds that have elapsed since January 1, 1970 (midnight UTC/GMT))

```json
{
  "start_at": 1615189066,
  "end_at": 1615192666
}
```


### Delete timetable by user_id and time_slot_id
Delete a time slot
API Path: `DELETE /users/:user_id/timeslots/:time_slot_id`

## Submission
Please

Your submission should have clear instructions on how to run your code. Please put the instruction in a file named `INSTRUCTION.md`. You should provide a `docker-compose.yml` file, so tester can simply run `docker-compose up -d` to test your application.


## Evaluation
Your submission will be evaluated by:
* Correctness of the API implementation.
* Code quality and architecture.
* Good programming practices such as clean code, clear comments, tests.
