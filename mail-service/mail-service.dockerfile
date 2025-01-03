# Build a tiny docker image
FROM alpine:latest

RUN mkdir /app

# Copy the Pre-built binary file(mailerApp)
COPY mailerApp /app
COPY templates /templates

# Command to run the executable
CMD ["/app/mailerApp"]
