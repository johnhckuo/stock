


# TimeSlot Manager
TimeSlot manager
  

## Informations

### Version

1.0.0

## Content negotiation

### URI Schemes
  * http

### Consumes
  * application/json
  * application/timeslot.v1+json

### Produces
  * application/json
  * application/timeslot.v1+json

## All endpoints

###  timeslot

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| POST | /users/{user_id}/time-slots | [add timeslot](#add-timeslot) |  |
| POST | /users/{username} | [add user](#add-user) |  |
| DELETE | /users/{user_id}/time-slots/{time_slot_id} | [destroy timeslot](#destroy-timeslot) |  |
| GET | /users/{user_id}/time-slots | [get timeslot](#get-timeslot) |  |
  


## Paths

### <span id="add-timeslot"></span> add timeslot (*addTimeslot*)

```
POST /users/{user_id}/time-slots
```

Add new time slot and return time slot id

#### Consumes
  * application/json

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| user_id | `path` | int64 (formatted integer) | `int64` |  | ✓ |  |  |
| body | `body` | [Timeslot](#timeslot) | `models.Timeslot` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#add-timeslot-201) | Created | Created |  | [schema](#add-timeslot-201-schema) |
| [default](#add-timeslot-default) | | error |  | [schema](#add-timeslot-default-schema) |

#### Responses


##### <span id="add-timeslot-201"></span> 201 - Created
Status: Created

###### <span id="add-timeslot-201-schema"></span> Schema
   
  

[Timeslot](#timeslot)

##### <span id="add-timeslot-default"></span> Default Response
error

###### <span id="add-timeslot-default-schema"></span> Schema

  

[Error](#error)

### <span id="add-user"></span> add user (*addUser*)

```
POST /users/{username}
```

Add new user

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| username | `path` | string | `string` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#add-user-201) | Created | Created |  | [schema](#add-user-201-schema) |
| [default](#add-user-default) | | error |  | [schema](#add-user-default-schema) |

#### Responses


##### <span id="add-user-201"></span> 201 - Created
Status: Created

###### <span id="add-user-201-schema"></span> Schema
   
  

[User](#user)

##### <span id="add-user-default"></span> Default Response
error

###### <span id="add-user-default-schema"></span> Schema

  

[Error](#error)

### <span id="destroy-timeslot"></span> destroy timeslot (*destroyTimeslot*)

```
DELETE /users/{user_id}/time-slots/{time_slot_id}
```

Delete time slot according to given id

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| time_slot_id | `path` | int64 (formatted integer) | `int64` |  | ✓ |  |  |
| user_id | `path` | int64 (formatted integer) | `int64` |  | ✓ |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [204](#destroy-timeslot-204) | No Content | Deleted |  | [schema](#destroy-timeslot-204-schema) |
| [default](#destroy-timeslot-default) | | error |  | [schema](#destroy-timeslot-default-schema) |

#### Responses


##### <span id="destroy-timeslot-204"></span> 204 - Deleted
Status: No Content

###### <span id="destroy-timeslot-204-schema"></span> Schema

##### <span id="destroy-timeslot-default"></span> Default Response
error

###### <span id="destroy-timeslot-default-schema"></span> Schema

  

[Error](#error)

### <span id="get-timeslot"></span> get timeslot (*getTimeslot*)

```
GET /users/{user_id}/time-slots
```

Retrieve time slots based on userId & endAt filter

#### Produces
  * application/json

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| user_id | `path` | int64 (formatted integer) | `int64` |  | ✓ |  |  |
| after_timestamp | `query` | int64 (formatted integer) | `int64` |  |  |  |  |
| before_timestamp | `query` | int64 (formatted integer) | `int64` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#get-timeslot-200) | OK | get available timeslots for the given user |  | [schema](#get-timeslot-200-schema) |
| [default](#get-timeslot-default) | | generic error response |  | [schema](#get-timeslot-default-schema) |

#### Responses


##### <span id="get-timeslot-200"></span> 200 - get available timeslots for the given user
Status: OK

###### <span id="get-timeslot-200-schema"></span> Schema
   
  

[][Timeslot](#timeslot)

##### <span id="get-timeslot-default"></span> Default Response
generic error response

###### <span id="get-timeslot-default-schema"></span> Schema

  

[Error](#error)

## Models

### <span id="error"></span> Error


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| code | int64 (formatted integer)| `int64` |  | |  |  |
| message | string| `string` | ✓ | |  |  |



### <span id="timeslot"></span> Timeslot


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| endAt | int64 (formatted integer)| `int64` | ✓ | |  |  |
| id | int64 (formatted integer)| `int64` |  | |  |  |
| startAt | int64 (formatted integer)| `int64` | ✓ | |  |  |
| userId | int64 (formatted integer)| `int64` |  | |  |  |



### <span id="user"></span> User


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| id | int64 (formatted integer)| `int64` |  | |  |  |
| name | string| `string` | ✓ | |  |  |


