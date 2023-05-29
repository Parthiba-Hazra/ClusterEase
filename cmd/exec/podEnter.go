/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Parthiba-Hazra/clstres/cmd"
	"github.com/spf13/cobra"
)

// podEnterCmd represents the podEnter command
var podEnterCmd = &cobra.Command{
	Use:   "podEnter",
	Short: "To get into a specific pod",
	Long:  `'podEnter' command allows you to execute commands within a pod running in a Kubernetes cluster. It provides a convenient way to interact with containers and perform tasks such as running scripts, troubleshooting, or debugging`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("podEnter called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(podEnterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// podEnterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// podEnterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
