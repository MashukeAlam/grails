package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"grails/models" // Adjust the import path accordingly
)

// GetCars retrieves all Cars from the database
func GetCars(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var Cars []models.Car
		if result := db.Find(&Cars); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Render("cars/index", fiber.Map{
			"Title": "All Cars",
			"Cars": Cars,
		}, "layouts/main")
	}
}

// InsertCar renders the insert form
func InsertCar() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("cars/insert", fiber.Map{
			"Title": "Add New Car",
		}, "layouts/main")
	}
}

// CreateCar handles the form submission for creating a new Car
func CreateCar(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		car := new(models.Car)
		if err := c.BodyParser(car); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if result := db.Create(car); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Redirect("/cars")
	}
}

// EditCar renders the edit form for a specific Car
func EditCar(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var car models.Car
		if err := db.First(&car, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Car not found",
			})
		}
		return c.Render("cars/edit", fiber.Map{"car": car, "Title": "Edit Entry"}, "layouts/main")
	}
}

// UpdateCar handles the form submission for updating a Car
func UpdateCar(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var car models.Car
		if err := db.First(&car, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Car not found",
			})
		}
		if err := c.BodyParser(&car); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if err := db.Save(&car).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update Car",
			})
		}
		return c.Redirect("/cars")
	}
}

// DeleteCar renders the delete confirmation view for a specific Car
func DeleteCar(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var car models.Car
		if err := db.First(&car, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Car not found",
			})
		}
		return c.Render("cars/delete", fiber.Map{"car": car, "Title": "Delete Entry"}, "layouts/main")
	}
}

// DestroyCar handles the deletion of a Car
func DestroyCar(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var car models.Car
		if err := db.First(&car, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Car not found",
			})
		}
		if err := db.Unscoped().Delete(&car).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete Car",
			})
		}
		return c.JSON(fiber.Map{"redirectUrl": "/cars"})
	}
}
