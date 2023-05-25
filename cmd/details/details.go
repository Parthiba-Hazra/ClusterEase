/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package details

import (
	"log"

	"github.com/Parthiba-Hazra/clstres/cmd"
	"github.com/spf13/cobra"
)

// detailsCmd represents the details command
var DetailsCmd = &cobra.Command{
	Use:   "details",
	Short: "show kubernetes resources details",
	Long:  `It show a detail view of a kubernetes resoureces according to user input`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Print("please provide resource type to get details")
	},
}

func init() {
	cmd.RootCmd.AddCommand(DetailsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// detailsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// detailsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
