package cmd

import (
	"github.com/souravbiswassanto/gitLocal/pkg"
	"os"

	"github.com/spf13/cobra"
)

func RootCMD() *cobra.Command {
	rootCMD := &cobra.Command{
		Use:   "gitLocal",
		Short: "Git local",
		Long:  "Git Local Contribution",
		Run: func(cmd *cobra.Command, args []string) {
			pkg.Test()
		},
	}
	rootCMD.AddCommand(NewAddCmd())
	return rootCMD
}

func Execute() {
	err := RootCMD().Execute()
	if err != nil {
		os.Exit(1)
	}
}
