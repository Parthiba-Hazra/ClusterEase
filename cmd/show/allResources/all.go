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

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Get all resources from provided namespace",
	Long:  `It shows all resources in the provided namspace (eg: clstres show all --ns=default)`,
	Run: func(cmd *cobra.Command, args []string) {
		namespace, _ := cmd.Flags().GetString("ns")
		client, err := client.GetClient()
		if err != nil {
			log.Printf("error getting kubernetes client: %v", err)
		}
		resources, err := helper.GetAllResources(client, namespace)
		if err != nil {
			log.Printf("error getting resources: %v", err)
		} else {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Resource Type", "Name", "Namespace", "Created At"})

			for _, resource := range resources {
				createdTime := resource.CreatedAt.Format("2006-01-02 15:04:05")
				row := []string{resource.Kind, resource.Name, resource.Namespace, createdTime}
				table.Append(row)
			}
			table.Render()
		}
	},
}

func init() {
	show.ShowCmd.AddCommand(allCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
