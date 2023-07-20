/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"log"
	"os"

	"github.com/michaelbui99/csbcu/internal/dotnet"
	"github.com/spf13/cobra"
)

var cleanQuietly bool

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean <project path>...",
	Short: "Clean obj and bin folder from project(s)",
	Long:  `Clean obj and bin folder from project(s)`,
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		stat, err := os.Stdin.Stat()
		if err != nil {
			return err
		}

		projectPaths := make([]string, 0)
		// project paths passed via pipe
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				projectPaths = append(projectPaths, scanner.Text())
			}
		}

		// project paths passed as arguments
		projectPaths = append(projectPaths, args...)

		projectsCleaned, err := dotnet.CleanBinaries(projectPaths)
		if err != nil {
			return err
		}

		if !cleanQuietly {
			log.Printf("Cleaned %v projects", projectsCleaned)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
	cleanCmd.Flags().BoolVarP(&cleanQuietly, "quiet", "q", false, "Disables all logging")
}
