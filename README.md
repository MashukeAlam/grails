# Grails - A GoFiber web framework for Rails lovers 💚!


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
WARNING: Remember to uncomment these lines after running once or they will keep adding .sql files.

**This code will return the automigrate code and routing codes in the terminal. Paste those in the main function.**

### Run migrations using goose

If you are not using Gorm automigrate, you can use goose. that way you can migrate up & down. This feature is not tested so use with caution. Ensure Goose in PATH.
```bash
go run app.go migrate up
```


Go to http://localhost:5000

Attributions:
This project is built on the boilerplate that Fiber provides and I respects all the people that implemented the initial foundation.

[![Repo](https://img.shields.io/badge/repository-link-cyan)](https://github.com/gofiber/boilerplate)

[![Discord](https://img.shields.io/badge/discord-join%20channel-7289DA)](https://gofiber.io/discord)