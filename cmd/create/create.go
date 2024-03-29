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

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Kubernetes resources from YAML",
	Long: `Create Kubernetes resources from a YAML file. This command allows you to deploy Kubernetes resources
based on the Kind specified in the YAML file. For example, to create a deployment, run the command as follows:
'clstres create --fp=./deployment.yaml --ns=default'`,
	Run: func(cmd *cobra.Command, args []string) {
		namespace, _ := cmd.Flags().GetString("ns")
		filePath, _ := cmd.Flags().GetString("fp")

		client, err := client.GetClient()
		if err != nil {
			log.Printf("error getting kubernetes client: %v", err)
		}
		err = helper.CreateResourcesFromYAML(client, namespace, filePath)
		if err != nil {
			log.Printf("error creating kubernetes resources: %v", err)
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(createCmd)
	createCmd.PersistentFlags().String("fp", "", "You need to provide the file path of your YAML file. (eg: --fp=./deployment.yaml)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
