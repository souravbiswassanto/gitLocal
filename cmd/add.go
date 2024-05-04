package cmd

import (
	"github.com/souravbiswassanto/gitLocal/pkg"
	"github.com/spf13/cobra"
)

func NewAddCmd() *cobra.Command {
	var email string
	var paths []string
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a dir",
		Long:  "Add a dir for viewing the git local stats",
		Run: func(cmd *cobra.Command, args []string) {
			// pkg.Test()
			pkg.ShowLocalGitContrib(email, paths)
		},
	}
	addCmd.Flags().StringSliceVarP(&paths, "paths", "d", []string{}, "provide the paths")
	addCmd.Flags().StringVarP(&email, "email", "e", "x@gmail.com", "provide your email")
	return addCmd
}
