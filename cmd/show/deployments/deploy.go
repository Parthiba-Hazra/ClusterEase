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

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Show deployments in a Kubernetes namespace",
	Long: `Display a list of deployments in a specified Kubernetes namespace, including their names,
namespaces, readiness status, and ages.`,
	Run: func(cmd *cobra.Command, args []string) {
		namespace, _ := cmd.Flags().GetString("ns")
		client, err := client.GetClient()
		if err != nil {
			log.Printf("error getting kubernetes client: %v", err)
		}
		deploymentList, err := helper.ShowDeployments(client, namespace)
		if err != nil {
			log.Printf("error getting deployment list: %v", err)
		} else {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Deployment", "Namespace", "Ready", "Age"})

			for _, deployment := range deploymentList {
				row := []string{deployment.Name, deployment.Namespace, deployment.Ready, deployment.Age}
				table.Append(row)
			}
			table.Render()
		}
	},
}

func init() {
	show.ShowCmd.AddCommand(deployCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deployCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deployCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
