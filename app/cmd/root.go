package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"runtime"
)

// Execute is the entry point for the command line
func Execute() {
	if err := newRootCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newRootCommand() *cobra.Command {
	var configFile = "./config.yaml"
	cmd := &cobra.Command{
		Use:   "fe26",
		Short: "fe26 is a basic http file server written in GO",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runServer(configFile)
		},
	}

	cmd.Flags().StringVarP(&configFile, "config", "c", configFile, "config file")

	// TODO start without subcommand

	cmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		_ = cmd.Help()
		return nil
	})

	cmd.AddCommand(
		versionCmd(),
	)

	return cmd
}

var Version = "devel"
var BuildTime = ""
var ShaVer = "undefined"

func versionCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "version",
		Short: "version ",
		Long:  `version long`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version: %s\n", Version)
			fmt.Printf("Build date: %s\n", BuildTime)
			fmt.Printf("Commit sha: %s\n", ShaVer)
			fmt.Printf("Compiler: %s\n", runtime.Version())
		},
	}

	// hide persistent flag on this command
	cmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		_ = command.Flags().MarkHidden("pers")
		command.Parent().HelpFunc()(command, strings)
	})

	return &cmd
}
