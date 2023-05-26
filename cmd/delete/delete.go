/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/Parthiba-Hazra/clstres/client"
	"github.com/Parthiba-Hazra/clstres/cmd"
	"github.com/Parthiba-Hazra/clstres/helper"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete resource in the kubernetes cluster",
	Long:  `It deletes the resources based on user provided info.`,
	Run: func(cmd *cobra.Command, args []string) {
		namespace, _ := cmd.Flags().GetString("ns")
		resourceName, _ := cmd.Flags().GetString("rn")
		kind, _ := cmd.Flags().GetString("k")

		client, err := client.GetClient()
		if err != nil {
			log.Printf("error getting kubernetes client: %v", err)
		}

		err = helper.DeleteResource(client, kind, resourceName, namespace)
		if err != nil {
			log.Printf("error while deleting resource: %v", err)
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(deleteCmd)
	deleteCmd.PersistentFlags().String("k", "", "You need to provide the kind of the resource that you want to delete. (eg: --k=deployment)")
	deleteCmd.PersistentFlags().String("rn", "", "You need to provide the name of the resource that you want to delete. (eg: --rn=deployment-name)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
