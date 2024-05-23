package main

import (
	"database/sql"
	"github.com/gofiber/template/html/v2"
	"regexp"
	"text/template"
	"time"

	"grails/database"
	"grails/handlers"
	"grails/internals"
	"grails/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"flag"
	"fmt"
	"log"
	"os"
	"strings"

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

type Field struct {
	Name string
	Type string
}

func capitalizeFirstLetter(str string) string {
	if len(str) == 0 {
		return str
	}
	return strings.ToUpper(string(str[0])) + str[1:]
}

// Converts a snake_case string to CamelCase.
func toCamelCase(str string) string {
	parts := strings.Split(str, "_")
	for i := range parts {
		parts[i] = capitalizeFirstLetter(parts[i])
	}
	return strings.Join(parts, "")
}

// toGoType maps SQL types to Go types
func toGoType(sqlType string) string {
	// Regular expression to match SQL types with optional length or precision
	re := regexp.MustCompile(`([a-zA-Z]+)(\(\d+\))?`)

	// Extract base type and optional length/precision
	matches := re.FindStringSubmatch(strings.ToUpper(sqlType))
	if len(matches) < 2 {
		return "string"
	}
	baseType := matches[1]

	switch baseType {
	case "VARCHAR", "CHAR", "NVARCHAR", "NCHAR", "CLOB", "TEXT":
		return "string"
	case "INT", "INTEGER", "TINYINT", "SMALLINT", "MEDIUMINT", "BIGINT":
		return "int"
	case "FLOAT", "DOUBLE", "REAL", "DECIMAL", "NUMERIC":
		return "float64"
	case "DATE", "DATETIME", "TIMESTAMP", "TIME", "YEAR":
		return "time.Time"
	case "BINARY", "VARBINARY", "BLOB", "LONGBLOB", "MEDIUMBLOB", "TINYBLOB":
		return "[]byte"
	case "BOOL", "BOOLEAN":
		return "bool"
	default:
		return "string"
	}
}

func generateHandlerFile(modelName string) {
	// Define the template
	const handlerTemplate = `package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"grails/models" // Adjust the import path accordingly
)

// Get{{.ModelName}}s retrieves all {{.ModelName}}s from the database
func Get{{.ModelName}}s(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var {{.ModelNamePlural}} []models.{{.ModelName}}
		if result := db.Find(&{{.ModelNamePlural}}); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Render("{{.ModelNameLowercase}}s/index", fiber.Map{
			"Title": "All {{.ModelName}}s",
			"{{.ModelNamePlural}}": {{.ModelNamePlural}},
		}, "layouts/main")
	}
}

// Insert{{.ModelName}} renders the insert form
func Insert{{.ModelName}}() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("{{.ModelNameLowercase}}s/insert", fiber.Map{
			"Title": "Add New {{.ModelName}}",
		}, "layouts/main")
	}
}

// Create{{.ModelName}} handles the form submission for creating a new {{.ModelName}}
func Create{{.ModelName}}(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		{{.ModelNameLowercase}} := new(models.{{.ModelName}})
		if err := c.BodyParser({{.ModelNameLowercase}}); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if result := db.Create({{.ModelNameLowercase}}); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Redirect("/{{.ModelNameLowercase}}s")
	}
}

// Edit{{.ModelName}} renders the edit form for a specific {{.ModelName}}
func Edit{{.ModelName}}(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var {{.ModelNameLowercase}} models.{{.ModelName}}
		if err := db.First(&{{.ModelNameLowercase}}, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "{{.ModelName}} not found",
			})
		}
		return c.Render("{{.ModelNameLowercase}}s/edit", fiber.Map{"{{.ModelNameLowercase}}": {{.ModelNameLowercase}}, "Title": "Edit Entry"}, "layouts/main")
	}
}

// Update{{.ModelName}} handles the form submission for updating a {{.ModelName}}
func Update{{.ModelName}}(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var {{.ModelNameLowercase}} models.{{.ModelName}}
		if err := db.First(&{{.ModelNameLowercase}}, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "{{.ModelName}} not found",
			})
		}
		if err := c.BodyParser(&{{.ModelNameLowercase}}); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if err := db.Save(&{{.ModelNameLowercase}}).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update {{.ModelName}}",
			})
		}
		return c.Redirect("/{{.ModelNameLowercase}}s")
	}
}

// Delete{{.ModelName}} renders the delete confirmation view for a specific {{.ModelName}}
func Delete{{.ModelName}}(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var {{.ModelNameLowercase}} models.{{.ModelName}}
		if err := db.First(&{{.ModelNameLowercase}}, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "{{.ModelName}} not found",
			})
		}
		return c.Render("{{.ModelNameLowercase}}s/delete", fiber.Map{"{{.ModelNameLowercase}}": {{.ModelNameLowercase}}, "Title": "Delete Entry"}, "layouts/main")
	}
}

// Destroy{{.ModelName}} handles the deletion of a {{.ModelName}}
func Destroy{{.ModelName}}(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var {{.ModelNameLowercase}} models.{{.ModelName}}
		if err := db.First(&{{.ModelNameLowercase}}, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "{{.ModelName}} not found",
			})
		}
		if err := db.Unscoped().Delete(&{{.ModelNameLowercase}}).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete {{.ModelName}}",
			})
		}
		return c.JSON(fiber.Map{"redirectUrl": "/{{.ModelNameLowercase}}s"})
	}
}
`
	// Prepare the data for the template
	data := struct {
		ModelName          string
		ModelNamePlural    string
		ModelNameLowercase string
	}{
		ModelName:          strings.Title(modelName),
		ModelNamePlural:    strings.Title(modelName) + "s",
		ModelNameLowercase: strings.ToLower(modelName),
	}

	// Parse and execute the template
	tmpl, err := template.New("handler").Parse(handlerTemplate)
	if err != nil {
		panic(err)
	}

	// Create the handler file
	handlerFileName := fmt.Sprintf("handlers/%s_handlers.go", strings.ToLower(modelName))
	handlerFile, err := os.Create(handlerFileName)
	if err != nil {
		panic(err)
	}
	defer handlerFile.Close()

	// Execute the template and write to the file
	err = tmpl.Execute(handlerFile, data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Handler file %s created successfully.\n", handlerFileName)
	// Generate the route registration code for the model
	routeRegistration := fmt.Sprintf(`
// %s routes
%s := app.Group("/%s")
%s.Get("/", handlers.Get%s(db))
%s.Get("/insert", handlers.Insert%s())
%s.Post("/", handlers.Create%s(db))
%s.Get("/:id/edit", handlers.Edit%s(db))
%s.Put("/:id", handlers.Update%s(db))
%s.Get("/:id/delete", handlers.Delete%s(db))
%s.Delete("/:id", handlers.Destroy%s(db))
`, strings.Title(modelName), modelName, modelName, modelName, strings.Title(modelName), modelName, strings.Title(modelName), modelName, strings.Title(modelName), modelName, strings.Title(modelName), modelName, strings.Title(modelName), modelName, strings.Title(modelName), modelName, modelName)

	// Print the route registration code in yellow color
	fmt.Println("\033[33m" + routeRegistration + "\033[0m")
}

func generateCreateMigration(tableName string, fields []Field, reference ...string) {
	// Define the migration directory
	migrationDir := "migrations"
	modelDir := "models"

	// Ensure the migration directory exists
	err := os.MkdirAll(migrationDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create migrations directory: %v", err)
	}

	// Generate the SQL for the migration
	migrationSQL := "-- +goose Up\n\n"
	migrationSQL += fmt.Sprintf("CREATE TABLE %ss (\n", tableName)
	migrationSQL += "  id INT AUTO_INCREMENT PRIMARY KEY,\n"
	for _, field := range fields {
		migrationSQL += fmt.Sprintf("  %s %s,\n", field.Name, field.Type)
	}
	if len(reference) > 0 {
		referenceTable := reference[0]
		migrationSQL += fmt.Sprintf("  %s INT NOT NULL,\n", referenceTable)
		migrationSQL += fmt.Sprintf("  FOREIGN KEY (%s) REFERENCES %s(id),\n", referenceTable, referenceTable)
	}
	migrationSQL += "  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP\n"
	migrationSQL += ");\n\n"

	migrationSQL += "-- +goose Down\n\n"
	migrationSQL += fmt.Sprintf("DROP TABLE %ss;", tableName)

	// Write the migration file
	time.Sleep(time.Millisecond * 1000)
	migrationFileName := fmt.Sprintf("%s/%s_create_%s_table.sql", migrationDir, time.Now().Format("20060102150405"), tableName)
	err = os.WriteFile(migrationFileName, []byte(migrationSQL), 0644)
	if err != nil {
		log.Fatalf("Failed to write migration file: %v", err)
	}

	fmt.Printf("Migration file %s created successfully.\n", migrationFileName)

	// Generate the Go model
	modelName := toCamelCase(tableName)
	modelContent := fmt.Sprintf("package models\n\nimport \"gorm.io/gorm\"\n\n// %s model\ntype %s struct {\n", modelName, modelName)
	modelContent += "	gorm.Model\n"
	for _, field := range fields {
		fieldName := toCamelCase(field.Name)
		goType := toGoType(field.Type)
		modelContent += fmt.Sprintf("	%s %s\n", fieldName, goType)
	}
	if len(reference) > 0 {
		referenceTable := reference[0]
		referenceField := toCamelCase(referenceTable)
		modelContent += fmt.Sprintf("	%sID int\n", referenceField)
		modelContent += fmt.Sprintf("	%s %s `gorm:\"foreignKey:%sID;references:ID\"`\n", referenceField, referenceField, referenceField)
	}
	modelContent += "}\n"

	// Write the model file.
	modelFileName := fmt.Sprintf("%s/%s.go", modelDir, tableName)
	err = os.WriteFile(modelFileName, []byte(modelContent), 0644)
	if err != nil {
		log.Fatalf("Failed to write model file: %v", err)
	}
	fmt.Printf("Model file %s created successfully.\n", modelFileName)

	// TODO: autoMigrate here gorm model.
	fmt.Printf("%s\n\n\ndbGorm.AutoMigrate(&models.%s{})%s\n\n\n", modelName, Green, Reset)
	generateHandlerFile(modelName)
	os.Exit(1)
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

	dsnWithDB := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?parseTime=true"
	db, err := sql.Open("mysql", dsnWithDB)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Database ready.")
	}
	defer db.Close()

	dbGorm, err := gorm.Open(mysql.Open(dsnWithDB), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	} else {
		fmt.Println("ORM Ready.")
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

	// Setup static files
	app.Static("/js", "./static/public/js")
	app.Static("/img", "./static/public/img")
	app.Static("/css", "./static/public/css")

	// Handle not founds
	app.Use(handlers.NotFound)
	//
	//tableName1 := "bike"
	//fields1 := []Field{
	//	{Name: "name", Type: "VARCHAR(100) NOT NULL"},
	//	{Name: "type", Type: "VARCHAR(10) NOT NULL"},
	//	{Name: "isCheap", Type: "TINYINT(1) NOT NULL"},
	//}
	//
	////Generate the migration files
	//generateCreateMigration(tableName1, fields1)

	// tableName2 := "player"
	// fields2 := []Field{
	// 	{Name: "name", Type: "VARCHAR(300) NOT NULL"},
	// }

	// // Generate the migration files
	// generateCreateMigration(tableName2, fields2, "game")

	// Listen on port 5000
	log.Fatal(app.Listen(*port))
}
