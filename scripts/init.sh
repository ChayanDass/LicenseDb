#!/bin/bash
set -e

# Default DB connection values
DB_HOST="${DB_HOST:-localhost}"
DB_PORT="${DB_PORT:-5432}"
DB_NAME="${DB_NAME:-licensedb}"
DB_USER="${DB_USER:-fossy}"
DB_PASSWORD="${DB_PASSWORD:-fossy}"

# Default user details
USER_NAME="fossy"
DISPLAY_NAME="Fossy Admin"
USER_EMAIL="daschayan8837@gmail.com"
USER_LEVEL="SUPER_ADMIN"
USER_PASSWORD="fossy"
ACTIVE="true"

# Determine root path (Docker vs local)
if [ -d "/app" ]; then
  echo "Running inside Docker"
  MIGRATION_PATH="/app/migrations"
  MIGRATE_BIN="/app/migrate"
else
  echo "Running locally"
  MIGRATION_PATH="$(pwd)/pkg/db/migrations"
  MIGRATE_BIN="/usr/local/bin/migrate"
fi

# Drop and recreate database and user (only locally for safety)
if [ ! -d "/app" ]; then
  echo "Resetting database and user (local only)..."

  if command -v sudo &> /dev/null; then
    echo "Dropping database and user if they exist..."
    sudo -u postgres psql <<EOF
DROP DATABASE IF EXISTS $DB_NAME;
DROP USER IF EXISTS $DB_USER;
CREATE USER $DB_USER WITH PASSWORD '$DB_PASSWORD';
CREATE DATABASE $DB_NAME OWNER $DB_USER;
GRANT ALL PRIVILEGES ON DATABASE $DB_NAME TO $DB_USER;
EOF
    echo "Database and user reset."
  else
    echo "Cannot reset database/user — 'sudo' not found."
  fi
else
  echo "Inside Docker — skipping DB and user reset."
fi


# Wait for PostgreSQL to be ready
echo "Waiting for PostgreSQL at $DB_HOST:$DB_PORT to become available..."
until PGPASSWORD="$DB_PASSWORD" psql -h "$DB_HOST" -U "$DB_USER" -p "$DB_PORT" -d "$DB_NAME" -c '\q' 2>/dev/null; do
  echo "PostgreSQL is unavailable - retrying in 2 seconds..."
  sleep 2
done
echo "PostgreSQL is ready."

# Download migrate binary if missing
if ! command -v migrate &> /dev/null && [ ! -f "$MIGRATE_BIN" ]; then
  echo "Installing migrate binary locally..."
  curl -sL https://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.tar.gz | tar xz

  mv migrate "$MIGRATE_BIN"
  chmod +x "$MIGRATE_BIN"
fi

# Run migrations
echo "Running migrations from: $MIGRATION_PATH"
"$MIGRATE_BIN" -path "$MIGRATION_PATH" -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" up

# Insert default user
echo "Inserting default admin user..."
PGPASSWORD="$DB_PASSWORD" psql -h "$DB_HOST" -p "$DB_PORT" -d "$DB_NAME" -U "$DB_USER" <<EOF
INSERT INTO users (user_name, display_name, user_email, user_level, user_password)
VALUES ('$USER_NAME', '$DISPLAY_NAME', '$USER_EMAIL', '$USER_LEVEL', '$USER_PASSWORD')
ON CONFLICT (user_name) DO NOTHING;
EOF
echo "User '$USER_NAME' created"

# Grant CREATEDB (only locally)
if [ ! -d "/app" ]; then
  echo "Granting CREATEDB to '$DB_USER' (local only)..."
  if command -v sudo &> /dev/null; then
    sudo -u postgres psql -c "ALTER USER $DB_USER WITH CREATEDB;"
    echo "'$DB_USER' granted CREATEDB."
  else
    echo "Cannot grant CREATEDB — 'sudo' not found."
  fi
else
  echo "Inside Docker — skipping ALTER USER"
fi
