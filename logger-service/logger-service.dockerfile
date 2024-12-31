# Build a tiny docker image
FROM alpine:latest

RUN mkdir /app

# Copy the Pre-built binary file(loggerApp)
COPY loggerApp /app

# Command to run the executable
CMD ["/app/loggerApp"]
