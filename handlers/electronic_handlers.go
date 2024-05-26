package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"grails/models" // Adjust the import path accordingly
)

// GetElectronics retrieves all Electronics from the database
func GetElectronics(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var Electronics []models.Electronic
		if result := db.Find(&Electronics); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Render("electronics/index", fiber.Map{
			"Title": "All Electronics",
			"Electronics": Electronics,
		}, "layouts/main")
	}
}

// InsertElectronic renders the insert form
func InsertElectronic() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("electronics/insert", fiber.Map{
			"Title": "Add New Electronic",
		}, "layouts/main")
	}
}

// CreateElectronic handles the form submission for creating a new Electronic
func CreateElectronic(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		electronic := new(models.Electronic)
		if err := c.BodyParser(electronic); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if result := db.Create(electronic); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Redirect("/electronics")
	}
}

// EditElectronic renders the edit form for a specific Electronic
func EditElectronic(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var electronic models.Electronic
		if err := db.First(&electronic, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Electronic not found",
			})
		}
		return c.Render("electronics/edit", fiber.Map{"electronic": electronic, "Title": "Edit Entry"}, "layouts/main")
	}
}

// UpdateElectronic handles the form submission for updating a Electronic
func UpdateElectronic(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var electronic models.Electronic
		if err := db.First(&electronic, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Electronic not found",
			})
		}
		if err := c.BodyParser(&electronic); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if err := db.Save(&electronic).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update Electronic",
			})
		}
		return c.Redirect("/electronics")
	}
}

// DeleteElectronic renders the delete confirmation view for a specific Electronic
func DeleteElectronic(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var electronic models.Electronic
		if err := db.First(&electronic, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Electronic not found",
			})
		}
		return c.Render("electronics/delete", fiber.Map{"electronic": electronic, "Title": "Delete Entry"}, "layouts/main")
	}
}

// DestroyElectronic handles the deletion of a Electronic
func DestroyElectronic(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var electronic models.Electronic
		if err := db.First(&electronic, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Electronic not found",
			})
		}
		if err := db.Unscoped().Delete(&electronic).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete Electronic",
			})
		}
		return c.JSON(fiber.Map{"redirectUrl": "/electronics"})
	}
}
