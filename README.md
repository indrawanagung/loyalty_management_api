# Loyalty Management System API

## Description

This project has 4 Domain layer :

- Model/Entity Layer
- Repository Layer
- Service Layer
- Controller Layer

### Tools Used:

In this project, I use some tools listed below. 

- Framework : Go  Fiber https://gofiber.io/
- ORM : GORM https://gorm.io/index.html
- Validation : Go Validator https://github.com/go-playground/validator
- Authentication : Json Web Token 
- Load Config : Viper https://github.com/spf13/viper

### How To Run :
#### Before you start 
- Install Docker https://www.docker.com/
- Install Golang Migrate https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
- Install Go https://go.dev/doc/install
#### Running Docker Postgre Image
Inside the project repository, run command `make postgres` to create and run container postgres database
```bash
### Run Docker Postgre SQL Image 
$ make postgres
```
#### Create Database Notes
To create database, run command `make createdb` or you can setting this configuration on Makefile `./Makefile`
```bash
### Create Database Notes on Posgre SQL
$ make createdb
```
#### Run Migration Schema
Run command `make migrateup` to create all schema to database notes automatically in file `db/migrations`. If you want to rollback all schema, run command `make migratedown`
```bash
### Create all schema to database notes
$ make migrateup
```
#### Install All Dependency 
To install all dependency in this project, run command `go get .`
```bash
### Install all dependency
$ go get .
```
#### Run Server API
Run express server api using command `npm start`
```bash
### Running Server API
$ go run main.go
```

#### Note :
Because I don't have much time to create open api documentation, I just create the postman files in the repository git
