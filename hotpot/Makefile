new_mod_name ?= "meal"

generate-module:
	@echo "Generating new module..."
	go run cmd/template/main.go generate --module=$(new_mod_name)

install:
	@echo "Installing dependencies..."
	go install github.com/swaggo/swag/cmd/swag@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go get github.com/subosito/gotenv
	go mod tidy

generate-doc:
	@echo "Generating documentation..."
	swag init --generalInfo cmd/server/main.go

migration-create:
	   @echo "Creating new migration..."
	   migrate create -ext sql -dir cmd/db/migrations -format 20060102150405 $(name)
	   git add *

migrate-db:
	@echo "Running migration..."
	go run cmd/db/migrate.go