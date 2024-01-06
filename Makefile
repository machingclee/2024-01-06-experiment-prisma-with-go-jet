DB_URL = postgresql://pguser:pguser@localhost:5432/udemy

from-db:
	jet -dsn=$(DB_URL)?sslmode=disable -schema=public -path=./.gen -ignore-tables=_prisma_migrations
run:
	go run ./cmd/api/*.go