/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/Parthiba-Hazra/clstres/cmd"
	"github.com/Parthiba-Hazra/clstres/helper"
	"github.com/spf13/cobra"
)

// podEnterCmd represents the podEnter command
var podEnterCmd = &cobra.Command{
	Use:   "podEnter",
	Short: "To get into a specific pod",
	Long:  `'podEnter' command allows you to execute commands within a pod running in a Kubernetes cluster. It provides a convenient way to interact with containers and perform tasks such as running scripts, troubleshooting, or debugging`,
	Run: func(cmd *cobra.Command, args []string) {
		namespace, _ := cmd.Flags().GetString("ns")
		podName, _ := cmd.Flags().GetString("p")
		command, _ := cmd.Flags().GetString("cmd")

		ans2, ans1, err := helper.EnterInPod(namespace, podName, command)
		if err != nil {
			log.Printf("EnterInPod returned an error: %v", err)
		}
		log.Print(ans1)
		log.Print(ans2)
	},
}

func init() {
	cmd.RootCmd.AddCommand(podEnterCmd)
	podEnterCmd.PersistentFlags().String("p", "", "You need to provide the name of pod in order to get details of that perticular pod (eg: --p=pod-name)")
	podEnterCmd.PersistentFlags().String("cmd", "", "You have to provide the command that you want to execute under the pod (eg: --cmd=ls)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// podEnterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// podEnterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
