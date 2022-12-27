package main

import (
	"log"

	"github.com/open-api-go/opa/command"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:     "opa",
	Short:   "opa: open-api toolkit",
	Long:    "opa: init open-api repository",
	Version: Version,
}

func Init() {
	cmd.AddCommand(command.CmdNew)
}

func main() {
	Init()
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
