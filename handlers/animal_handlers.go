package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"grails/models" // Adjust the import path accordingly
)

// GetAnimals retrieves all Animals from the database
func GetAnimals(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var Animals []models.Animal
		if result := db.Find(&Animals); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Render("animals/index", fiber.Map{
			"Title": "All Animals",
			"Animals": Animals,
		}, "layouts/main")
	}
}

// InsertAnimal renders the insert form
func InsertAnimal() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("animals/insert", fiber.Map{
			"Title": "Add New Animal",
		}, "layouts/main")
	}
}

// CreateAnimal handles the form submission for creating a new Animal
func CreateAnimal(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		animal := new(models.Animal)
		if err := c.BodyParser(animal); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if result := db.Create(animal); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Redirect("/animals")
	}
}

// EditAnimal renders the edit form for a specific Animal
func EditAnimal(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var animal models.Animal
		if err := db.First(&animal, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Animal not found",
			})
		}
		return c.Render("animals/edit", fiber.Map{"animal": animal, "Title": "Edit Entry"}, "layouts/main")
	}
}

// UpdateAnimal handles the form submission for updating a Animal
func UpdateAnimal(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var animal models.Animal
		if err := db.First(&animal, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Animal not found",
			})
		}
		if err := c.BodyParser(&animal); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if err := db.Save(&animal).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update Animal",
			})
		}
		return c.Redirect("/animals")
	}
}

// DeleteAnimal renders the delete confirmation view for a specific Animal
func DeleteAnimal(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var animal models.Animal
		if err := db.First(&animal, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Animal not found",
			})
		}
		return c.Render("animals/delete", fiber.Map{"animal": animal, "Title": "Delete Entry"}, "layouts/main")
	}
}

// DestroyAnimal handles the deletion of a Animal
func DestroyAnimal(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var animal models.Animal
		if err := db.First(&animal, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Animal not found",
			})
		}
		if err := db.Unscoped().Delete(&animal).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete Animal",
			})
		}
		return c.JSON(fiber.Map{"redirectUrl": "/animals"})
	}
}
