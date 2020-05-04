# Gimic

Gimic is simple boilerplate for building Microservice with Gin Gonic + Go-Micro.

## Installation

#### Clone the project

```bash
git clone https://github.com/hajidalakhtar/Gimic.git
```

#### Install required library

Mysql database
```bash
go get github.com/go-sql-driver/mysql
```
Go Framework
```bash
go get -u github.com/gin-gonic/gin
```
Go-Micro
```bash
go get github.com/micro/micro/v2
```
Object Relation Mapping
```bash
go get -u github.com/jinzhu/gorm
```

#### Running
Serving the project
```bash
go run main.go serve
```

## Avaible CLI Commands
|COMMAND|DESCRIPTION|
|-------|-----------|
|```go run main.go serve```|running server instance|
|```go run main.go migrate```|running auto migration|
|```go run main.go microservice```|running auto microservice|





## License
[MIT](https://choosealicense.com/licenses/mit/)

