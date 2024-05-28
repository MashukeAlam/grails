package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/MashukeAlam/grails/cmd"
	"github.com/MashukeAlam/grails/handlers"
	"github.com/MashukeAlam/grails/internals"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

// ANSI escape codes for colors
const (
	Red   = "\033[31m"
	Green = "\033[32m"
	Reset = "\033[0m"
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

func main() {
	cmd.Execute()
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("%sError loading .env file%s", Red, Reset)
	} else {
		fmt.Printf("%sENV Loaded.%s\n", Green, Reset)
	}

	// Database server connection string without database
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/"
	dbNot, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%sDatabase server connected.%s\n", Green, Reset)
	}
	defer dbNot.Close()

	// Verify the connection
	if err := dbNot.Ping(); err != nil {
		log.Fatal(err)
	}

	// Create the database if not exist
	// TODO: send this back at other functions may be in internals
	err = createDatabase(dbNot, os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	} else {
		fmt.Printf("%sDatabase connected.%s\n", Green, Reset)
	}

	// Database connection string with DB
	dsnWithDB := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?parseTime=true"
	db, err := sql.Open("mysql", dsnWithDB)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%sDatabase ready.%s\n", Green, Reset)
	}
	defer db.Close()

	// DBGORM
	dbGorm, err := gorm.Open(mysql.Open(dsnWithDB), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	} else {
		fmt.Printf("%sORM Ready.%s\n", Green, Reset)
		DB = dbGorm
	}

	if len(os.Args) > 1 {
		// Check for migration command
		if os.Args[1] == "migrate" {
			internals.Migrate(dbGorm) // Run the migrate function with the direction
			return
		} else {
			log.Println("For Migration, please type: go run app.go migrate")
			os.Exit(1)
		}
	}

	// Create a new engine
	engine := html.New("views", ".html")

	// Parse command-line flags
	flag.Parse()

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -prod
		Views:   engine,
	})

	app = internals.FiberAppStart(app)
	internals.SetupRoutes(app, dbGorm)
	app.Use(handlers.NotFound)

	// Listen on port 5000
	log.Fatal(app.Listen(*port))
}
