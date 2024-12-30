# Build a tiny docker image
FROM alpine:latest

RUN mkdir /app

# Copy the Pre-built binary file(authApp)
COPY authApp /app

# Command to run the executable
CMD ["/app/authApp"]
