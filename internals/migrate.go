package internals

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func getDatabaseURL() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {
		fmt.Println("ENV Loaded.")
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, name)
}

func runGoose(direction string) error {
	dbURL := getDatabaseURL()
	cmd := exec.Command("goose", "-dir", "./migrations", "mysql", dbURL, direction)
	cmd.Env = append(cmd.Env, os.Environ()...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run migrations with goose: %v, output: %s", err, string(out))
	}
	log.Printf("Migrations executed successfully with goose:\n%s", string(out))
	return nil
}

func MigrateOld(direction string) {
	if direction != "up" && direction != "down" {
		fmt.Println("Invalid migration direction. Use 'up' or 'down'.")
		os.Exit(1)
	}

	err := runGoose(direction)
	if err != nil {
		log.Fatalf("Error running migrations with goose: %v", err)
	}
}
