package cmd

import "github.com/spf13/cobra"

var migrationCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate command is used for database migration",
	Long:  "migrate command is used for database migration: migrate < up | down >",
}

func init() {
	rootCmd.AddCommand(migrationCmd)
}
