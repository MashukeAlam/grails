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
	blue      = "\033[34m"
	magenta   = "\033[35m"
	cyan      = "\033[36m"
	bold      = "\033[1m"
	underline = "\033[4m"
	blink     = "\033[5m"
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

					fmt.Printf("%s%sSetup intializing project %s with %s db.%s\n", bold, magenta, projectName,databaseName, reset)

					repoURL := "https://github.com/MashukeAlam/grails-template.git"

					// Create the project directory
					projectDir := filepath.Join(".", projectName)
					if err := os.MkdirAll(projectDir, os.ModePerm); err != nil {
						log.Fatalf("Failed to create project directory: %v", err)
					}

					// Clone the repository into the project directory
					fmt.Printf("%s%s📂 Cloning repository...%s\n", bold, cyan, reset)
					cmd := exec.Command("git", "clone", repoURL, projectName)
					cmd.Stdout = nil
					cmd.Stderr = nil
					err := cmd.Run()
					if err != nil {
						log.Fatalf("%s❌ Failed to clone the repository: %v%s\n", red, err, reset)
					}
					fmt.Printf("%s%s✅ Repository cloned successfully!%s\n", bold, green, reset)

					// Change directory to the cloned project
					err = os.Chdir(projectName)
					if err != nil {
						log.Fatalf("%s❌ Failed to change directory to %s: %v%s\n", red, projectName, err, reset)
					}

					// Edit the module name using go mod edit
					// cmd = exec.Command("go", "mod", "edit", "-module", projectName)
					// err = cmd.Run()
					// if err != nil {
					// 	log.Fatalf("%s❌ Failed to edit module name: %v%s\n", red, err, reset)
					// }
					fmt.Printf("%s%s📄 Final touch...%s\n", blink, yellow, reset)

					// Ask the user if they want to run 'go mod tidy'
					reader := bufio.NewReader(os.Stdin)
					fmt.Printf("%s%s%s❓ Do you want to run 'go mod tidy'? (y/n): %s", bold, underline, yellow, reset)
					response, err := reader.ReadString('\n')
					if err != nil {
						log.Fatalf("%s❌ Failed to read input: %v%s\n", red, err, reset)
					}
					response = strings.TrimSpace(response)

					if strings.ToLower(response) == "y" {
						fmt.Printf("%s🔄 Running 'go mod tidy'...%s\n", yellow, reset)
						// Run 'go mod tidy'
						cmd = exec.Command("go", "mod", "tidy")
						cmd.Stdout = nil
						cmd.Stderr = nil
						err = cmd.Run()
						if err != nil {
							log.Fatalf("%s❌ Failed to run 'go mod tidy': %v%s\n", red, err, reset)
						}
						fmt.Printf("%s✅ 'go mod tidy' completed successfully!%s\n", green, reset)
					} else {
						fmt.Printf("%s🚫 Skipped 'go mod tidy'%s\n", yellow, reset)
					}
					// Git commit the changes
					cmd = exec.Command("git", "add", "go.mod", "go.sum")
					err = cmd.Run()
					if err != nil {
						log.Fatalf("%s❌ Failed to stage changes: %v%s\n", red, err, reset)
					}
					cmd = exec.Command("git", "commit", "-m", "Grails project setup and module renamed")
					err = cmd.Run()
					if err != nil {
						log.Fatalf("%s❌ Failed to commit changes: %v%s\n", red, err, reset)
					}
					fmt.Printf("%s%s✅ Project setup complete!%s\n", bold, green, reset)

					// Provide instructions to the user
					fmt.Printf("%sTo get going %s\n", magenta, reset)
					fmt.Printf("\t1. %s cd %s %s\n", cyan, projectName, reset)
					fmt.Printf("\t2. %s go run app.go %s\n", cyan, reset)
					fmt.Printf("%s%s🚀 All set! Happy coding!%s\n", bold, green, reset)
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