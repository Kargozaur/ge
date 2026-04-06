DB_PATH=./local.db
MIGRATIONS_DIR=./migrations
DB=sqlite3

.PHONY: migrate-status migrate-up migrate-down migrate-create

migrate-status:
	goose -dir $(MIGRATIONS_DIR) $(DB) $(DB_PATH) status

migrate-up:
	goose -dir $(MIGRATIONS_DIR) $(DB) $(DB_PATH) up

migrate-down:
	goose -dir $(MIGRATIONS_DIR) $(DB) $(DB_PATH) down

migrate-create:
	@read -p "Enter migration name: " name; \
	goose -dir $(MIGRATIONS_DIR) create $$name sql
