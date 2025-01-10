# Go Microservices Project (`go-micro`)

## Overview

`go-micro` is a collection of small, self-contained, and loosely coupled microservices designed to communicate with one another through various methods such as REST APIs, gRPC, RPC, and message queues (using AMQP - Advanced Message Queuing Protocol). The project also includes a simple front-end application.

### Microservices Included:

1. **Front-End Service**  
   Displays web pages and serves as the user interface.

2. **Authentication Service**  
   Manages user authentication, backed by a PostgreSQL database.

3. **Logging Service**  
   Logs data to a MongoDB database.

4. **Listener Service**  
   Consumes messages from RabbitMQ and acts upon them.

5. **Broker Service**  
   Acts as an optional single point of entry to the microservice cluster.

6. **Mail Service**  
   Accepts JSON payloads, formats them into emails, and sends them.

### Why Go?  
Go (Golang) is chosen for its efficiency and simplicity in building distributed web applications, making it particularly suited for this project.

## Prerequisites

1. **Docker**  
   Install Docker from [Docker Desktop](https://www.docker.com/products/docker-desktop).

2. **Make** (for macOS)  
   Install Make using Homebrew:  
   ```bash
   brew install make
   ```

3. **Git**  
   Ensure Git is installed to clone the repository.

## Getting Started

1. **Clone the Repository**  
   ```bash
   git clone https://github.com/Youngprinnce/microservices-in-Go.git
   ```
   Navigate to the project folder:  
   ```bash
   cd microservices-in-Go
   ```

2. **Start Docker**  
   Ensure Docker is running on your machine.

3. **Build the Microservices**  
   Use the `Makefile` to build all the services:  
   ```bash
   make up_build
   ```

4. **Start the Front-End Application**  
   ```bash
   make start
   ```

5. **Access the Application**  
   Open your browser and navigate to:  
   ```
   http://localhost:80
   ```

## Logging Options

The Logging Service supports three methods for logging items to MongoDB:
1. **JSON Request** (REST API)
2. **RabbitMQ** (Message Queue)
3. **RPC** (Remote Procedure Call)

By default, the Broker Service uses **RPC** for logging. Other methods are commented out in the source code. You can uncomment the desired method in the `broker-service` code to use it.

## Project Features

### Microservice Communication
The microservices interact through multiple protocols:
- **REST API**: For standard web-based interactions.
- **gRPC**: For high-performance communication.
- **RPC**: For remote procedure calls.
- **AMQP**: For messaging via RabbitMQ.

### Containerization
All services are containerized using Docker to ensure portability and consistent deployments.

### Makefile Integration
A Makefile is included for streamlined commands:
- `make up_build`: Builds all services individually.
- `make start`: Builds and starts the front-end application.

## Technologies Used

- **Go (Golang)**: The programming language used to build all microservices.
- **PostgreSQL**: Database for the Authentication Service.
- **MongoDB**: Database for the Logging Service.
- **RabbitMQ**: Message broker for inter-service communication.
- **Docker**: For containerizing the services.
- **Make**: For automating the build and deployment process.

## Contributing

Contributions are welcome! Feel free to fork the repository and submit a pull request with your improvements or fixes.

## License

This project is licensed under the [MIT License](LICENSE).

---

For any questions or issues, please raise an issue on the repository or contact [Youngprinnce](https://github.com/Youngprinnce).

