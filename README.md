# Backend CRM Service
Mini project CRM Service Backen for e-commerce using Go Gin and Gorm with Onion Architecture.

[![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/doc/)

![go workflow](https://github.com/frederikoadr/backend-crm/actions/workflows/go.yml/badge.svg)

Database repo : https://gitlab.com/frederikoadr/crm-service

## Run

Use this command to run API app from root directory:

```shell
go run main.go
```

## Unit Tests

### Generate Mocks

To generate mock, run:

```shell
go generate ./...
```

### Run Unit Tests

To run unit tests:
```shell
go test ./...
```

---

## Coach Appointment Tech Spec

This README would normally document whatever steps are necessary to get your application up and running.

### Feature Description ###
there are three endpoint that could use for appointment proccess, which are:
1. Create Actors
2. Register Actors
3. Approval Actors
4. Login SuperAdmin and Admin Actors
5. Remove Admin Actors
6. Update (Activate/Deactivate Admin) Actors
7. Get All Customer
8. Get Customer by "first_name"/"last_name"/"email"
9. Get All Actor
10. Get Actor by "username"
11. Post API if Admin Get All Customer(7)

### Acceptance Criteria ###
1. Create Admin on Actor
2. Admin could Register
3. SuperAdmin could approve (change status) Admin on Register
4. SuperAdmin and Admin could login
5. SuperAdmin could remove Admin
6. SuperAdmin could activate/deactivate Admin on Actors
7. SuperAdmins and Admin could get all a customer data with parameter (by first_name and email)
8. SuperAdmins and Admin could get all a admin data with parameter (by username)
9. Admin could save data customers from API https://reqres.in/api/users?page=2 if customers is empty
10. Unit test

### Architecture and Design ###
This service using onion architecture, there are 5 layers 
from inner to outer which are entity, repository, use case,
controller, and request handler. the usage and responsibility of
each layer are follow:
1. **Entity**: this layer contains the domain model or entities
of the system. These are the core objects that 
represent the business concepts and rules.
2. **Repository**: This layer provides an interface for the 
application to access and manipulate the entities. 
It encapsulates the data access logic and provides
a way to abstract the database implementation details.
3. **Use case** : This layer contains the business logic 
or use cases of the system. It defines the operations 
that can be performed on the entities and orchestrates 
the interactions between the entities and the repository layer.
4. **Controller**: This layer handles the HTTP requests and
responses. It maps the incoming requests to the appropriate 
use case and returns the response to the client.
5. **Request handler**: This layer is responsible for handling 
the incoming HTTP requests and passing them on to 
the controller layer.
