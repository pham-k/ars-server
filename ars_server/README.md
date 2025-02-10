# Project ars_server

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

Create DB container
```bash
make docker-run
```

Shutdown DB container
```bash
make docker-down
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```

migrate create
```bash
migrate create -seq -ext=.sql -dir=./migrations create_token_table 
```

migrate up
```bash
migrate -path=./migrations -database="postgres://ars_dev:ars_dev@localhost:5432/ars_dev?sslmode=disable" up 
```

migrate down
```bash
migrate -path=./migrations -database="postgres://ars_dev:ars_dev@localhost:5432/ars_dev?sslmode=disable" down
```

migrate version
```bash
migrate -path=./migrations -database="postgres://ars_dev:ars_dev@localhost:5432/ars_dev?sslmode=disable" version
```