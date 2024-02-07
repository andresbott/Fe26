package cmd

import (
	"github.com/andresbott/Fe26/app/server"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

func Server() *cobra.Command {
	addr := ":8090"
	cmd := cobra.Command{
		Use:   "start",
		Short: "Todo",
		Long:  `Also to do`,
		RunE: func(cmd *cobra.Command, args []string) error {

			cfg := server.Cfg{
				Addr:   addr,
				Logger: nil,
			}
			srv := server.NewServer(cfg)
			osSigExit(func() {
				srv.Stop()
			})
			err := srv.Start()
			if err != nil {
				return err
			}
			return nil

		},
	}
	cmd.Flags().StringVarP(&addr, "addr", "a", addr, "listen address")
	return &cmd
}

func osSigExit(fn func()) chan bool {
	signalCh := make(chan os.Signal, 1)
	stopDone := make(chan bool, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	// handle exit
	go func() {
		<-signalCh
		fn()
		stopDone <- true
	}()
	return stopDone
}
