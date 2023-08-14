# Implementing CRUD in GoLang REST API with Mux & GORM

![CRUD in GoLang REST API with Mux & GORM](https://codewithmukesh.com/wp-content/uploads/2022/03/CRUD-in-Golang-REST-API-with-Mux-GORM-Simple-Guide.png)

In this article, we will learn about implementing CRUD in Golang REST API with Gorilla Mux for routing requests, GORM as the ORM to access the database, Viper for Loading configurations, and PostgreSQL as the database provider. We will also follow some clean development practices to help organize the GoLang project folder structure in a more easy-to-understand fashion.

## Topics Covered
- Setting up the Golang Project
- Loading Configurations using Viper
- Defining the Product Entity
- Connecting to the database
- Routing
- Implementing CRUD in Golang Rest API
	 - Create
	 - Get By ID
	 - Get All
	 - Update
	 - Delete
- Testing CRUD Operations

Read the entire article here - https://codewithmukesh.com/blog/implementing-crud-in-golang-rest-api/

## Config your own Golang
- Create your own database on PostgreSQL Server
- Rename config.json.example to config.json
- Change the value of 'connection_string' in the config file:
```
{
  "connection_string": "user=username password=yourpassword dbname=yourdb host=localhost port=5432 sslmode=disable TimeZone=Asia/Singapore",
  "domain": "http://localhost",
  "port": 8080
}
```

## Up to date the dependencies version (Recommended)
```
go get -u
```

## Or only download current dependencies version (Not Recommended)
```
go get
```

## Run on local/dev mode
```
go run .
```

## Build
```
go build -o file_name
```

## Run Build at Terminal
```
./file_name
```

## Test Rest API
Install plugin on vscode named: ```REST Client``` or others like ```'ext:rest'```\
```Test the rest API at folder:```
```
~/golang-crud-rest-api/collection/products.rest
```
