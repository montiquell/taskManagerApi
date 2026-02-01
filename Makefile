include .env
export 

migrations-up:
	migrate -path ./internal/migrations -database $(DATABASE_URL) up

migrations-down:
	migrate -path ./internal/migrations -database $(DATABASE_URL) down
