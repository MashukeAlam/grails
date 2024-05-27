package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"grails/models" // Adjust the import path accordingly
)

// GetPerfumes retrieves all Perfumes from the database
func GetPerfumes(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var Perfumes []models.Perfume
		if result := db.Find(&Perfumes); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Render("perfumes/index", fiber.Map{
			"Title": "All Perfumes",
			"Records": Perfumes,
		}, "layouts/main")
	}
}

// InsertPerfume renders the insert form
func InsertPerfume() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("perfumes/insert", fiber.Map{
			"Title": "Add New Perfume",
		}, "layouts/main")
	}
}

// CreatePerfume handles the form submission for creating a new Perfume
func CreatePerfume(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		perfume := new(models.Perfume)
		if err := c.BodyParser(perfume); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if result := db.Create(perfume); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Redirect("/perfumes")
	}
}

// EditPerfume renders the edit form for a specific Perfume
func EditPerfume(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var perfume models.Perfume
		if err := db.First(&perfume, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Perfume not found",
			})
		}
		return c.Render("perfumes/edit", fiber.Map{"perfume": perfume, "Title": "Edit Entry"}, "layouts/main")
	}
}

// UpdatePerfume handles the form submission for updating a Perfume
func UpdatePerfume(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var perfume models.Perfume
		if err := db.First(&perfume, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Perfume not found",
			})
		}
		if err := c.BodyParser(&perfume); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if err := db.Save(&perfume).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update Perfume",
			})
		}
		return c.Redirect("/perfumes")
	}
}

// DeletePerfume renders the delete confirmation view for a specific Perfume
func DeletePerfume(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var perfume models.Perfume
		if err := db.First(&perfume, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Perfume not found",
			})
		}
		return c.Render("perfumes/delete", fiber.Map{"perfume": perfume, "Title": "Delete Entry"}, "layouts/main")
	}
}

// DestroyPerfume handles the deletion of a Perfume
func DestroyPerfume(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var perfume models.Perfume
		if err := db.First(&perfume, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Perfume not found",
			})
		}
		if err := db.Unscoped().Delete(&perfume).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete Perfume",
			})
		}
		return c.JSON(fiber.Map{"redirectUrl": "/perfumes"})
	}
}
