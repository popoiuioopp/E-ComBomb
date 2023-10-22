# Craft Cart

An e-commerce platform built with Go (using the Gin framework) for the backend and a frontend (currently under development).

## Table of Contents

- [Getting Started](#getting-started)
- [Prerequisites](#prerequisites)
- [Running the Project](#running-the-project)
- [Components](#components)
- [License](#license)

## Getting Started

These instructions will help you set up the project on your local machine for development and testing purposes.
Prerequisites

- Go
- Gin (This will be installed with the go mod commands provided below)

## Running the Project

### To run the project on docker

```
docker-compose up -d
```

### To run backend and frontend locally

- backend

```bash
cd backend
go run main.go
```

- frontend

```bash
cd frontend
bun install
bun dev
```

- database
  - Currently solely running database is still in progress.

## Components

### Backend:

Developed with Go and the Gin framework.  
Handles all the API endpoints, database connections, and business logic.

### Frontend:

Developed with NextJS and BunJS.  
Handles server side and client side rendered interactable page for users.

## License

This project is licensed under the MIT License - see the LICENSE.md file for details.
