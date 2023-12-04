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

3. Run App
   ```sh
   go run cmd/main.go

4. Using Postman 
**API Endpoints**
The application provides the following RESTful endpoints:

_GET /profile: Retrieve the profile of the current user.
UPDATE /profile: Update the profile details of the current user.
POST /signin: Authenticate a user and create a new session.
POST /signup: Register a new user._

**1. POST /signup: Register a new user.**
<img width="660" alt="image" src="https://github.com/thtrangphu/CleanArchitecture-TKPM/assets/76843467/7003d147-8e12-4950-9f58-b313b0081b87">

**2. UPDATE /profile: Update the profile details of the current user.**
<img width="674" alt="image" src="https://github.com/thtrangphu/CleanArchitecture-TKPM/assets/76843467/52863979-96f0-4f7f-9304-3b7a173b2de8">

**3. POST /signin: Authenticate a user and create a new session.**
<img width="638" alt="image" src="https://github.com/thtrangphu/CleanArchitecture-TKPM/assets/76843467/0213279a-1db4-4b00-87b4-abb42fe30fcd">

**4. POST /signup: Register a new user.**
<img width="639" alt="image" src="https://github.com/thtrangphu/CleanArchitecture-TKPM/assets/76843467/93b6dd44-204a-4b09-bbee-002446eda029">

