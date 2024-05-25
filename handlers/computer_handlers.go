package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"grails/models" // Adjust the import path accordingly
)

// GetComputers retrieves all Computers from the database
func GetComputers(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var Computers []models.Computer
		if result := db.Find(&Computers); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Render("computers/index", fiber.Map{
			"Title": "All Computers",
			"Computers": Computers,
		}, "layouts/main")
	}
}

// InsertComputer renders the insert form
func InsertComputer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("computers/insert", fiber.Map{
			"Title": "Add New Computer",
		}, "layouts/main")
	}
}

// CreateComputer handles the form submission for creating a new Computer
func CreateComputer(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		computer := new(models.Computer)
		if err := c.BodyParser(computer); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if result := db.Create(computer); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Redirect("/computers")
	}
}

// EditComputer renders the edit form for a specific Computer
func EditComputer(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var computer models.Computer
		if err := db.First(&computer, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Computer not found",
			})
		}
		return c.Render("computers/edit", fiber.Map{"computer": computer, "Title": "Edit Entry"}, "layouts/main")
	}
}

// UpdateComputer handles the form submission for updating a Computer
func UpdateComputer(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var computer models.Computer
		if err := db.First(&computer, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Computer not found",
			})
		}
		if err := c.BodyParser(&computer); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if err := db.Save(&computer).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update Computer",
			})
		}
		return c.Redirect("/computers")
	}
}

// DeleteComputer renders the delete confirmation view for a specific Computer
func DeleteComputer(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var computer models.Computer
		if err := db.First(&computer, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Computer not found",
			})
		}
		return c.Render("computers/delete", fiber.Map{"computer": computer, "Title": "Delete Entry"}, "layouts/main")
	}
}

// DestroyComputer handles the deletion of a Computer
func DestroyComputer(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var computer models.Computer
		if err := db.First(&computer, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Computer not found",
			})
		}
		if err := db.Unscoped().Delete(&computer).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete Computer",
			})
		}
		return c.JSON(fiber.Map{"redirectUrl": "/computers"})
	}
}
