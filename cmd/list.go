/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/michaelbui99/csbcu/internal/dotnet"
	"github.com/spf13/cobra"
)

var listQuiently bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all paths to folders that contains a file ending in .csproj",
	Long:  `List all paths to folders that contains a file ending in .csproj`,
	RunE: func(cmd *cobra.Command, args []string) error {
		currentDir, err := os.Getwd()
		if err != nil {
			return err
		}

		if !listQuiently {
			log.Printf("Searching for all c# project paths from %v ...", currentDir)
		}

		projectPaths, err := dotnet.GetProjectPaths(currentDir)
		if err != nil {
			return err
		}

		if !listQuiently {
			log.Printf("Found %v projects", len(projectPaths))
		}

		for _, projectPath := range projectPaths {
			os.Stdout.WriteString(projectPath + "\n")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&listQuiently, "quiet", "q", false, "Disables all logging. Can be useful when you want to pipe this into the 'clean' command: cscbu list | cscbu clean")
}
