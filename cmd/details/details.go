/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package details

import (
	"fmt"

	"github.com/Parthiba-Hazra/clstres/cmd"
	"github.com/spf13/cobra"
)

// detailsCmd represents the details command
var detailsCmd = &cobra.Command{
	Use:   "details",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("details called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(detailsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// detailsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// detailsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
