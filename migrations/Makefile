# ---------- Config ----------

ENV            ?= dev                      # dev or prod
ENV_FILE       := config/.$(ENV).env
MIGRATE        ?= migrate
MIGRATIONS_DIR ?= ./migrations

# ---------- Migrations ----------

migrations-i: ## Install golang-migrate via Homebrew (macOS)
	brew list golang-migrate >/dev/null 2>&1 || brew install golang-migrate

# Usage: make migrations-new name=add_users_table ENV=dev
migrations-new: ## Create new migration (make migrations-new name=add_users_table)
	@test -n "$(name)" || (echo ">>> Please provide migration name: make migrations-new name=add_users_table"; exit 1)
	$(MIGRATE) create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)

# Usage: make migrations-up ENV=dev
migrations-up: ## Migrate all the way up
	@test -f "$(ENV_FILE)" || (echo ">>> Env file $(ENV_FILE) not found"; exit 1)
	set -a; . $(ENV_FILE); set +a; \
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$$DB_URL" up

migrations-down: ## Roll back one migration
	@test -f "$(ENV_FILE)" || (echo ">>> Env file $(ENV_FILE) not found"; exit 1)
	set -a; . $(ENV_FILE); set +a; \
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$$DB_URL" down 1

migrations-version: ## Show current migration version
	@test -f "$(ENV_FILE)" || (echo ">>> Env file $(ENV_FILE) not found"; exit 1)
	set -a; . $(ENV_FILE); set +a; \
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "$$DB_URL" version

.PHONY: migrations-i migrations-new migrations-up migrations-down migrations-version
