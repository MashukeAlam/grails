package main

import (
	"database/sql"
	"github.com/gofiber/template/html/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"grails/database"
	"grails/handlers"
	"grails/internals"
	"grails/models"

	"flag"
	"fmt"
	"log"
	"os"
	// "golang.org/x/text/cases"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
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

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		fmt.Printf("%sENV Loaded.%s\n", Green, Reset)
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
		fmt.Printf("%sDatabase server connected.%s\n", Green, Reset)
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
		fmt.Printf("%sDatabase connected.%s\n", Green, Reset)
	}

	dsnWithDB := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?parseTime=true"
	db, err := sql.Open("mysql", dsnWithDB)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%sDatabase ready.%s\n", Green, Reset)
	}
	defer db.Close()

	dbGorm, err := gorm.Open(mysql.Open(dsnWithDB), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	} else {
		fmt.Printf("%sORM Ready.%s\n", Green, Reset)
		DB = dbGorm
	}

	// Create a new engine
	engine := html.New("views", ".html")

	// Parse command-line flags
	flag.Parse()

	// Connected with database
	database.Connect()

	// Create fiber app
	app := fiber.New(fiber.Config{
		Prefork: *prod, // go run app.go -prod
		Views:   engine,
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Route to render the Slim template
	app.Get("/", func(c *fiber.Ctx) error {
		// Pass the title to the template
		return c.Render("index", fiber.Map{
			"Title": "Hello, Fiber with Slim!",
		}, "layouts/main")
	})

	// add models to watch for migration.
	dbGorm.AutoMigrate(&models.Game{})
	dbGorm.AutoMigrate(&models.Player{})
	dbGorm.AutoMigrate(&models.Country{})
	dbGorm.AutoMigrate(&models.Food{})
	dbGorm.AutoMigrate(&models.Car{})
	dbGorm.AutoMigrate(&models.Bike{})
	dbGorm.AutoMigrate(&models.Animal{})
	dbGorm.AutoMigrate(&models.Human{})
	dbGorm.AutoMigrate(&models.Language{})
	dbGorm.AutoMigrate(&models.Bird{})

	// Create a /game endpoint
	game := app.Group("/game")
	game.Get("/", handlers.GetGames(dbGorm))
	game.Get("/insert", handlers.InsertGame())
	game.Post("/", handlers.CreateGame(dbGorm))
	game.Get("/:id/edit", handlers.EditGame(dbGorm))
	game.Get("/:id/delete", handlers.DeleteGame(dbGorm))
	game.Post("/:id", handlers.UpdateGame(dbGorm))
	game.Delete("/:id", handlers.DestroyGame(dbGorm))

	// Car routes
	Car := app.Group("/Car")
	Car.Get("/", handlers.GetCars(dbGorm))
	Car.Get("/insert", handlers.InsertCar())
	Car.Post("/", handlers.CreateCar(dbGorm))
	Car.Get("/:id/edit", handlers.EditCar(dbGorm))
	Car.Post("/:id", handlers.UpdateCar(dbGorm))
	Car.Get("/:id/delete", handlers.DeleteCar(dbGorm))
	Car.Delete("/:id", handlers.DeleteCar(dbGorm))

	// Bike routes
	Bike := app.Group("/Bike")
	Bike.Get("/", handlers.GetBikes(dbGorm))
	Bike.Get("/insert", handlers.InsertBike())
	Bike.Post("/", handlers.CreateBike(dbGorm))
	Bike.Get("/:id/edit", handlers.EditBike(dbGorm))
	Bike.Put("/:id", handlers.UpdateBike(dbGorm))
	Bike.Get("/:id/delete", handlers.DeleteBike(dbGorm))
	Bike.Delete("/:id", handlers.DestroyBike(dbGorm))

	// Animal routes
	Animal := app.Group("/Animal")
	Animal.Get("/", handlers.GetAnimals(dbGorm))
	Animal.Get("/insert", handlers.InsertAnimal())
	Animal.Post("/", handlers.CreateAnimal(dbGorm))
	Animal.Get("/:id/edit", handlers.EditAnimal(dbGorm))
	Animal.Put("/:id", handlers.UpdateAnimal(dbGorm))
	Animal.Get("/:id/delete", handlers.DeleteAnimal(dbGorm))
	Animal.Delete("/:id", handlers.DestroyAnimal(dbGorm))

	// Human routes
	Human := app.Group("/Human")
	Human.Get("/", handlers.GetHumans(dbGorm))
	Human.Get("/insert", handlers.InsertHuman())
	Human.Post("/", handlers.CreateHuman(dbGorm))
	Human.Get("/:id/edit", handlers.EditHuman(dbGorm))
	Human.Put("/:id", handlers.UpdateHuman(dbGorm))
	Human.Get("/:id/delete", handlers.DeleteHuman(dbGorm))
	Human.Delete("/:id", handlers.DestroyHuman(dbGorm))

	// Language routes
	Language := app.Group("/Language")
	Language.Get("/", handlers.GetLanguages(dbGorm))
	Language.Get("/insert", handlers.InsertLanguage())
	Language.Post("/", handlers.CreateLanguage(dbGorm))
	Language.Get("/:id/edit", handlers.EditLanguage(dbGorm))
	Language.Put("/:id", handlers.UpdateLanguage(dbGorm))
	Language.Get("/:id/delete", handlers.DeleteLanguage(dbGorm))
	Language.Delete("/:id", handlers.DestroyLanguage(dbGorm))

	// Bird routes
	Bird := app.Group("/Bird")
	Bird.Get("/", handlers.GetBirds(dbGorm))
	Bird.Get("/insert", handlers.InsertBird())
	Bird.Post("/", handlers.CreateBird(dbGorm))
	Bird.Get("/:id/edit", handlers.EditBird(dbGorm))
	Bird.Put("/:id", handlers.UpdateBird(dbGorm))
	Bird.Get("/:id/delete", handlers.DeleteBird(dbGorm))
	Bird.Delete("/:id", handlers.DestroyBird(dbGorm))

	// Dev routes
	Dev := app.Group("/dev")
	Dev.Get("/", handlers.GetDevView())
	Dev.Post("/", handlers.ProcessIncomingScaffoldData(dbGorm))

	// Setup static files
	app.Static("/js", "./static/public/js")
	app.Static("/img", "./static/public/img")
	app.Static("/css", "./static/public/css")

	// Handle not founds
	app.Use(handlers.NotFound)

	// Listen on port 5000
	log.Fatal(app.Listen(*port))
}
