package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func Execute() {
	if err := rootCmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func rootCmd() *cobra.Command {
	var pers string
	cmd := cobra.Command{
		Use:   "fe26",
		Short: "Todo",
		Long:  `Also to do`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
	cmd.PersistentFlags().StringVar(&pers, "pers", "", "persistent flag")

	cmd.AddCommand(
		versionCmd(),
		Server(),
	)

	return &cmd
}

func versionCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "version",
		Short: "version ",
		Long:  `version long`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			//per, _ := cmd.Flags().GetString("pers")
			//fmt.Println(per)
			fmt.Println("version information goes here")
		},
	}

	// hide persistent flag on this command
	cmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		_ = command.Flags().MarkHidden("pers")
		command.Parent().HelpFunc()(command, strings)
	})

	return &cmd
}
