package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/MashukeAlam/grails/models" // Adjust the import path accordingly
)

// GetGames retrieves all games from the database
func GetGames(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var games []models.Game
		if result := db.Find(&games); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Render("games/index", fiber.Map{
			"Title": "All Games",
			"Games": games,
		}, "layouts/main")
	}
}

// InsertGame renders the insert form
func InsertGame() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Render("games/insert", fiber.Map{
			"Title": "Add New Game",
		}, "layouts/main")
	}
}

// CreateGame handles the form submission for creating a new game
func CreateGame(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		game := new(models.Game)
		if err := c.BodyParser(game); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		fmt.Println("Create")
		if result := db.Create(&game); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Redirect("/game")
	}
}

// EditGame renders the edit form for a specific game
func EditGame(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var game models.Game
		if err := db.First(&game, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Game not found",
			})
		}
		return c.Render("games/edit", fiber.Map{"game": game, "Title": "Edit Entry"}, "layouts/main")
	}
}

// UpdateGame handles the form submission for updating a game
func UpdateGame(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var game models.Game
		if err := db.First(&game, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Game not found",
			})
		}
		if err := c.BodyParser(&game); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}
		if err := db.Save(&game).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update game",
			})
		}
		fmt.Println("Edit")
		return c.Redirect("/game")
	}
}

// DeleteGame renders the delete confirmation view for a specific game
func DeleteGame(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var game models.Game
		if err := db.First(&game, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Game not found",
			})
		}
		fmt.Println("Delete")
		return c.Render("games/delete", fiber.Map{"game": game, "Title": "Delete Entry"}, "layouts/main")
	}
}

// DestroyGame handles the deletion of a game
func DestroyGame(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var game models.Game
		fmt.Println("Deleting game with ID:", c.Params("id"))
		if err := db.First(&game, c.Params("id")).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Game not found",
			})
		}
		fmt.Println("Delete")
		if err := db.Unscoped().Delete(&game).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete game",
			})
		}

		return c.JSON(fiber.Map{"redirectUrl": "/game"})
	}
}
