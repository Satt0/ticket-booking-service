# Define the migration directory
MIGRATION_DIR=./migrations

# Use the DB_URL from the environment or default to an empty string if not set
include .env
export $(shell sed 's/=.*//' .env)
dev: 
	swag init && go run main.go api:launch
# Target to create a new migration
create-migration:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq $$name

# Target to apply all up migrations
# using sql-migrate
migrate-up:
	migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" up

# Target to apply a single down migration (undo last migration)
migrate-down:
	migrate -path $(MIGRATION_DIR) -database "$(DB_URL)" down 1

# docker
docker-up: 
	docker compose up -d --remove-orphans

docker-build: 
	docker build . -t booking:1.0.0