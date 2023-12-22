package cmd

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
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

	versCdm := versionCmd()
	cmd.AddCommand(
		subCmd(),
		versCdm,
	)

	return &cmd
}

func subCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "sub",
		Short: "Todo",
		Long:  `Also to do`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			per, _ := cmd.Flags().GetString("pers")
			fmt.Println(per)
		},
	}

	return &cmd
}

func versionCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "version",
		Short: "Todo",
		Long:  `Also to do`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			//per, _ := cmd.Flags().GetString("pers")
			//fmt.Println(per)
			fmt.Println("version information goes here")
		},
	}
	spew.Dump()

	return &cmd
}
