package cmd

import(
	"github.com/spf13/cobra"
)

func NewAddCmd() {
	var email string
	addCmd := &cobra.Command{
		Use : "add",
		Short: "Add a dir",
		Long: "Add a dir for viewing the git local stats",
		Run: func(cmd *cobra.Command, args []string){

		},
	}
	addCmd.Flags().StringVarP(&email, "email", "e", "x@gmail.com", "provide your email")

}