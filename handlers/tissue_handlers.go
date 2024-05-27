package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"grails/models" // Adjust the import path accordingly
)

// GetTissues retrieves all Tissues from the database
func GetTissues(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var Tissues []models.Tissue
		if result := db.Find(&Tissues); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Render("tissues/index", fiber.Map{
			"Title": "All Tissues",
			"Records": Tissues,
		}, "layouts/main")
	}
}

// InsertTissue renders the insert form
func InsertTissue() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("tissues/insert", fiber.Map{
			"Title": "Add New Tissue",
		}, "layouts/main")
	}
}

// CreateTissue handles the form submission for creating a new Tissue
func CreateTissue(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tissue := new(models.Tissue)
		if err := c.BodyParser(tissue); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if result := db.Create(tissue); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Redirect("/tissues")
	}
}

// EditTissue renders the edit form for a specific Tissue
func EditTissue(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var tissue models.Tissue
		if err := db.First(&tissue, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Tissue not found",
			})
		}
		return c.Render("tissues/edit", fiber.Map{"tissue": tissue, "Title": "Edit Entry"}, "layouts/main")
	}
}

// UpdateTissue handles the form submission for updating a Tissue
func UpdateTissue(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var tissue models.Tissue
		if err := db.First(&tissue, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Tissue not found",
			})
		}
		if err := c.BodyParser(&tissue); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if err := db.Save(&tissue).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update Tissue",
			})
		}
		return c.Redirect("/tissues")
	}
}

// DeleteTissue renders the delete confirmation view for a specific Tissue
func DeleteTissue(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var tissue models.Tissue
		if err := db.First(&tissue, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Tissue not found",
			})
		}
		return c.Render("tissues/delete", fiber.Map{"tissue": tissue, "Title": "Delete Entry"}, "layouts/main")
	}
}

// DestroyTissue handles the deletion of a Tissue
func DestroyTissue(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var tissue models.Tissue
		if err := db.First(&tissue, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Tissue not found",
			})
		}
		if err := db.Unscoped().Delete(&tissue).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete Tissue",
			})
		}
		return c.JSON(fiber.Map{"redirectUrl": "/tissues"})
	}
}
