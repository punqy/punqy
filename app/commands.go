package app

import (
	"context"
	punqy "github.com/punqy/core"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func Commands() punqy.Commands {
	return punqy.Commands{
		punqy.Command{
			Use:   "http-serve",
			Short: "Run http server",
			Long:  "",
			Args:  nil,
			Run: func(cmd *cobra.Command, args []string) {
				ctx := context.Background()
				registry, err := BuildRegistry(ctx)
				if err != nil {
					logger.Error(err)
					return
				}
				registry.HttpServer().Serve(ctx)
			},
		},
	}
}
