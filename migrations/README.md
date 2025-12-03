# Migrations
## Description

This is migrations service to do migrations and setup database.


## How to run

you need to have a golang-migrate

```
brew install golang-migrate

migrate -help

// Create migration 
migrate create -ext sql -dir . -seq init_schema
```

usage of env

```
# Use dev config
make migrations-up ENV=dev

# Use prod config
make migrations-up ENV=prod
```

commands

```
# Database Migrations (Makefile targets)

This document explains the Makefile targets used to manage database migrations.

## Overview

The Makefile provides a small set of migration-related commands and configuration variables:

- ENV: Environment selection (`dev` or `prod`). Default: `dev`.
- ENV_FILE: Path to the environment file that contains DB connection variables. Default: `config/.$(ENV).env`.
- MIGRATE: Migration CLI command (default `migrate`).
- MIGRATIONS_DIR: Directory where migration SQL files live. Default: `./migrations`.

All migration commands that apply changes to the database (up, down, version) rely on the environment file to provide `DB_URL`. They will fail early if the expected env file is missing.

## Targets

- migrations-i
    - Description: Installs golang-migrate via Homebrew on macOS if not already installed.
    - Usage: `make migrations-i`
    - Notes: macOS/Homebrew specific. On other platforms install the migrate CLI using your platform's package manager or download a binary.

- migrations-new
    - Description: Create a new SQL migration file pair (up/down) using the migrate CLI.
    - Usage: `make migrations-new name=add_users_table [ENV=dev]`
    - Notes: The `name` parameter is required and provides the migration name suffix. Files are created in `MIGRATIONS_DIR`.

- migrations-up
    - Description: Apply all pending migrations (migrate up).
    - Usage: `make migrations-up [ENV=dev]`
    - Requirements: The environment file (`ENV_FILE`) must exist and define `DB_URL`. The Makefile sources it before running the migrate command.

- migrations-down
    - Description: Roll back a single migration (migrate down 1).
    - Usage: `make migrations-down [ENV=dev]`
    - Requirements: The environment file (`ENV_FILE`) must exist and define `DB_URL`.

- migrations-version
    - Description: Show the current migration version applied to the database.
    - Usage: `make migrations-version [ENV=dev]`
    - Requirements: The environment file (`ENV_FILE`) must exist and define `DB_URL`.

## Examples

- Create a new migration:
    - `make migrations-new name=add_orders_table`
- Apply migrations for production (ensure `config/.prod.env` exists and contains DB_URL):
    - `make migrations-up ENV=prod`
- Roll back one migration:
    - `make migrations-down`
- Show current migration version:
    - `make migrations-version`

## Notes

- Ensure the environment file (default `config/.$(ENV).env`) is present and contains a `DB_URL` value before running commands that touch the database.
- The Makefile uses `set -a; . <envfile>; set +a` to export variables from the env file into the shell environment for the migrate command.
- If you prefer a different migrate binary path or a different migrations directory, set the `MIGRATE` and `MIGRATIONS_DIR` variables when invoking `make`, e.g.:
    - `make migrations-up MIGRATE=/usr/local/bin/migrate MIGRATIONS_DIR=./db/migrations`
```

