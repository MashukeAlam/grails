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
        // Render the Slim template with the games data
        return c.Render("games/index", fiber.Map{
            "games": games,
            "title": "All Games",
        })
    }
}
