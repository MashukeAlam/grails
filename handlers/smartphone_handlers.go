package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/MashukeAlam/grails/models" // Adjust the import path accordingly
)

// GetSmartphones retrieves all Smartphones from the database
func GetSmartphones(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var Smartphones []models.Smartphone
		if result := db.Find(&Smartphones); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Render("smartphones/index", fiber.Map{
			"Title": "All Smartphones",
			"Records": Smartphones,
		}, "layouts/main")
	}
}

// InsertSmartphone renders the insert form
func InsertSmartphone() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("smartphones/insert", fiber.Map{
			"Title": "Add New Smartphone",
		}, "layouts/main")
	}
}

// CreateSmartphone handles the form submission for creating a new Smartphone
func CreateSmartphone(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		smartphone := new(models.Smartphone)
		if err := c.BodyParser(smartphone); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if result := db.Create(smartphone); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Redirect("/smartphones")
	}
}

// EditSmartphone renders the edit form for a specific Smartphone
func EditSmartphone(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var smartphone models.Smartphone
		if err := db.First(&smartphone, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Smartphone not found",
			})
		}
		return c.Render("smartphones/edit", fiber.Map{"smartphone": smartphone, "Title": "Edit Entry"}, "layouts/main")
	}
}

// UpdateSmartphone handles the form submission for updating a Smartphone
func UpdateSmartphone(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var smartphone models.Smartphone
		if err := db.First(&smartphone, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Smartphone not found",
			})
		}
		if err := c.BodyParser(&smartphone); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if err := db.Save(&smartphone).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update Smartphone",
			})
		}
		return c.Redirect("/smartphones")
	}
}

// DeleteSmartphone renders the delete confirmation view for a specific Smartphone
func DeleteSmartphone(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var smartphone models.Smartphone
		if err := db.First(&smartphone, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Smartphone not found",
			})
		}
		return c.Render("smartphones/delete", fiber.Map{"smartphone": smartphone, "Title": "Delete Entry"}, "layouts/main")
	}
}

// DestroySmartphone handles the deletion of a Smartphone
func DestroySmartphone(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var smartphone models.Smartphone
		if err := db.First(&smartphone, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Smartphone not found",
			})
		}
		if err := db.Unscoped().Delete(&smartphone).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete Smartphone",
			})
		}
		return c.JSON(fiber.Map{"redirectUrl": "/smartphones"})
	}
}
