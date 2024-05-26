package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"grails/models" // Adjust the import path accordingly
)

// GetGirls retrieves all Girls from the database
func GetGirls(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var Girls []models.Girl
		if result := db.Find(&Girls); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Render("girls/index", fiber.Map{
			"Title": "All Girls",
			"Girls": Girls,
		}, "layouts/main")
	}
}

// InsertGirl renders the insert form
func InsertGirl() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("girls/insert", fiber.Map{
			"Title": "Add New Girl",
		}, "layouts/main")
	}
}

// CreateGirl handles the form submission for creating a new Girl
func CreateGirl(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		girl := new(models.Girl)
		if err := c.BodyParser(girl); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if result := db.Create(girl); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Redirect("/girls")
	}
}

// EditGirl renders the edit form for a specific Girl
func EditGirl(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var girl models.Girl
		if err := db.First(&girl, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Girl not found",
			})
		}
		return c.Render("girls/edit", fiber.Map{"girl": girl, "Title": "Edit Entry"}, "layouts/main")
	}
}

// UpdateGirl handles the form submission for updating a Girl
func UpdateGirl(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var girl models.Girl
		if err := db.First(&girl, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Girl not found",
			})
		}
		if err := c.BodyParser(&girl); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if err := db.Save(&girl).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update Girl",
			})
		}
		return c.Redirect("/girls")
	}
}

// DeleteGirl renders the delete confirmation view for a specific Girl
func DeleteGirl(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var girl models.Girl
		if err := db.First(&girl, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Girl not found",
			})
		}
		return c.Render("girls/delete", fiber.Map{"girl": girl, "Title": "Delete Entry"}, "layouts/main")
	}
}

// DestroyGirl handles the deletion of a Girl
func DestroyGirl(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var girl models.Girl
		if err := db.First(&girl, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Girl not found",
			})
		}
		if err := db.Unscoped().Delete(&girl).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete Girl",
			})
		}
		return c.JSON(fiber.Map{"redirectUrl": "/girls"})
	}
}
