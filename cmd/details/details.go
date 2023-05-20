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
	Short: "show a pod's details",
	Long:  `It provide all details of a pod`,
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
