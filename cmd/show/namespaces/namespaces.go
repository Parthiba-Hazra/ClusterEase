/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/Parthiba-Hazra/clstres/client"
	"github.com/Parthiba-Hazra/clstres/cmd/show"
	"github.com/Parthiba-Hazra/clstres/helper"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// namespacesCmd represents the namespaces command
var namespacesCmd = &cobra.Command{
	Use:   "namespaces",
	Short: "It will show all name-spaces in kubernetes cluster",
	Long:  `Show all namespace's details like name, status, age`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := client.GetClient()
		if err != nil {
			log.Printf("error getting kubernetes client: %v", err)
		}
		namespaceDetails, err := helper.ShowNameSpaces(client)
		if err != nil {
			log.Printf("Can't get the namespacces: %v", err)
		} else {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Namespace-Name", "status", "Age"})

			for _, namespace := range namespaceDetails {
				row := []string{namespace.Name, namespace.Status, namespace.Age}
				table.Append(row)
			}
			table.Render()
		}
	},
}

func init() {
	show.ShowCmd.AddCommand(namespacesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// namespacesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// namespacesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
