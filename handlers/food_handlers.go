package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"grails/models" // Adjust the import path accordingly
)

// GetFoods retrieves all Foods from the database
func GetFoods(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var Foods []models.Food
		if result := db.Find(&Foods); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Render("foods/index", fiber.Map{
			"Title": "All Foods",
			"Foods": Foods,
		}, "layouts/main")
	}
}

// InsertFood renders the insert form
func InsertFood() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("foods/insert", fiber.Map{
			"Title": "Add New Food",
		}, "layouts/main")
	}
}

// CreateFood handles the form submission for creating a new Food
func CreateFood(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		food := new(models.Food)
		if err := c.BodyParser(food); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if result := db.Create(food); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Redirect("/foods")
	}
}

// EditFood renders the edit form for a specific Food
func EditFood(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var food models.Food
		if err := db.First(&food, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Food not found",
			})
		}
		return c.Render("foods/edit", fiber.Map{"food": food, "Title": "Edit Entry"}, "layouts/main")
	}
}

// UpdateFood handles the form submission for updating a Food
func UpdateFood(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var food models.Food
		if err := db.First(&food, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Food not found",
			})
		}
		if err := c.BodyParser(&food); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if err := db.Save(&food).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update Food",
			})
		}
		return c.Redirect("/foods")
	}
}

// DeleteFood renders the delete confirmation view for a specific Food
func DeleteFood(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var food models.Food
		if err := db.First(&food, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Food not found",
			})
		}
		return c.Render("foods/delete", fiber.Map{"food": food, "Title": "Delete Entry"}, "layouts/main")
	}
}

// DestroyFood handles the deletion of a Food
func DestroyFood(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var food models.Food
		if err := db.First(&food, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Food not found",
			})
		}
		if err := db.Unscoped().Delete(&food).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete Food",
			})
		}
		return c.JSON(fiber.Map{"redirectUrl": "/foods"})
	}
}
