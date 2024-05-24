package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"grails/models" // Adjust the import path accordingly
)

// GetBirdss retrieves all Birdss from the database
func GetBirdss(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var Birdss []models.Birds
		if result := db.Find(&Birdss); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Render("birdss/index", fiber.Map{
			"Title": "All Birdss",
			"Birdss": Birdss,
		}, "layouts/main")
	}
}

// InsertBirds renders the insert form
func InsertBirds() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("birdss/insert", fiber.Map{
			"Title": "Add New Birds",
		}, "layouts/main")
	}
}

// CreateBirds handles the form submission for creating a new Birds
func CreateBirds(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		birds := new(models.Birds)
		if err := c.BodyParser(birds); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if result := db.Create(birds); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Redirect("/birdss")
	}
}

// EditBirds renders the edit form for a specific Birds
func EditBirds(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var birds models.Birds
		if err := db.First(&birds, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Birds not found",
			})
		}
		return c.Render("birdss/edit", fiber.Map{"birds": birds, "Title": "Edit Entry"}, "layouts/main")
	}
}

// UpdateBirds handles the form submission for updating a Birds
func UpdateBirds(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var birds models.Birds
		if err := db.First(&birds, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Birds not found",
			})
		}
		if err := c.BodyParser(&birds); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if err := db.Save(&birds).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update Birds",
			})
		}
		return c.Redirect("/birdss")
	}
}

// DeleteBirds renders the delete confirmation view for a specific Birds
func DeleteBirds(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var birds models.Birds
		if err := db.First(&birds, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Birds not found",
			})
		}
		return c.Render("birdss/delete", fiber.Map{"birds": birds, "Title": "Delete Entry"}, "layouts/main")
	}
}

// DestroyBirds handles the deletion of a Birds
func DestroyBirds(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var birds models.Birds
		if err := db.First(&birds, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Birds not found",
			})
		}
		if err := db.Unscoped().Delete(&birds).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete Birds",
			})
		}
		return c.JSON(fiber.Map{"redirectUrl": "/birdss"})
	}
}
