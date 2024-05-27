package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"grails/models" // Adjust the import path accordingly
)

// GetMonitors retrieves all Monitors from the database
func GetMonitors(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var Monitors []models.Monitor
		if result := db.Find(&Monitors); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Render("monitors/index", fiber.Map{
			"Title": "All Monitors",
			"Records": Monitors,
		}, "layouts/main")
	}
}

// InsertMonitor renders the insert form
func InsertMonitor() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("monitors/insert", fiber.Map{
			"Title": "Add New Monitor",
		}, "layouts/main")
	}
}

// CreateMonitor handles the form submission for creating a new Monitor
func CreateMonitor(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		monitor := new(models.Monitor)
		if err := c.BodyParser(monitor); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if result := db.Create(monitor); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Redirect("/monitors")
	}
}

// EditMonitor renders the edit form for a specific Monitor
func EditMonitor(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var monitor models.Monitor
		if err := db.First(&monitor, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Monitor not found",
			})
		}
		return c.Render("monitors/edit", fiber.Map{"monitor": monitor, "Title": "Edit Entry"}, "layouts/main")
	}
}

// UpdateMonitor handles the form submission for updating a Monitor
func UpdateMonitor(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var monitor models.Monitor
		if err := db.First(&monitor, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Monitor not found",
			})
		}
		if err := c.BodyParser(&monitor); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if err := db.Save(&monitor).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update Monitor",
			})
		}
		return c.Redirect("/monitors")
	}
}

// DeleteMonitor renders the delete confirmation view for a specific Monitor
func DeleteMonitor(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var monitor models.Monitor
		if err := db.First(&monitor, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Monitor not found",
			})
		}
		return c.Render("monitors/delete", fiber.Map{"monitor": monitor, "Title": "Delete Entry"}, "layouts/main")
	}
}

// DestroyMonitor handles the deletion of a Monitor
func DestroyMonitor(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var monitor models.Monitor
		if err := db.First(&monitor, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Monitor not found",
			})
		}
		if err := db.Unscoped().Delete(&monitor).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete Monitor",
			})
		}
		return c.JSON(fiber.Map{"redirectUrl": "/monitors"})
	}
}
