package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/MashukeAlam/grails/models" // Adjust the import path accordingly
)

// GetMobiles retrieves all Mobiles from the database
func GetMobiles(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var Mobiles []models.Mobile
		if result := db.Find(&Mobiles); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Render("mobiles/index", fiber.Map{
			"Title": "All Mobiles",
			"Mobiles": Mobiles,
		}, "layouts/main")
	}
}

// InsertMobile renders the insert form
func InsertMobile() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("mobiles/insert", fiber.Map{
			"Title": "Add New Mobile",
		}, "layouts/main")
	}
}

// CreateMobile handles the form submission for creating a new Mobile
func CreateMobile(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		mobile := new(models.Mobile)
		if err := c.BodyParser(mobile); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if result := db.Create(mobile); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Redirect("/mobiles")
	}
}

// EditMobile renders the edit form for a specific Mobile
func EditMobile(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var mobile models.Mobile
		if err := db.First(&mobile, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Mobile not found",
			})
		}
		return c.Render("mobiles/edit", fiber.Map{"mobile": mobile, "Title": "Edit Entry"}, "layouts/main")
	}
}

// UpdateMobile handles the form submission for updating a Mobile
func UpdateMobile(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var mobile models.Mobile
		if err := db.First(&mobile, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Mobile not found",
			})
		}
		if err := c.BodyParser(&mobile); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if err := db.Save(&mobile).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update Mobile",
			})
		}
		return c.Redirect("/mobiles")
	}
}

// DeleteMobile renders the delete confirmation view for a specific Mobile
func DeleteMobile(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var mobile models.Mobile
		if err := db.First(&mobile, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Mobile not found",
			})
		}
		return c.Render("mobiles/delete", fiber.Map{"mobile": mobile, "Title": "Delete Entry"}, "layouts/main")
	}
}

// DestroyMobile handles the deletion of a Mobile
func DestroyMobile(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var mobile models.Mobile
		if err := db.First(&mobile, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Mobile not found",
			})
		}
		if err := db.Unscoped().Delete(&mobile).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete Mobile",
			})
		}
		return c.JSON(fiber.Map{"redirectUrl": "/mobiles"})
	}
}
