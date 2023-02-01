########### BUILD STAGE ###########

# Choose the base image.
FROM golang:1.18-alpine AS builder

# Set the working directory.
WORKDIR /app

# Copy the source from the current directory to the working directory inside the container.
COPY . .

# Build the Go app, install curl and migrate command.
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

########### RUN STAGE ###########

# Choose the base image.
FROM alpine:3.13

# Set the working directory.
WORKDIR /app

# Copy the built binary and migrate binary from the previous stage, as well as the migration 
# files and environment vars. Also copy start script which starts the app and runs the migrations.
# Finally, also copy the wait-for script which waits for the database to be ready before connecting
# the app to it.
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY db/migration /app/migration
COPY app.env .
COPY start.sh .
COPY start-local.sh .
COPY wait-for.sh .

# Expose port 8080 to the outside world.
EXPOSE 8080

# Command to run the executable.
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]