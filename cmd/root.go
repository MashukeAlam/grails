package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "grails",
	Short: "A brief description of your application",
	// Add other properties as needed
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Add subcommands here
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(generateCmd)
	// Add other commands similarly
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize project directories",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Directories created successfully.")
		// Add logic to initialize directories
	},
}

var generateCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate project structure",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generating project structure...")

		dirs := []string{
			"internals",
			"models",
			"handlers",
			"helpers",
			"views",
			"static",
		}

		for _, dir := range dirs {
			err := os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				fmt.Printf("Failed to create directory %s: %v\n", dir, err)
			} else {
				fmt.Printf("Directory %s created successfully.\n", dir)
			}
		}
	},
}

func appendRoutesCode1() error {
	// Define the file path
	filePath := "habijabi/habijabi.txt"
	// Check if the file already exists
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		// Create the file if it doesn't exist
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Sprintf("%s", err)
			return err
		}
		defer file.Close()
	}

	// Read the existing content of the file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Convert content to string
	fileContent := string(content)

	codeToAdd := fmt.Sprint("func appendRoutesCode(codeToAdd string) error {\n\t// Define the file path\n\tfilePath := \"internals/routes.go\"\n\n\t// Check if the file already exists\n\t_, err := os.Stat(filePath)\n\tif os.IsNotExist(err) {\n\t\t// Create the file if it doesn't exist\n\t\tfile, err := os.Create(filePath)\n\t\tif err != nil {\n\t\t\treturn err\n\t\t}\n\t\tdefer file.Close()\n\t}\n\n\t// Read the existing content of the file\n\tcontent, err := ioutil.ReadFile(filePath)\n\tif err != nil {\n\t\treturn err\n\t}\n\n\t// Convert content to string\n\tfileContent := string(content)\n\n\t// Find the index of closing braces of SetupRoutes function\n\tidx := strings.LastIndex(fileContent, \"}\")\n\n\t// Append the codeToAdd before the closing braces\n\tnewContent := fileContent[:idx] + codeToAdd + \"\\n}\" + fileContent[idx+1:]\n\n\t// Write the updated content back to the file\n\terr = ioutil.WriteFile(filePath, []byte(newContent), 0644)\n\tif err != nil {\n\t\treturn err\n\t}\n\n\treturn nil\n}")

	// Find the index of closing braces of SetupRoutes function
	idx := strings.LastIndex(fileContent, "}")

	// Append the codeToAdd before the closing braces
	newContent := fileContent[:idx] + codeToAdd + "\n}" + fileContent[idx+1:]

	// Write the updated content back to the file
	err = ioutil.WriteFile(filePath, []byte(newContent), 0644)
	if err != nil {
		return err
	}
	fmt.Sprintf("here")
	return nil
}

// Define other commands similarly
