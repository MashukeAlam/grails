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
        return c.JSON(games)
    }
}
