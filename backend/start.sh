#!/bin/sh

set -e # Exit on error

echo "run db migrations"
source /app/app.env
echo "TWILIO_AUTH_TOKEN: $TWILIO_AUTH_TOKEN"

/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@" # Execute the CMD from Dockerfile by taking all args passed to CMD.