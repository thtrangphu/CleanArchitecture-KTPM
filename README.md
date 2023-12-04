# CleanArchitecture-TKPM

CleanArchitecture-TKPM is a web application implementing clean architecture principles. 

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

What things you need to install the software and how to install them:
- [Docker](https://www.docker.com/)
- [Go](https://go.dev/)
- [Postman](https://www.postman.com/)
- Make [optional]
- Migrate DB


### Installing

A step-by-step series of examples that tell you how to get a development environment running:

1. Clone the repository:
   ```sh
   git clone https://github.com/thtrangphu/CleanArchitecture-TKPM.git

2. Start Docker and Set Up DB
  ```sh
  docker run --name postgres15 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine
  docker exec -it postgres15 psql
  docker exec -it postgres15 createdb --username=root --owner=root clean-architecture
  migrate  -path .\database\migrations\ -datebase "postgresql://root:password@localhost:5433/clean-architecture?sslmode=disable" -verbose up




3. Run
  ```sh
  go run cmd/main.go

4. Postman 

