#! /usr/bin/env bash

#!/bin/bash

# Function to kill ngrok process
cleanup() {
  echo "Cleaning up..."
  kill $(jobs -p)
}

# Trap EXIT signal and call cleanup function
trap cleanup EXIT

# Start ngrok
ngrok http --domain="$NGROK_DOMAIN" "$PORT" > /dev/null &

# Run the go server

go run ./main.go

# Keep script running to maintain ngrok process
wait
