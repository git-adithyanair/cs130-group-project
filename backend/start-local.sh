#!/bin/sh

set -e # Exit on error

echo "run db migrations"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

export TWILIO_ACCOUNT_SID=AC0e098fc36218f0f7360b0a541b5c45da
export TWILIO_AUTH_TOKEN=acc402636baeac620f22316009456b7d

echo "start the app"
exec "$@" # Execute the CMD from Dockerfile by taking all args passed to CMD.