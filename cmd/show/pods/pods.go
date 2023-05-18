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

// podsCmd represents the pods command
var podsCmd = &cobra.Command{
	Use:   "pods",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := client.GetClient()
		if err != nil {
			log.Printf("error getting kubernetes client: %v", err)
		}
		podDetails, err := helper.ShowPod(client)
		if err != nil {
			log.Printf("error getting pods: %v", err)
		} else {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Pod name", "Namespace", "status"})

			for _, pod := range podDetails {
				row := []string{pod.Name, pod.Namespace, pod.Status}
				table.Append(row)
			}
			table.Render()
		}
	},
}

func init() {
	show.ShowCmd.AddCommand(podsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// podsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// podsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
