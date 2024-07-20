#!/bin/bash
export DOCKER_BUILDKIT=1

# Name of your Docker application
APP_NAME="vuetify-app"

# Stop the currently running container
docker stop $APP_NAME

# Remove the stopped container
docker rm $APP_NAME

# Build the new Docker image
docker build -t $APP_NAME .

# Run the new Docker container
docker run -d --name $APP_NAME -p 8081:8080 $APP_NAME

cd ..
rm -rf ./backend