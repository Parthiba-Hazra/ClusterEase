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

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Get details of a Kubernetes service",
	Long: `Retrieve detailed information about a Kubernetes service, including its name, namespace,
creation time, labels, type, cluster IP, external IPs, load balancer IP, ports, selector,
and session affinity.`,
	Run: func(cmd *cobra.Command, args []string) {
		namespace, _ := cmd.Flags().GetString("ns")
		serviceName, _ := cmd.Flags().GetString("s")

		client, err := client.GetClient()
		if err != nil {
			log.Printf("error getting kubernetes client: %v", err)
		}
		serviceDetailsList, err := helper.ShowServiceDetails(client, namespace, serviceName)
		if err != nil {
			log.Printf("error getting service details: %v", err)
		} else {
			for _, service := range serviceDetailsList {
				fmt.Println("Service Name:", service.Name)
				fmt.Println("Namespace:", service.Namespace)
				fmt.Println("Creation Time:", service.CreationTime)
				fmt.Println("Labels:", service.Labels)
				fmt.Println("Type:", service.Type)
				fmt.Println("Cluster IP:", service.ClusterIP)
				fmt.Println("External IPs:", service.ExternalIPs)
				fmt.Println("LoadBalancer IP:", service.LoadBalancerIP)

				fmt.Println("Ports:")
				for _, port := range service.Ports {
					fmt.Println("  - Name:", port.Name)
					fmt.Println("    Protocol:", port.Protocol)
					fmt.Println("    Port:", port.Port)
					fmt.Println("    Target Port:", port.TargetPort)
					fmt.Println("    Node Port:", port.NodePort)
				}

				fmt.Println("Selector:", service.Selector)
				fmt.Println("Session Affinity:", service.SessionAffinity)

				fmt.Println("---------------------------")
			}
		}
	},
}

func init() {
	details.DetailsCmd.AddCommand(serviceCmd)
	serviceCmd.PersistentFlags().String("s", "", "You need to provide the name of pod in order to get details of that perticular pod (eg: --s=service-name)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
