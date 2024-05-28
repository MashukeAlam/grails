package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/MashukeAlam/grails/models" // Adjust the import path accordingly
)

// GetHumans retrieves all Humans from the database
func GetHumans(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var Humans []models.Human
		if result := db.Find(&Humans); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Render("humans/index", fiber.Map{
			"Title": "All Humans",
			"Humans": Humans,
		}, "layouts/main")
	}
}

// InsertHuman renders the insert form
func InsertHuman() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("humans/insert", fiber.Map{
			"Title": "Add New Human",
		}, "layouts/main")
	}
}

// CreateHuman handles the form submission for creating a new Human
func CreateHuman(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		human := new(models.Human)
		if err := c.BodyParser(human); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if result := db.Create(human); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Redirect("/humans")
	}
}

// EditHuman renders the edit form for a specific Human
func EditHuman(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var human models.Human
		if err := db.First(&human, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Human not found",
			})
		}
		return c.Render("humans/edit", fiber.Map{"human": human, "Title": "Edit Entry"}, "layouts/main")
	}
}

// UpdateHuman handles the form submission for updating a Human
func UpdateHuman(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var human models.Human
		if err := db.First(&human, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Human not found",
			})
		}
		if err := c.BodyParser(&human); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if err := db.Save(&human).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update Human",
			})
		}
		return c.Redirect("/humans")
	}
}

// DeleteHuman renders the delete confirmation view for a specific Human
func DeleteHuman(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var human models.Human
		if err := db.First(&human, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Human not found",
			})
		}
		return c.Render("humans/delete", fiber.Map{"human": human, "Title": "Delete Entry"}, "layouts/main")
	}
}

// DestroyHuman handles the deletion of a Human
func DestroyHuman(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var human models.Human
		if err := db.First(&human, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Human not found",
			})
		}
		if err := db.Unscoped().Delete(&human).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete Human",
			})
		}
		return c.JSON(fiber.Map{"redirectUrl": "/humans"})
	}
}
