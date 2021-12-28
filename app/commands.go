package app

import (
	"context"
	punqy "github.com/punqy/core"
	"github.com/punqy/punqy/app/config"
	"github.com/punqy/punqy/migrations"
	"github.com/punqy/punqy/model/storage"
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
			Use:   "create-admin",
			Short: "Create def user",
			Long:  "",
			Args:  nil,
			Run: func(cmd *cobra.Command, args []string) {
				ctx := context.Background()
				registry, err := BuildRegistry(ctx)
				if err != nil {
					logger.Error(err)
					return
				}
				password, err := registry.PasswordEncoder.EncodePassword("Qwerty123", nil)
				if err != nil {
					logger.Error(err)
					return
				}
				u := storage.User{Username: "admin@admin.com", Password: password, Roles: storage.StringList{storage.RoleSuperAdmin}}
				if err := u.Init(); err != nil {
					logger.Error(err)
					return
				}
				if err := registry.UserRepository().Insert(ctx, u); err != nil {
					logger.Error(err)
					return
				}
			},
		},
		punqy.Command{
			Use:   "create-oauth-client",
			Short: "Create oauth client",
			Long:  "",
			Args:  nil,
			Run: func(cmd *cobra.Command, args []string) {
				ctx := context.Background()
				registry, err := BuildRegistry(ctx)
				if err != nil {
					logger.Error(err)
					return
				}
				client, err := registry.ClientRepository().NewOauthClient(ctx)
				if err != nil {
					logger.Error(err)
					return
				}
				logger.Infof("OAuth client created")
				logger.Infof("Id: %s", client.ID)
				logger.Infof("Secret: %s", client.ClientSecret)
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
