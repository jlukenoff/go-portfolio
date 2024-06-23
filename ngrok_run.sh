#! /usr/bin/env bash

PORT=8080
NGROK_DOMAIN="sought-toucan-together.ngrok-free.app"

# Function to kill ngrok process
cleanup() {
  echo "Cleaning up..."
  kill "$(jobs -p)"
}

# Trap EXIT signal and call cleanup function
trap cleanup EXIT


# Run the go server
go run ./main.go > /dev/null &

# Start ngrok
ngrok http --domain="$NGROK_DOMAIN" "$PORT"

# Keep script running to maintain ngrok process
wait
