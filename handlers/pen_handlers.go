package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"grails/models" // Adjust the import path accordingly
)

// GetPens retrieves all Pens from the database
func GetPens(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var Pens []models.Pen
		if result := db.Find(&Pens); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Render("pens/index", fiber.Map{
			"Title": "All Pens",
			"Records": Pens,
		}, "layouts/main")
	}
}

// InsertPen renders the insert form
func InsertPen() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("pens/insert", fiber.Map{
			"Title": "Add New Pen",
		}, "layouts/main")
	}
}

// CreatePen handles the form submission for creating a new Pen
func CreatePen(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		pen := new(models.Pen)
		if err := c.BodyParser(pen); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if result := db.Create(pen); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Redirect("/pens")
	}
}

// EditPen renders the edit form for a specific Pen
func EditPen(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var pen models.Pen
		if err := db.First(&pen, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Pen not found",
			})
		}
		return c.Render("pens/edit", fiber.Map{"pen": pen, "Title": "Edit Entry"}, "layouts/main")
	}
}

// UpdatePen handles the form submission for updating a Pen
func UpdatePen(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var pen models.Pen
		if err := db.First(&pen, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Pen not found",
			})
		}
		if err := c.BodyParser(&pen); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if err := db.Save(&pen).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update Pen",
			})
		}
		return c.Redirect("/pens")
	}
}

// DeletePen renders the delete confirmation view for a specific Pen
func DeletePen(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var pen models.Pen
		if err := db.First(&pen, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Pen not found",
			})
		}
		return c.Render("pens/delete", fiber.Map{"pen": pen, "Title": "Delete Entry"}, "layouts/main")
	}
}

// DestroyPen handles the deletion of a Pen
func DestroyPen(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var pen models.Pen
		if err := db.First(&pen, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Pen not found",
			})
		}
		if err := db.Unscoped().Delete(&pen).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete Pen",
			})
		}
		return c.JSON(fiber.Map{"redirectUrl": "/pens"})
	}
}
