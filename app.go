package main

import (
	"database/sql"
	"time"

	"grails/database"
	"grails/handlers"
	"grails/internals"

	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/slim/v2"

	"github.com/joho/godotenv"
)

var (
	port = flag.String("port", ":5000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func createDatabase(db *sql.DB, dbName string) error {
    // Create the database if it doesn't exist
    _, err := db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
    return err
}

type Field struct {
	Name string
	Type string
}

func generateMigration(tableName string, fields []Field) {
	// Define the migration directory
	migrationDir := "migrations"

	// Ensure the migration directory exists
	err := os.MkdirAll(migrationDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create migrations directory: %v", err)
	}

	// Generate the SQL for the UP migration
	upSQL := "-- +goose Up\n\n"
	upSQL += fmt.Sprintf("CREATE TABLE %s (\n", tableName)
	upSQL += "  id INT AUTO_INCREMENT PRIMARY KEY,\n"
	for _, field := range fields {
		upSQL += fmt.Sprintf("  %s %s,\n", field.Name, field.Type)
	}
	upSQL += "  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP\n"
	upSQL += ");"
	
	// Generate the SQL for the DOWN migration
	// downSQL := "-- +goose Down\n\n"
	// downSQL += fmt.Sprintf("DROP TABLE %s;", tableName)
	
	// Write the UP migration file
	time.Sleep(time.Millisecond * 1000)
	upFileName := fmt.Sprintf("%s/%s_create_%s_table.up.sql", migrationDir, time.Now().Format("20060102150405"), tableName)
	err = os.WriteFile(upFileName, []byte(upSQL), 0644)
	if err != nil {
		log.Fatalf("Failed to write UP migration file: %v", err)
	}
	
	// Write the DOWN migration file
	// time.Sleep(time.Millisecond * 1000)
	// downFileName := fmt.Sprintf("%s/%s_create_%s_table.down.sql", migrationDir, time.Now().Format("20060102150405"), tableName)
	// err = os.WriteFile(downFileName, []byte(downSQL), 0644)
	// if err != nil {
	// 	log.Fatalf("Failed to write DOWN migration file: %v", err)
	// }

	fmt.Printf("Migration files %s created successfully.\n", upFileName)
}


func main() {
	
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    } else {
		fmt.Println("ENV Loaded.")
	}

	if len(os.Args) > 1 {
		// Check for migration command
		if os.Args[1] == "migrate" && len(os.Args) == 3 {
			internals.Migrate(os.Args[2]) // Run the migrate function with the direction
			return
		} else {
			log.Println("Usage: app migrate <up|down>")
			os.Exit(1)
		}
	}
    // Database connection string
    dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/"
    dbNot, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    } else {
		fmt.Println("Database server connected.")
	}
    defer dbNot.Close()

    // Verify the connection
    if err := dbNot.Ping(); err != nil {
        log.Fatal(err)
    }

	// Create the
	err = createDatabase(dbNot, os.Getenv("DB_NAME"))
    if err != nil {
        log.Fatalf("Failed to create database: %v", err)
    } else {
		fmt.Println("Database connected.")
	}

	dsnWithDB := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME")
    db, err := sql.Open("mysql", dsnWithDB)
    if err != nil {
        log.Fatal(err)
    } else {
		fmt.Println("Database ready.")
	}
    defer db.Close()

	// Create a new engine
	engine := slim.New("./views", ".slim")

	// Parse command-line flags
	flag.Parse()

	// Connected with database
	database.Connect()

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -prod
		Views: engine,
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Route to render the Slim template
	app.Get("/", func(c *fiber.Ctx) error {
		// Pass the title to the template
		return c.Render("index", fiber.Map{
			"Title": "Hello, Fiber with Slim!",
		})
	})

	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")

	// Bind handlers
	v1.Get("/users", handlers.UserList)
	v1.Post("/users", handlers.UserCreate)

	// Setup static files
	app.Static("/js", "./static/public/js")
	app.Static("/img", "./static/public/img")

	// Handle not founds
	app.Use(handlers.NotFound)

	// tableName1 := "users"
	// fields1 := []Field{
	// 	{Name: "name", Type: "VARCHAR(100) NOT NULL"},
	// 	{Name: "email", Type: "VARCHAR(100) NOT NULL UNIQUE"},
	// }

	// // Generate the migration files
	// generateMigration(tableName1, fields1)

	// tableName2 := "tweets"
	// fields2 := []Field{
	// 	{Name: "body", Type: "VARCHAR(300) NOT NULL"},
	// 	{Name: "title", Type: "VARCHAR(100) NOT NULL"},
	// }

	// // Generate the migration files
	// generateMigration(tableName2, fields2)

	//Listen on port 5000
	log.Fatal(app.Listen(*port)) 
}
