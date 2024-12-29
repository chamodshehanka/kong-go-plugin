#!/bin/bash

# Start the database container
docker compose up -d kong-database

# Wait for the database to be ready
echo "Waiting for the database to be ready..."
sleep 10

# Run the migrations bootstrap command
docker compose run --rm kong kong migrations bootstrap

# Start the Kong container
docker compose up -d kong

# Start the Konga container
docker compose up -d konga

echo "Kong has been started successfully."