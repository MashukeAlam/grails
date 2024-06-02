package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "gen",
				Aliases: []string{"g"},
				Usage:   "generate necessary files to start a project",
				Action: func(cCtx *cli.Context) error {
					// Read project name
					fmt.Print("Enter project name: ")
					projectName := readInput()

					// Read database name
					fmt.Print("Enter database name: ")
					databaseName := readInput()

					fmt.Printf("Project Name: %s\n", projectName)
					fmt.Printf("Database Name: %s\n", databaseName)

					repoURL := "https://github.com/MashukeAlam/grails-template.git"

					// Create the project directory
					projectDir := filepath.Join(".", projectName)
					if err := os.MkdirAll(projectDir, os.ModePerm); err != nil {
						log.Fatalf("Failed to create project directory: %v", err)
					}

					// Clone the repository into the project directory
					fmt.Printf("%s📂 Cloning repository...%s\n", yellow, reset)
					cmd := exec.Command("git", "clone", repoURL, projectName)
					cmd.Stdout = nil
					cmd.Stderr = nil
					err := cmd.Run()
					if err != nil {
						log.Fatalf("%s❌ Failed to clone the repository: %v%s\n", red, err, reset)
					}
					fmt.Printf("%s✅ Repository cloned successfully!%s\n", green, reset)

					// Change directory to the cloned project
					err = os.Chdir(projectName)
					if err != nil {
						log.Fatalf("%s❌ Failed to change directory to %s: %v%s\n", red, projectName, err, reset)
					}

					// Edit the module name using go mod edit
					cmd = exec.Command("go", "mod", "edit", "-module", projectName)
					err = cmd.Run()
					if err != nil {
						log.Fatalf("%s❌ Failed to edit module name: %v%s\n", red, err, reset)
					}
					fmt.Printf("%s✅ Project setup complete!%s\n", green, reset)

					// Provide instructions to the user
					fmt.Printf("%sOne step that is left: %s", yellow, reset)
					fmt.Printf("1. %s🔄 Running 'go mod tidy'...%s\n", yellow, reset)
					fmt.Printf("%s🚀 All set! Happy coding!%s\n", green, reset)
					return nil
				},
			},
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "show the version",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("GRAILS Version: 0.9.2beta")
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	return strings.TrimSpace(input)
}
