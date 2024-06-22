#! /usr/bin/env bash

NGROK_DOMAIN='sought-toucan-together.ngrok-free.app'
PORT=80

# Start ngrok
ngrok http --domain="$NGROK_DOMAIN" "$PORT" > /dev/null &

# Start the go server
go run main.go

