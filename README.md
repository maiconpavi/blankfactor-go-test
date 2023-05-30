# blankfactor-go-test

## The Test

### Double Booked

Senior Golang

When maintaining a calendar of events, it is important to know if an event overlaps with another event.

Given a sequence of events, each having a start and end time, write a program that will return the sequence of all pairs of overlapping events.

## The Solution

I decided to use SQLite as Database and Golang as programming language.
To get the pairs of overlapping events I used the following query:

```sql
SELECT e1.id AS id1,
       e1.title AS title1,
       e1.start_time AS start_time1,
       e1.end_time AS end_time1,
       e2.id AS id2,
       e2.title AS title2,
       e2.start_time AS start_time2,
       e2.end_time AS end_time2
FROM EVENTS e1
INNER JOIN EVENTS e2 ON e2.end_time >= e1.start_time
AND e1.end_time >= e2.start_time
AND e1.id != e2.id
AND e1.start_time < e2.start_time
```

## How to run

### Prerequisites

- Golang 1.20
- SQLite 3
- Make

### Run

Use `make run` to run the program.
Access the API at `http://localhost:8080/`.
The Swagger UI is available at `http://localhost:8080/swagger/index.html`.

### Build

Use `make build` to build the program.
 