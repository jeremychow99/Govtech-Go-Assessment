# GovTech Internship Assessment (GDS-OCV)

REST API built using Golang, Gin and GORM with MySQL.

## Installation & Setup
Requirements: Golang installed, and a MySQL database to connect to.
1. Create a database/schema in MySQL.
2. In the .env file as shown below, modify the `DB_URL` accordingly based on your MySQL Credentials and name of schema/db you just created. Can also change `PORT` if required.
```
PORT=3000
DB_URL=user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
```
1. For example, I used a MySQL schema with dbname: test123, user: root, pass: root, so my DB_URL was `root:root@tcp(127.0.0.1:3306)/test123?charset=utf8mb4&parseTime=True&loc=Local`.
2. From folder root, run `go run migrate/migrate.go` to create database tables.
3. Run `go run main.go`.
4. API ready to be used.

## Some Notes
`main_test.go` which contains the unit tests does not work. Need to learn how to do proper unit tests for GORM (maybe using SqlMock?).