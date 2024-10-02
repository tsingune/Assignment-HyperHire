# Distributed File Storage Server

This project is a distributed file storage server built with **Golang** and **PostgreSQL**. The application allows you to upload files, split them into chunks, store them in a database, retrieve the files, and reassemble them upon download.

## Table of Contents
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Folder Structure](#folder-structure)
- [Installation](#installation)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Environment Variables](#environment-variables)
- [Future Enhancements](#future-enhancements)

## Features

1. **File Upload**:
   - Splits a file into multiple chunks and stores them in a PostgreSQL database using multithreading.
   - Returns a unique file ID for later retrieval.

2. **File Retrieval**:
   - Allows retrieving file data by file ID and returns it in base64 format.

3. **File Download**:
   - Reassembles file chunks from the database and returns the complete file.

4. **Dockerized Setup**:
   - The entire project can be built and run with a single `docker-compose` command.

## Tech Stack

- **Backend**: Go (Golang)
- **Database**: PostgreSQL
- **Docker**: For containerization of the application and database
- **Multithreading**: For parallel file uploads and downloads

---

## Folder Structure

```plaintext
├── go.mod
├── go.sum
├── cmd
│   └── fileStorageService
│       ├── Dockerfile
│       ├── docker-compose.yml
│       └── main.go
├── fileStorageService
│   ├── api.go
│   ├── middleware.go
│   ├── config.go
│   ├── repository.go
│   ├── service.go
│   ├── utils.go
│   ├── model.go
│   └── server.go
└── db.go
```

- `cmd/fileStorageService`: Contains the main entry point and Docker configuration.
- `fileStorageService`: Contains all business logic, models, and API handler code.
- `db.go`: Contains database connection logic.

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/tushar/Assignment-HyperHire.git
   cd Assignment-HyperHire
   ```

2. Ensure you have Docker and Docker Compose installed on your machine.

3. Build and run the application using Docker Compose:
   ```sh
   docker-compose up --build
   ```

## Running the Application

Once the application is running, it will be available at `http://localhost:8080`.

### Makefile Commands

This project includes a Makefile to simplify common development tasks. Here are the available commands:

- `make build`: Builds the Go application using Docker Compose.
- `make up`: Starts the services using Docker Compose.
- `make down`: Stops the running services.
- `make clean`: Cleans up Docker containers, volumes, and images.
- `make rebuild`: Rebuilds and restarts the service (equivalent to running `make down`, `make build`, and `make up` in sequence).

To use these commands, run `make <command>` in the terminal from the project root directory.

## API Endpoints

- **Upload File**: `POST /upload`
  - Uploads a file and returns a unique file ID.
  - Example:
    ```sh
    curl -F "file=@/path/to/your/file" http://localhost:8080/upload
    ```

- **Get Uploaded Files**: `GET /files`
  - Returns a list of all uploaded file IDs.
  - Example:
    ```sh
    curl http://localhost:8080/files
    ```

- **Download File**: `GET /download?id=<fileID>`
  - Downloads the file with the given file ID.
  - Example:
    ```sh
    curl -O http://localhost:8080/download?id=<fileID>
    ```

## Environment Variables

The application requires the following environment variables to be set:

- `DB_HOST`: Database host
- `DB_USER`: Database user
- `DB_PASSWORD`: Database password
- `DB_NAME`: Database name
- `DB_PORT`: Database port

These can be set in a `.env` file or directly in your environment.

