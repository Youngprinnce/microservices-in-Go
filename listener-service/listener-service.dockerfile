# Build a tiny docker image
FROM alpine:latest

RUN mkdir /app

# Copy the Pre-built binary file(listenerApp)
COPY listenerApp /app

# Command to run the executable
CMD ["/app/listenerApp"]
