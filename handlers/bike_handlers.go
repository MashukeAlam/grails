package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"grails/models" // Adjust the import path accordingly
)

// GetBikes retrieves all Bikes from the database
func GetBikes(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var Bikes []models.Bike
		if result := db.Find(&Bikes); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Render("bikes/index", fiber.Map{
			"Title": "All Bikes",
			"Bikes": Bikes,
		}, "layouts/main")
	}
}

// InsertBike renders the insert form
func InsertBike() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("bikes/insert", fiber.Map{
			"Title": "Add New Bike",
		}, "layouts/main")
	}
}

// CreateBike handles the form submission for creating a new Bike
func CreateBike(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bike := new(models.Bike)
		if err := c.BodyParser(bike); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if result := db.Create(bike); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Redirect("/bikes")
	}
}

// EditBike renders the edit form for a specific Bike
func EditBike(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var bike models.Bike
		if err := db.First(&bike, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Bike not found",
			})
		}
		return c.Render("bikes/edit", fiber.Map{"bike": bike, "Title": "Edit Entry"}, "layouts/main")
	}
}

// UpdateBike handles the form submission for updating a Bike
func UpdateBike(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var bike models.Bike
		if err := db.First(&bike, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Bike not found",
			})
		}
		if err := c.BodyParser(&bike); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if err := db.Save(&bike).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update Bike",
			})
		}
		return c.Redirect("/bikes")
	}
}

// DeleteBike renders the delete confirmation view for a specific Bike
func DeleteBike(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var bike models.Bike
		if err := db.First(&bike, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Bike not found",
			})
		}
		return c.Render("bikes/delete", fiber.Map{"bike": bike, "Title": "Delete Entry"}, "layouts/main")
	}
}

// DestroyBike handles the deletion of a Bike
func DestroyBike(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var bike models.Bike
		if err := db.First(&bike, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Bike not found",
			})
		}
		if err := db.Unscoped().Delete(&bike).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete Bike",
			})
		}
		return c.JSON(fiber.Map{"redirectUrl": "/bikes"})
	}
}
