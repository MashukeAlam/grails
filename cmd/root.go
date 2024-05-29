package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
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

// copyDir copies a whole directory recursively
func copyDir(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}

	for _, fd := range fds {
		srcfp := filepath.Join(src, fd.Name())
		dstfp := filepath.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = copyDir(srcfp, dstfp); err != nil {
				return err
			}
		} else {
			if err = copyFile(srcfp, dstfp); err != nil {
				return err
			}
		}
	}
	return nil
}

// copyFile copies a single file from src to dst
func copyFile(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}

var generateCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate project structure",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter project name: ")
		projectName, _ := reader.ReadString('\n')
		projectName = strings.TrimSpace(projectName)

		if projectName == "" {
			fmt.Println("Project name cannot be empty.")
			return
		}

		fmt.Printf("Generating project structure for %s...\n", projectName)

		dirs := []string{
			"internals",
			"models",
			"handlers",
			"helpers",
			"views",
			"static",
		}

		for _, dir := range dirs {
			fullPath := filepath.Join(projectName, dir)
			err := os.MkdirAll(fullPath, os.ModePerm)
			if err != nil {
				fmt.Printf("Failed to create directory %s: %v\n", fullPath, err)
			} else {
				fmt.Printf("Directory %s created successfully.\n", fullPath)
			}
		}

		// Assuming the source 'views' directory is located at "./grails/views"
		srcViews := "./grails/views"
		dstViews := filepath.Join(projectName, "views")

		err := copyDir(srcViews, dstViews)
		if err != nil {
			fmt.Printf("Failed to copy views directory: %v\n", err)
		} else {
			fmt.Printf("Views directory copied successfully to %s.\n", dstViews)
		}
	},
}
