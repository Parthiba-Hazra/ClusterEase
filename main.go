/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/Parthiba-Hazra/clstres/cmd"
	_ "github.com/Parthiba-Hazra/clstres/cmd/show"
	_ "github.com/Parthiba-Hazra/clstres/cmd/show/pods"
)

func main() {
	cmd.Execute()
}
