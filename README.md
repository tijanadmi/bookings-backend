# Room Reservation Backend API

his repository contains the backend API for a room reservation system.

The application manages room availability for specific periods and allows users to book rooms for desired dates. 

The backend is developed using the Go programming language, PostgreSQL for data storage, and Docker for database management. It also supports both gRPC and HTTP requests using gRPC-Gateway.

## Features

- Manage room availability
- Book rooms for a specific time period
- User session management
- Serve both gRPC and HTTP requests via gRPC-Gateway
- User authentication and authorization
- Secure API endpoints with JWT and PASETO access tokens
- Configuration management using Viper
- Robust unit tests with Testify
- Mocking database interactions for testing
- PostgreSQL database integration via Docker


## Technology Stack

* using [go1.22](https://tip.golang.org/doc/go1.22)
* using [gRPC](https://grpc.io/docs/languages/go/quickstart/) and [gRPC-Gateway](https://github.com/grpc-ecosystem/grpc-gateway) for API communication
* using [viper](https://github.com/spf13/viper) as a configuration solution
* using [PostgreSQL](https://www.postgresql.org/download/) for relational database management
* using [jwt-go](github.com/dgrijalva/jwt-go) to provide an implementation of JWT
* using [x/crypto](golang.org/x/crypto), Go Cryptography package 
* using [testify](https://github.com/stretchr/testify), for write robust unit tests 
* using [Gomock](https://github.com/golang/mock) for mocking database
* using [Docker](https://www.docker.com/products/docker-desktop/) for containerized PostgreSQL database


## Getting Started

### Prerequisites
-	Go (version go1.22.4 or higher)
-	PostgreSQL(managed via Docker)
-	Git
-   Docker (for running PostgreSQL)

### Installation

1.	**Clone the repository:**

```sh
git clone https://github.com/tijanadmi/bookings-backend.git
cd bookings-backend
```

2. **Install dependencies:**

```sh
go mod download
```

3. **Install Docker**
   
Instructions for installation are https://docs.docker.com/engine/install/

4. **Set up environment variables**

Create a `.env` file in the root directory of the project and add the necessary configuration variables.

```sh
ENVIRONMENT=development
DB_SOURCE=postgresql://root:root@localhost:5433/bookings?sslmode=disable
MIGRATION_URL=file://db/migration
HTTP_SERVER_ADDRESS=0.0.0.0:8080
GRPC_SERVER_ADDRESS=0.0.0.0:9090
TOKEN_SYMMETRIC_KEY=12345678901234567890123456789012
ACCESS_TOKEN_DURATION=15m
REFRESH_TOKEN_DURATION=24h
```

## Running the Application

1. **Start Posgres container:**

```sh
make postgres
```

2. **Run the server:**

```sh
make server
```

3. **Access the API:**

-   The HTTP API will be available at http://localhost:8080
-   The gRPC API will be available at http://localhost:9090

## Running Tests

To run tests, use the following command:

```sh
go test ./...
```

This will run all the tests in the project, including unit tests with mocked database interactions.

## API Documentation

Detailed documentation and descriptions of the API endpoints are available at:

http://localhost:8080/swagger/index.html

## QR Code for GitHub Repository




## License

[MIT](https://choosealicense.com/licenses/mit/)
