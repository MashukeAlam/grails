package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"grails/models" // Adjust the import path accordingly
)

// GetLaptops retrieves all Laptops from the database
func GetLaptops(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var Laptops []models.Laptop
		if result := db.Find(&Laptops); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Render("laptops/index", fiber.Map{
			"Title": "All Laptops",
			"Records": Laptops,
		}, "layouts/main")
	}
}

// InsertLaptop renders the insert form
func InsertLaptop() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("laptops/insert", fiber.Map{
			"Title": "Add New Laptop",
		}, "layouts/main")
	}
}

// CreateLaptop handles the form submission for creating a new Laptop
func CreateLaptop(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		laptop := new(models.Laptop)
		if err := c.BodyParser(laptop); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if result := db.Create(laptop); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Redirect("/laptops")
	}
}

// EditLaptop renders the edit form for a specific Laptop
func EditLaptop(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var laptop models.Laptop
		if err := db.First(&laptop, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Laptop not found",
			})
		}
		return c.Render("laptops/edit", fiber.Map{"laptop": laptop, "Title": "Edit Entry"}, "layouts/main")
	}
}

// UpdateLaptop handles the form submission for updating a Laptop
func UpdateLaptop(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var laptop models.Laptop
		if err := db.First(&laptop, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Laptop not found",
			})
		}
		if err := c.BodyParser(&laptop); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if err := db.Save(&laptop).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update Laptop",
			})
		}
		return c.Redirect("/laptops")
	}
}

// DeleteLaptop renders the delete confirmation view for a specific Laptop
func DeleteLaptop(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var laptop models.Laptop
		if err := db.First(&laptop, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Laptop not found",
			})
		}
		return c.Render("laptops/delete", fiber.Map{"laptop": laptop, "Title": "Delete Entry"}, "layouts/main")
	}
}

// DestroyLaptop handles the deletion of a Laptop
func DestroyLaptop(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var laptop models.Laptop
		if err := db.First(&laptop, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Laptop not found",
			})
		}
		if err := db.Unscoped().Delete(&laptop).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete Laptop",
			})
		}
		return c.JSON(fiber.Map{"redirectUrl": "/laptops"})
	}
}
