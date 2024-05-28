package helpers

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func CreateModel(tableName string, fields []Field, reference ...string) {
	modelDir := "models"
	// Ensure the migration directory exists
	err := os.MkdirAll(modelDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create migrations directory: %v", err)
	}

	modelName := ToCamelCase(tableName)
	modelContent := fmt.Sprintf("package models\n\nimport \"gorm.io/gorm\"\n\n// %s model\ntype %s struct {\n", modelName, modelName)
	modelContent += "	gorm.Model\n"
	for _, field := range fields {
		fieldName := ToCamelCase(field.Name)
		goType := field.Type
		modelContent += fmt.Sprintf("	%s %s\n", fieldName, goType)
	}
	if len(reference) > 0 {
		referenceTable := reference[0]
		referenceField := ToCamelCase(referenceTable)
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

	appendMigrationCode(modelName)

	// TODO: autoMigrate here gorm model.
	fmt.Printf("%s\n\n\ndbGorm.AutoMigrate(&models.%s{})%s\n\n\n", Green, modelName, Reset)
	generateHandlerFile(modelName)

	// Generate index view template
	viewContent := generateIndexViewContent(tableName, fields)
	// Write the view file
	// View directory
	viewDirPlural := strings.ToLower(tableName)
	viewDirPlural += "s"
	viewDir := filepath.Join("views", viewDirPlural)

	// Ensure the view directory exists
	if err := os.MkdirAll(viewDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create views directory: %v", err)
	}
	viewFileName := filepath.Join(viewDir, "index.html")
	if err := os.WriteFile(viewFileName, []byte(viewContent), 0644); err != nil {
		log.Fatalf("Failed to write view file: %v", err)
	}
	fmt.Printf("Index View file %s created successfully.\n", viewFileName)

	insertContent := generateInsertViewContent(tableName, fields)
	viewFileName = filepath.Join(viewDir, "insert.html")
	if err := os.WriteFile(viewFileName, []byte(insertContent), 0644); err != nil {
		log.Fatalf("Failed to write view file: %v", err)
	}
	fmt.Printf("Insert View file %s created successfully.\n", viewFileName)

	os.Exit(1)

}

func appendMigrationCode(modelName string) {
	migrationFileName := "internals/migrations.go"
	migrationFunction := `package internals

import (
	"gorm.io/gorm"
	"github.com/MashukeAlam/grails/models"
)

func Migrate(db *gorm.DB) {
`
	migrationCode := fmt.Sprintf("\tdb.AutoMigrate(&models.%s{})\n", modelName)

	// Check if the migration file exists
	if _, err := os.Stat(migrationFileName); os.IsNotExist(err) {
		// Create the migration file with the function definition
		err := os.WriteFile(migrationFileName, []byte(migrationFunction+migrationCode+"}\n"), 0644)
		if err != nil {
			log.Fatalf("Failed to create migration file: %v", err)
		}
		fmt.Printf("Migration file %s created successfully.\n", migrationFileName)
	} else {
		content, err := os.ReadFile(migrationFileName)
		if err != nil {
			log.Fatalf("Failed to read migration file: %v", err)
		}

		contentStr := string(content)
		if strings.HasSuffix(contentStr, "}\n") {
			contentStr = strings.TrimSuffix(contentStr, "}\n")
		}

		contentStr += fmt.Sprintf("\n\t%s", migrationCode)
		contentStr += "\n}\n"

		err = os.WriteFile(migrationFileName, []byte(contentStr), 0644)
		if err != nil {
			log.Fatalf("Failed to write to migration file: %v", err)
		}
		fmt.Printf("Migration for %s appended to %s successfully.\n", modelName, migrationFileName)

	}
}

func appendRoutesCode(codeToAdd string) error {
	// Define the file path
	filePath := "internals/routes.go"

	// Check if the file already exists
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		// Create the file if it doesn't exist
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
	}

	// Read the existing content of the file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Convert content to string
	fileContent := string(content)

	// Find the index of closing braces of SetupRoutes function
	idx := strings.LastIndex(fileContent, "}")

	// Append the codeToAdd before the closing braces
	newContent := fileContent[:idx] + codeToAdd + "\n}" + fileContent[idx+1:]

	// Write the updated content back to the file
	err = ioutil.WriteFile(filePath, []byte(newContent), 0644)
	if err != nil {
		return err
	}

	return nil
}

func generateIndexViewContent(tableName string, fields []Field) string {
	var tableHeaders, tableRows strings.Builder

	// Generate table headers
	for _, field := range fields {
		tableHeaders.WriteString(fmt.Sprintf("<th>%s</th>", field.Name))
	}

	// Generate table rows
	tableRows.WriteString("{{range .Records}}<tr>")
	for _, field := range fields {
		tableRows.WriteString(fmt.Sprintf("<td>{{.%s}}</td>", ToCamelCase(field.Name)))
	}
	tableRows.WriteString(fmt.Sprintf(`
        <td>
            <a href="%ss/{{.ID}}/edit">Edit</a> |
            <a href="%ss/{{.ID}}/delete">Delete</a>
        </td>
        <td>{{.CreatedAt}}</td>
    </tr>{{end}}`, strings.ToLower(tableName), strings.ToLower(tableName)))

	return fmt.Sprintf(`
    <h2>All %s</h2>
    <a href="/%ss/insert">Add +</a>
    <table>
        <thead>
            <tr>%s<th>Actions</th><th>Created At</th></tr>
        </thead>
        <tbody>%s</tbody>
    </table>
    `, tableName, tableName, tableHeaders.String(), tableRows.String())
}

func generateInsertViewContent(tableName string, fields []Field) string {
	var formFields strings.Builder
	for _, field := range fields {
		formFields.WriteString(fmt.Sprintf(`
            <label for="%s">%s:</label>
            <input type="%s" id="%s" name="%s" required>
        `, field.Name, field.Name, field.Type, field.Name, field.Name))
	}

	return fmt.Sprintf(`
    <h2>Add %s</h2>
    <form action="/%ss" method="POST">
        %s
        <button type="submit">Add %s</button>
    </form>
    `, tableName, tableName, formFields.String(), tableName)
}

func generateHandlerFile(modelName string) {
	// Define the template
	const handlerTemplate = `package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/MashukeAlam/grails/models" // Adjust the import path accordingly
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
			"Records": {{.ModelNamePlural}},
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

	// Generate the route registration code for the model
	routeRegistration := fmt.Sprintf(`
// %s routes
%s := app.Group("/%ss")
%s.Get("/", handlers.Get%ss(dbGorm))
%s.Get("/insert", handlers.Insert%s())
%s.Post("/", handlers.Create%s(dbGorm))
%s.Get("/:id/edit", handlers.Edit%s(dbGorm))
%s.Put("/:id", handlers.Update%s(dbGorm))
%s.Get("/:id/delete", handlers.Delete%s(dbGorm))
%s.Delete("/:id", handlers.Destroy%s(dbGorm))
`, strings.Title(modelName), modelName, modelName, modelName, strings.Title(modelName), modelName, strings.Title(modelName), modelName, strings.Title(modelName), modelName, strings.Title(modelName), modelName, strings.Title(modelName), modelName, strings.Title(modelName), modelName, modelName)

	// Print the route registration code in yellow color
	fmt.Println("\033[33m" + routeRegistration + "\033[0m")
	appendRoutesCode(routeRegistration)
	fmt.Printf("Handler file %s created successfully.\n", handlerFileName)
}
