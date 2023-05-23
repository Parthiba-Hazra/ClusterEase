/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package show

import (
	"github.com/Parthiba-Hazra/clstres/cmd"
	"github.com/emicklei/go-restful/v3/log"
	"github.com/spf13/cobra"
)

var Verbose bool
var Source string

// showCmd represents the show command
var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "It show the rosurce details metion in the command (like.. pods, servcies, deployments, etc...)",
	Long:  `It show the rosurce details metion in the command (like.. pods, servcies, deployments, etc...)`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Print("Please mention the name of resorces you want to see")
	},
}

func init() {
	cmd.RootCmd.AddCommand(ShowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
