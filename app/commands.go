package app

import (
	"context"
	punqy "github.com/punqy/core"
	"github.com/punqy/punqy/app/config"
	"github.com/punqy/punqy/migrations"
	logger "github.com/sirupsen/logrus"
	"github.com/slmder/migrate"
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
		punqy.Command{
			Use:   "migration",
			Short: "migration",
			Args:  cobra.MinimumNArgs(1),
			Children: punqy.Commands{
				punqy.Command{
					Use:   "up",
					Short: "up migrations",
					Run: func(cmd *cobra.Command, args []string) {
						conf := config.NewModule()
						store := punqy.NewModule(
							conf.Config().DatabaseDriverName,
							conf.Config().DatabaseDSN)
						app := migrate.NewManager(conf.Config().MigrationsDir, "", logger.New(), conf.Config().MigrationsTableName, store.DbConnection())
						if err := app.Prepare(migrations.Versions()); err != nil {
							logger.Error(err)
							return
						}
						if err := app.Up(context.Background(), migrate.TransactionModeGeneral); err != nil {
							logger.Error(err)
							return
						}
					},
					Children: punqy.Commands{
						punqy.Command{
							Use:   "version [version-name]",
							Short: "up migration by version name",
							Args:  cobra.MinimumNArgs(1),
							Run: func(cmd *cobra.Command, args []string) {
								conf := config.NewModule()
								store := punqy.NewModule(
									conf.Config().DatabaseDriverName,
									conf.Config().DatabaseDSN)
								app := migrate.NewManager(conf.Config().MigrationsDir, "", logger.New(), conf.Config().MigrationsTableName, store.DbConnection())
								version, err := app.Lookup(migrations.Versions(), args...)
								if err != nil {
									logger.Error(err)
									return
								}
								if err := app.Prepare(version); err != nil {
									logger.Error(err)
									return
								}
								if err := app.Up(context.Background(), migrate.TransactionModeGeneral); err != nil {
									logger.Error(err)
									return
								}
							},
						},
					},
				},
				punqy.Command{
					Use:   "down",
					Short: "down migrations",
					Run: func(cmd *cobra.Command, args []string) {
						conf := config.NewModule()
						store := punqy.NewModule(
							conf.Config().DatabaseDriverName,
							conf.Config().DatabaseDSN)
						app := migrate.NewManager(conf.Config().MigrationsDir, "", logger.New(), conf.Config().MigrationsTableName, store.DbConnection())
						if err := app.Prepare(migrations.Versions()); err != nil {
							logger.Error(err)
							return
						}
						if err := app.Down(context.Background(), migrate.TransactionModeGeneral); err != nil {
							logger.Error(err)
							return
						}
					},
					Children: punqy.Commands{
						punqy.Command{
							Use:   "version [version-name]",
							Short: "down migration by version name",
							Args:  cobra.MinimumNArgs(1),
							Run: func(cmd *cobra.Command, args []string) {
								conf := config.NewModule()
								store := punqy.NewModule(
									conf.Config().DatabaseDriverName,
									conf.Config().DatabaseDSN)
								app := migrate.NewManager(conf.Config().MigrationsDir, "", logger.New(), conf.Config().MigrationsTableName, store.DbConnection())
								version, err := app.Lookup(migrations.Versions(), args...)
								if err != nil {
									logger.Error(err)
									return
								}
								if err := app.Prepare(version); err != nil {
									logger.Error(err)
									return
								}
								if err := app.Down(context.Background(), migrate.TransactionModeGeneral); err != nil {
									logger.Error(err)
									return
								}
							},
						},
					},
				},
				punqy.Command{
					Use:   "generate",
					Short: "generate migration",
					Run: func(cmd *cobra.Command, args []string) {
						conf := config.NewModule()
						store := punqy.NewModule(
							conf.Config().DatabaseDriverName,
							conf.Config().DatabaseDSN)
						app := migrate.NewManager(conf.Config().MigrationsDir, "", logger.New(), conf.Config().MigrationsTableName, store.DbConnection())
						if err := app.Generate(); err != nil {
							logger.Error(err)
							return
						}
					},
				},
			},
		},
	}
}
