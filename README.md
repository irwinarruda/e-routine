# e-Routine

```go

type Todo struct {
  id          number
  title       string
  description string
  weekDays   []number
  createdAt  time.Time
  updatedAt  time.Time
}

type TodoDay struct {
  id         number
  todo       Todo
  isFinished bool
  createdAt  time.Time
  updatedAt  time.Time
}

type Day struct {
  id       number
  weekDay  number
  todosDay []TodoDay
  createdAt  time.Time
  updatedAt  time.Time
}

```

To run the migrations, first you need to install some cmd apps.

```bash
# install
go install github.com/pressly/goose/v3/cmd/goose@latest
go install ./cmd

# run
cd ./server
# -u or --up for running the migrations
e-routine-cmd --up ./migrations
# -d or --down for removing the migrations
e-routine-cmd --down ./migrations
```

To run the server, firs install air

```bash
go install github.com/cosmtrek/air@latest
```

```bash
# start the server and watch changes
air bench
```
