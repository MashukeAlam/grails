package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"grails/models" // Adjust the import path accordingly
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
		if result := db.Create(&game); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": result.Error.Error(),
			})
		}
		return c.Redirect("/game")
	}
}
