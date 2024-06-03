package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/MashukeAlam/grails/models" // Adjust the import path accordingly
)

// GetVats retrieves all Vats from the database
func GetVats(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var Vats []models.Vat
		if result := db.Find(&Vats); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Render("vats/index", fiber.Map{
			"Title": "All Vats",
			"Records": Vats,
		}, "layouts/main")
	}
}

// InsertVat renders the insert form
func InsertVat() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("vats/insert", fiber.Map{
			"Title": "Add New Vat",
		}, "layouts/main")
	}
}

// CreateVat handles the form submission for creating a new Vat
func CreateVat(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		vat := new(models.Vat)
		if err := c.BodyParser(vat); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if result := db.Create(vat); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Redirect("/vats")
	}
}

// EditVat renders the edit form for a specific Vat
func EditVat(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var vat models.Vat
		if err := db.First(&vat, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Vat not found",
			})
		}
		return c.Render("vats/edit", fiber.Map{"vat": vat, "Title": "Edit Entry"}, "layouts/main")
	}
}

// UpdateVat handles the form submission for updating a Vat
func UpdateVat(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var vat models.Vat
		if err := db.First(&vat, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Vat not found",
			})
		}
		if err := c.BodyParser(&vat); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if err := db.Save(&vat).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update Vat",
			})
		}
		return c.Redirect("/vats")
	}
}

// DeleteVat renders the delete confirmation view for a specific Vat
func DeleteVat(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var vat models.Vat
		if err := db.First(&vat, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Vat not found",
			})
		}
		return c.Render("vats/delete", fiber.Map{"vat": vat, "Title": "Delete Entry"}, "layouts/main")
	}
}

// DestroyVat handles the deletion of a Vat
func DestroyVat(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var vat models.Vat
		if err := db.First(&vat, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Vat not found",
			})
		}
		if err := db.Unscoped().Delete(&vat).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete Vat",
			})
		}
		return c.JSON(fiber.Map{"redirectUrl": "/vats"})
	}
}
