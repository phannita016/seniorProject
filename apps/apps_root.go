package apps

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	sn "github.com/phannita016/seniorProject"
	"github.com/phannita016/seniorProject/config"
	"github.com/phannita016/seniorProject/x"
	"github.com/spf13/cobra"
)

func fileConfig() string {
	defer x.Recover()
	return "config.yaml"
}

func NewAppsRoot() *cobra.Command {
	var file string

	rootcmd := cobra.Command{
		Use:     "service",
		Short:   "Application by Computer Engineering Nareasuan university",
		Long:    "The 'seniorProject' application version - " + sn.AppVersion,
		Version: sn.AppVersion + " - " + sn.TagApplication,

		SilenceErrors: true,

		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := config.New(file)
			if err != nil {
				return err
			}

			app, err := NewApps(conf)
			if err != nil {
				return err
			}

			var ctx = context.Background()

			if err = app.Runner(ctx); err != nil {
				return err
			}

			done := make(chan os.Signal, 1)
			signal.Notify(done, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM)

			<-done

			ctx, cancel := context.WithTimeout(ctx, time.Second*5)
			defer cancel()

			return app.Close(ctx)
		},
	}

	flags := rootcmd.PersistentFlags()

	flags.StringVar(
		&file,
		"config",
		fileConfig(),
		"Config file locatioin path",
	)

	return &rootcmd
}
