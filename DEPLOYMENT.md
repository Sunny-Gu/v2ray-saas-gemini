# Deployment Guide

This document provides instructions on how to build and run the entire V2Ray SaaS application stack using Docker and Docker Compose.

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

Ensure that both Docker and Docker Compose are installed on your system and that the Docker daemon is running.

## Running the Application

The entire application stack, including all Go microservices, the MySQL database, and the Redis cache, is defined in the `docker-compose.yml` file.

To build and run all services in detached mode (in the background), navigate to the project's root directory and execute the following command:

```bash
docker-compose up --build -d
```

### Command Breakdown

- `docker-compose up`: This is the standard command to start the services defined in the `docker-compose.yml` file.
- `--build`: This flag forces Docker Compose to build the images for our Go services from their respective `Dockerfile`s before starting the containers. You should use this flag the first time you run the command or whenever you make changes to the source code.
- `-d`: This flag runs the containers in detached mode, meaning they will run in the background and you will not see their logs directly in your terminal.

## Verifying the Deployment

After running the command, you can check the status of the running containers:

```bash
docker-compose ps
```

You should see all services (`db`, `redis`, `portal-api`, `admin-api`, `node-service`, `node-agent`) with a `State` of `Up`.

You can also view the logs for a specific service:

```bash
# Example: View logs for the portal-api
docker-compose logs -f portal-api
```

## Accessing the Services

Once the containers are running, the services will be accessible at the following default ports on your local machine:

- **Portal API**: `http://localhost:8080`
- **Admin API**: `http://localhost:8081`
- **Node Service**: `http://localhost:8082`
- **Node Agent**: `http://localhost:8083`
- **MySQL Database**: `localhost:3306`
- **Redis**: `localhost:6379`

## Stopping the Application

To stop all running services, use the following command:

```bash
docker-compose down
```

If you also want to remove the persistent data volume for the database (this will delete all your data), you can add the `-v` flag:

```bash
docker-compose down -v
```
