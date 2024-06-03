package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/MashukeAlam/grails/models" // Adjust the import path accordingly
)

// GetMouses retrieves all Mouses from the database
func GetMouses(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var Mouses []models.Mouse
		if result := db.Find(&Mouses); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Render("mouses/index", fiber.Map{
			"Title": "All Mouses",
			"Records": Mouses,
		}, "layouts/main")
	}
}

// InsertMouse renders the insert form
func InsertMouse() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("mouses/insert", fiber.Map{
			"Title": "Add New Mouse",
		}, "layouts/main")
	}
}

// CreateMouse handles the form submission for creating a new Mouse
func CreateMouse(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		mouse := new(models.Mouse)
		if err := c.BodyParser(mouse); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if result := db.Create(mouse); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Redirect("/mouses")
	}
}

// EditMouse renders the edit form for a specific Mouse
func EditMouse(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var mouse models.Mouse
		if err := db.First(&mouse, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Mouse not found",
			})
		}
		return c.Render("mouses/edit", fiber.Map{"mouse": mouse, "Title": "Edit Entry"}, "layouts/main")
	}
}

// UpdateMouse handles the form submission for updating a Mouse
func UpdateMouse(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var mouse models.Mouse
		if err := db.First(&mouse, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Mouse not found",
			})
		}
		if err := c.BodyParser(&mouse); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if err := db.Save(&mouse).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update Mouse",
			})
		}
		return c.Redirect("/mouses")
	}
}

// DeleteMouse renders the delete confirmation view for a specific Mouse
func DeleteMouse(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var mouse models.Mouse
		if err := db.First(&mouse, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Mouse not found",
			})
		}
		return c.Render("mouses/delete", fiber.Map{"mouse": mouse, "Title": "Delete Entry"}, "layouts/main")
	}
}

// DestroyMouse handles the deletion of a Mouse
func DestroyMouse(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var mouse models.Mouse
		if err := db.First(&mouse, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Mouse not found",
			})
		}
		if err := db.Unscoped().Delete(&mouse).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete Mouse",
			})
		}
		return c.JSON(fiber.Map{"redirectUrl": "/mouses"})
	}
}
