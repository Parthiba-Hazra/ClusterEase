/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/Parthiba-Hazra/clstres/cmd"
	_ "github.com/Parthiba-Hazra/clstres/cmd/create"
	_ "github.com/Parthiba-Hazra/clstres/cmd/details"
	_ "github.com/Parthiba-Hazra/clstres/cmd/details/deployments"
	_ "github.com/Parthiba-Hazra/clstres/cmd/details/namespace"
	_ "github.com/Parthiba-Hazra/clstres/cmd/details/pod"
	_ "github.com/Parthiba-Hazra/clstres/cmd/details/service"
	_ "github.com/Parthiba-Hazra/clstres/cmd/exec"
	_ "github.com/Parthiba-Hazra/clstres/cmd/show"
	_ "github.com/Parthiba-Hazra/clstres/cmd/show/allResources"
	_ "github.com/Parthiba-Hazra/clstres/cmd/show/deployments"
	_ "github.com/Parthiba-Hazra/clstres/cmd/show/namespaces"
	_ "github.com/Parthiba-Hazra/clstres/cmd/show/pods"
	_ "github.com/Parthiba-Hazra/clstres/cmd/show/services"
)

func main() {
	cmd.Execute()
}
