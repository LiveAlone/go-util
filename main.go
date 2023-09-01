package main

import (
	"fmt"
	"github.com/LiveAlone/go-util/appfx"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(appfx.AppConstruct()...),
		fx.WithLogger(
			func(logger *zap.Logger) fxevent.Logger {
				return &fxevent.ZapLogger{Logger: logger}
			},
		),
		fx.Invoke(func(shut fx.Shutdowner, rootCmd *cobra.Command) {
			err := rootCmd.Execute()
			if err != nil {
				fmt.Printf("rootCmd.Execute error %v", zap.Error(err))
			}
			err = shut.Shutdown()
			if err != nil {
				fmt.Printf("shutdown error %v", zap.Error(err))
			}
		}),
	).Run()
}
