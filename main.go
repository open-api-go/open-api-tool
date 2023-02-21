package main

import (
	"log"

	"github.com/open-api-go/openapi-tool/command"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:     "openapi-tool new -n name -d dir",
	Short:   "openapi-tool: open-api toolkit",
	Long:    "openapi-tool: init open-api repository",
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
