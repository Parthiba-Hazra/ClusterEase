/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package details

import (
	"fmt"
	"log"

	"github.com/Parthiba-Hazra/clstres/client"
	"github.com/Parthiba-Hazra/clstres/cmd/details"
	"github.com/Parthiba-Hazra/clstres/helper"
	"github.com/spf13/cobra"
)

// deploymentCmd represents the deployment command
var deploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Show details of a Kubernetes deployment",
	Long: `Show detailed information about a Kubernetes deployment in the specified namespace. This command
provides information such as creation time, replicas, available replicas, ready replicas, updated replicas,
strategy, selector, and container details for the specified deployment.`,
	Run: func(cmd *cobra.Command, args []string) {
		namespace, _ := cmd.Flags().GetString("ns")
		deploymentName, _ := cmd.Flags().GetString("d")

		client, err := client.GetClient()
		if err != nil {
			log.Printf("error getting kubernetes client: %v", err)
		}
		deploymentDetailsList, err := helper.GetDeploymentDetails(client, namespace, deploymentName)
		if err != nil {
			log.Printf("error getting deployment details: %v", err)
		} else {
			for _, deployment := range deploymentDetailsList {
				fmt.Println("Name:", deployment.Name)
				fmt.Println("Namespace:", deployment.Namespace)
				fmt.Println("Creation Time:", deployment.CreationTime)
				fmt.Println("Replicas:", deployment.Replicas)
				fmt.Println("Available Replicas:", deployment.AvailableReplicas)
				fmt.Println("Ready Replicas:", deployment.ReadyReplicas)
				fmt.Println("Updated Replicas:", deployment.UpdatedReplicas)
				fmt.Println("Strategy:", deployment.Strategy)
				fmt.Println("Selector:", deployment.Selector)

				fmt.Println("Containers:")
				for _, container := range deployment.Containers {
					fmt.Println("\tContainer Name:", container.ContainerName)
					fmt.Println("\tPorts:")
					for _, port := range container.Ports {
						fmt.Println("\t\tPort Name:", port.PortName)
						fmt.Println("\t\tProtocol:", port.Protocol)
						fmt.Println("\t\tContainer Port:", port.ContainerPort)
					}
				}

				fmt.Println("-----------------------------------")
			}

		}
	},
}

func init() {
	details.DetailsCmd.AddCommand(deploymentCmd)
	deploymentCmd.PersistentFlags().String("d", "", "You need to provide the name of deployment to get details (eg: --d=deployment-name)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deploymentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deploymentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
