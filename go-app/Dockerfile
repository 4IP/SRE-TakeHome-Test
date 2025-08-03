# Use the official Go image as a parent image
FROM golang:1.19-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY go.mod ./
COPY go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the container's workspace
COPY . .

# Build the Go app
RUN go build -o /go-hello-app .

# Expose port 8081 to the outside world
EXPOSE 8081

# Command to run the executable
CMD [ "/go-hello-app" ]
