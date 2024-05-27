package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"grails/models" // Adjust the import path accordingly
)

// GetCards retrieves all Cards from the database
func GetCards(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var Cards []models.Card
		if result := db.Find(&Cards); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Render("cards/index", fiber.Map{
			"Title": "All Cards",
			"Records": Cards,
		}, "layouts/main")
	}
}

// InsertCard renders the insert form
func InsertCard() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("cards/insert", fiber.Map{
			"Title": "Add New Card",
		}, "layouts/main")
	}
}

// CreateCard handles the form submission for creating a new Card
func CreateCard(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		card := new(models.Card)
		if err := c.BodyParser(card); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if result := db.Create(card); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Redirect("/cards")
	}
}

// EditCard renders the edit form for a specific Card
func EditCard(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var card models.Card
		if err := db.First(&card, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Card not found",
			})
		}
		return c.Render("cards/edit", fiber.Map{"card": card, "Title": "Edit Entry"}, "layouts/main")
	}
}

// UpdateCard handles the form submission for updating a Card
func UpdateCard(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var card models.Card
		if err := db.First(&card, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Card not found",
			})
		}
		if err := c.BodyParser(&card); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if err := db.Save(&card).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update Card",
			})
		}
		return c.Redirect("/cards")
	}
}

// DeleteCard renders the delete confirmation view for a specific Card
func DeleteCard(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var card models.Card
		if err := db.First(&card, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Card not found",
			})
		}
		return c.Render("cards/delete", fiber.Map{"card": card, "Title": "Delete Entry"}, "layouts/main")
	}
}

// DestroyCard handles the deletion of a Card
func DestroyCard(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var card models.Card
		if err := db.First(&card, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Card not found",
			})
		}
		if err := db.Unscoped().Delete(&card).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete Card",
			})
		}
		return c.JSON(fiber.Map{"redirectUrl": "/cards"})
	}
}
