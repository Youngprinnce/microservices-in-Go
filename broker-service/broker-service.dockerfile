# The dockerfile that tells dockercompaose how to build the broker-service image

# this is the builder image (as builder is an alias)
FROM golang:1.22-alpine as builder    

# Set the Current Working Directory inside the container
RUN mkdir /app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . /app

# Set the Current Working Directory inside the container
WORKDIR /app

# Build the Go app (brokerApp) and output it to the /app directory
RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

# Optional test just to make sure it is executable
RUN chmod +x /app/brokerApp

#############################################################################

# This is a new docker image below

# Build a tiny docker image
FROM alpine:latest

RUN mkdir /app

# Copy the Pre-built binary file(builder) from the previous stage
COPY --from=builder /app/brokerApp /app

# Command to run the executable
CMD ["/app/brokerApp"]
