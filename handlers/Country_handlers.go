package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/MashukeAlam/grails/models" // Adjust the import path accordingly
)

// GetCountrys retrieves all Countrys from the database
func GetCountrys(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var Countrys []models.Country
		if result := db.Find(&Countrys); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Render("countrys/index", fiber.Map{
			"Title":    "All Countrys",
			"Countrys": Countrys,
		}, "layouts/main")
	}
}

// InsertCountry renders the insert form
func InsertCountry() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("countrys/insert", fiber.Map{
			"Title": "Add New Country",
		}, "layouts/main")
	}
}

// CreateCountry handles the form submission for creating a new Country
func CreateCountry(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		country := new(models.Country)
		if err := c.BodyParser(country); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if result := db.Create(country); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Redirect("/countrys")
	}
}

// EditCountry renders the edit form for a specific Country
func EditCountry(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var country models.Country
		if err := db.First(&country, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Country not found",
			})
		}
		return c.Render("countrys/edit", fiber.Map{"country": country, "Title": "Edit Entry"}, "layouts/main")
	}
}

// UpdateCountry handles the form submission for updating a Country
func UpdateCountry(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var country models.Country
		if err := db.First(&country, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Country not found",
			})
		}
		if err := c.BodyParser(&country); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if err := db.Save(&country).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update Country",
			})
		}
		return c.Redirect("/countrys")
	}
}

// DeleteCountry renders the delete confirmation view for a specific Country
func DeleteCountry(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var country models.Country
		if err := db.First(&country, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Country not found",
			})
		}
		return c.Render("countrys/delete", fiber.Map{"country": country, "Title": "Delete Entry"}, "layouts/main")
	}
}

// DestroyCountry handles the deletion of a Country
func DestroyCountry(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var country models.Country
		if err := db.First(&country, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Country not found",
			})
		}
		if err := db.Unscoped().Delete(&country).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete Country",
			})
		}
		return c.JSON(fiber.Map{"redirectUrl": "/countrys"})
	}
}
