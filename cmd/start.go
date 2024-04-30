package cmd

import "github.com/spf13/cobra"

rootCMD = &cobra.Command{
	Use: "gitLocal",
	Short: "Git local",
	Long: "Git Local Contribution"
}

func Execute() {
	err := rootCMD.Execute();
	if err != nil {
		os.Exit(1)
	}
}