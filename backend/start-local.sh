#!/bin/sh

set -e # Exit on error

echo "run db migrations"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@" # Execute the CMD from Dockerfile by taking all args passed to CMD.