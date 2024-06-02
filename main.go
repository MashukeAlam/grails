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
					cmd := exec.Command("git", "clone", repoURL, projectDir)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					if err := cmd.Run(); err != nil {
						log.Fatalf("Failed to clone repository: %v", err)
					}

					// Change to the project directory
					if err := os.Chdir(projectDir); err != nil {
						log.Fatalf("Failed to change directory: %v", err)
					}

					// Update the module name in go.mod
					modCmd := exec.Command("go", "mod", "edit", "-module", projectName)
					modCmd.Stdout = os.Stdout
					modCmd.Stderr = os.Stderr
					if err := modCmd.Run(); err != nil {
						log.Fatalf("Failed to update module name: %v", err)
					}

					// Provide instructions to the user
					fmt.Println("\nProject setup complete!")
					fmt.Println("Next steps:")
					fmt.Println("1. Run `go mod tidy` to clean up the module dependencies.")
					fmt.Println("2. Start coding and enjoy your project!")
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
