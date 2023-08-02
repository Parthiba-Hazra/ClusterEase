/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/Parthiba-Hazra/clstres/client"
	"github.com/Parthiba-Hazra/clstres/cmd/details"
	"github.com/Parthiba-Hazra/clstres/helper"
	"github.com/spf13/cobra"
)

// namespaceCmd represents the namespace command
var namespaceCmd = &cobra.Command{
	Use:   "namespace",
	Short: "Show details of a namespace",
	Long: `Show details of a namespace including its name, creation time,
status, labels, annotations, and associated resource quota.`,
	Run: func(cmd *cobra.Command, args []string) {
		namespace, _ := cmd.Flags().GetString("ns")

		client, err := client.GetClient()
		if err != nil {
			log.Printf("error getting kubernetes client: %v", err)
			return
		}

		nsDetailsList, err := helper.GetNamespaceDetails(client, namespace)
		if err != nil {
			log.Printf("error getting namespace details: %v", err)
			return
		}

		for _, ns := range nsDetailsList {
			fmt.Println("Name:", ns.Name)
			fmt.Println("Creation Time:", ns.CreationTime)
			fmt.Println("Status:", ns.Status)
			fmt.Println("Labels:", ns.Labels)
			fmt.Println("Annotations:", ns.Annotations)
			fmt.Println("Resource Quota:", ns.ResourceQuota)
			fmt.Println("-----------------------------------")
		}
	},
}

func init() {
	details.DetailsCmd.AddCommand(namespaceCmd)
	namespaceCmd.PersistentFlags().String("ns", "", "Provide the name of the namespace to get its details (e.g., --ns=namespace-name)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// namespaceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// namespaceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
