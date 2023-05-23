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

// detailsCmd represents the details command
var detailsCmd = &cobra.Command{
	Use:   "pod",
	Short: "show a pod's details",
	Long:  `It a detail view of a pod`,
	Run: func(cmd *cobra.Command, args []string) {
		namespace, _ := cmd.Flags().GetString("ns")
		podName, _ := cmd.Flags().GetString("p")

		client, err := client.GetClient()
		if err != nil {
			log.Printf("error getting kubernetes client: %v", err)
		}
		podDetailsList, err := helper.GetPodDetails(client, namespace, podName)
		if err != nil {
			log.Printf("error getting pod details: %v", err)
		} else {
			for _, pod := range podDetailsList {
				fmt.Println("Name:", pod.Name)
				fmt.Println("Namespace:", pod.Namespace)
				fmt.Println("Creation Time:", pod.CreationTime)
				fmt.Println("Phase:", pod.Phase)
				fmt.Println("IP:", pod.IP)

				fmt.Println("Conditions:")
				for _, condition := range pod.Conditions {
					fmt.Println("\tType:", condition.Type)
					fmt.Println("\tStatus:", condition.Status)
					fmt.Println("\tLast Transition Time:", condition.LastTransitionTime)
					fmt.Println("\tReason:", condition.Reason)
					fmt.Println("\tMessage:", condition.Message)
				}

				fmt.Println("Container Details:")
				for _, container := range pod.ContainerDetails {
					fmt.Println("\tContainer Name:", container.ContainerName)
					fmt.Println("\tPorts:")
					for _, port := range container.Ports {
						fmt.Println("\t\tPort Name:", port.PortName)
						fmt.Println("\t\tProtocol:", port.Protocol)
						fmt.Println("\t\tContainer Port:", port.ContainerPort)
						fmt.Println("\t\tHost Port:", port.HostPort)
					}
				}

				fmt.Println("-----------------------------------")
			}

		}
	},
}

func init() {
	details.DetailsCmd.AddCommand(detailsCmd)
	detailsCmd.PersistentFlags().String("p", "", "You need to provide the name of pod in order to get details of that perticular pod (eg: --p=pod-name)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// detailsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// detailsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
