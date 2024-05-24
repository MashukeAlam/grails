package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"grails/models" // Adjust the import path accordingly
)

// GetLanguages retrieves all Languages from the database
func GetLanguages(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var Languages []models.Language
		if result := db.Find(&Languages); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Render("languages/index", fiber.Map{
			"Title": "All Languages",
			"Languages": Languages,
		}, "layouts/main")
	}
}

// InsertLanguage renders the insert form
func InsertLanguage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("languages/insert", fiber.Map{
			"Title": "Add New Language",
		}, "layouts/main")
	}
}

// CreateLanguage handles the form submission for creating a new Language
func CreateLanguage(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		language := new(models.Language)
		if err := c.BodyParser(language); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if result := db.Create(language); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Redirect("/languages")
	}
}

// EditLanguage renders the edit form for a specific Language
func EditLanguage(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var language models.Language
		if err := db.First(&language, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Language not found",
			})
		}
		return c.Render("languages/edit", fiber.Map{"language": language, "Title": "Edit Entry"}, "layouts/main")
	}
}

// UpdateLanguage handles the form submission for updating a Language
func UpdateLanguage(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var language models.Language
		if err := db.First(&language, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Language not found",
			})
		}
		if err := c.BodyParser(&language); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if err := db.Save(&language).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update Language",
			})
		}
		return c.Redirect("/languages")
	}
}

// DeleteLanguage renders the delete confirmation view for a specific Language
func DeleteLanguage(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var language models.Language
		if err := db.First(&language, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Language not found",
			})
		}
		return c.Render("languages/delete", fiber.Map{"language": language, "Title": "Delete Entry"}, "layouts/main")
	}
}

// DestroyLanguage handles the deletion of a Language
func DestroyLanguage(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var language models.Language
		if err := db.First(&language, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Language not found",
			})
		}
		if err := db.Unscoped().Delete(&language).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete Language",
			})
		}
		return c.JSON(fiber.Map{"redirectUrl": "/languages"})
	}
}
