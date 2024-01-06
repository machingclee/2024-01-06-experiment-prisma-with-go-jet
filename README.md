# 2024-01-06-experiment-prisma-with-go-jet

This is an experiment to integrate prisma and go-jet

# To run the Project

- Docker-compose up
- `go get -u github.com/go-jet/jet/v2`
- `go install github.com/go-jet/jet/v2/cmd/jet@latest`
- `yarn`
- `yarn migrate`
- Add a record in the table (student)
- `make from-db`
- `make run`
