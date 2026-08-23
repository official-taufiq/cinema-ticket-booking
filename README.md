# Cinema Booking

A simple cinema seat-booking application built with Go and Redis.

## Features

- Browse available movies and seats
- Temporarily hold seats
- Confirm or release bookings
- Track seat availability with Redis

## Run

Requirements: Go 1.25+ and Docker.

```bash
docker compose up -d
go run ./cmd
```

Open `http://localhost:8080/` in your browser.

## Test

```bash
go test ./...
```
