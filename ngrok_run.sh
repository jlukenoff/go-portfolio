#! /usr/bin/env bash

#!/bin/bash

# Function to kill ngrok process
cleanup() {
  echo "Cleaning up..."
  kill $(jobs -p)
}

# Trap EXIT signal and call cleanup function
trap cleanup EXIT


# Run the go server
go run ./main.go > /dev/null &

# Start ngrok
ngrok http --domain="$NGROK_DOMAIN" "$PORT"

# Keep script running to maintain ngrok process
wait
