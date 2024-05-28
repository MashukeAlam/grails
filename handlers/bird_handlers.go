package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/MashukeAlam/grails/models" // Adjust the import path accordingly
)

// GetBirds retrieves all Birds from the database
func GetBirds(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var Birds []models.Bird
		if result := db.Find(&Birds); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Render("birds/index", fiber.Map{
			"Title": "All Birds",
			"Birds": Birds,
		}, "layouts/main")
	}
}

// InsertBird renders the insert form
func InsertBird() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("birds/insert", fiber.Map{
			"Title": "Add New Bird",
		}, "layouts/main")
	}
}

// CreateBird handles the form submission for creating a new Bird
func CreateBird(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bird := new(models.Bird)
		if err := c.BodyParser(bird); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if result := db.Create(bird); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Redirect("/birds")
	}
}

// EditBird renders the edit form for a specific Bird
func EditBird(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var bird models.Bird
		if err := db.First(&bird, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Bird not found",
			})
		}
		return c.Render("birds/edit", fiber.Map{"bird": bird, "Title": "Edit Entry"}, "layouts/main")
	}
}

// UpdateBird handles the form submission for updating a Bird
func UpdateBird(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var bird models.Bird
		if err := db.First(&bird, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Bird not found",
			})
		}
		if err := c.BodyParser(&bird); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if err := db.Save(&bird).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update Bird",
			})
		}
		return c.Redirect("/birds")
	}
}

// DeleteBird renders the delete confirmation view for a specific Bird
func DeleteBird(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var bird models.Bird
		if err := db.First(&bird, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Bird not found",
			})
		}
		return c.Render("birds/delete", fiber.Map{"bird": bird, "Title": "Delete Entry"}, "layouts/main")
	}
}

// DestroyBird handles the deletion of a Bird
func DestroyBird(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var bird models.Bird
		if err := db.First(&bird, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Bird not found",
			})
		}
		if err := db.Unscoped().Delete(&bird).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete Bird",
			})
		}
		return c.JSON(fiber.Map{"redirectUrl": "/birds"})
	}
}
