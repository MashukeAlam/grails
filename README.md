# Grails - A GoFiber web framework for Rails lovers!

![Release](https://img.shields.io/github/release/gofiber/boilerplate.svg)
[![Discord](https://img.shields.io/badge/discord-join%20channel-7289DA)](https://gofiber.io/discord)
![Test](https://github.com/gofiber/boilerplate/workflows/Test/badge.svg)
![Security](https://github.com/gofiber/boilerplate/workflows/Security/badge.svg)
![Linter](https://github.com/gofiber/boilerplate/workflows/Linter/badge.svg)


## IDE Development

### Visual Studio Code

Use the following plugins, in this boilerplate project:
- Name: Go
  - ID: golang.go
  - Description: Rich Go language support for Visual Studio Code
  - Version: 0.29.0
  - Editor: Go Team at Google
  - Link to Marketplace to VS: https://marketplace.visualstudio.com/items?itemName=golang.Go

## Development

### Start the application 


```bash
go run app.go
```

### Create migrations
Write less code and call generateMigrations. This will come to CLI soon.

```bash
tableName1 := "users"
fields1 := []Field{
	{Name: "name", Type: "VARCHAR(100) NOT NULL"},
	{Name: "email", Type: "VARCHAR(100) NOT NULL UNIQUE"},
}

// Generate the migration files
generateMigration(tableName1, fields1)

tableName2 := "tweets"
fields2 := []Field{
 	{Name: "body", Type: "VARCHAR(300) NOT NULL"},
 	{Name: "title", Type: "VARCHAR(100) NOT NULL"},
}

// Generate the migration files
generateMigration(tableName2, fields2)
```


### Run migrations


```bash
go run app.go migrate up
```


Go to http://localhost:5000:
