# Use an official Node.js runtime as a parent image
FROM node:16

# Set the working directory in the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Install dependencies
RUN npm install

# Build your app for production
RUN npm run build

# Install a simple HTTP server to serve your static content
RUN npm install -g http-server

# Make port 8080 available to the world outside this container
EXPOSE 8080

# Serve the app
CMD ["http-server", "dist", "-p 8080"]
