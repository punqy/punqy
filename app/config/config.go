package config

type AppEnv string

const (
	EnvDev  = AppEnv("dev")
	EnvTest = AppEnv("test")
	EnvProd = AppEnv("prod")
)

type Config struct {
	AppEnv                    AppEnv
	HttpServerPort            int    `env:"HTTP_SERVER_PORT,default=6009"`
	HttpFirewallEnabled       bool   `env:"HTTP_FIREWALL_ENABLED,default=true"`
	DatabaseDriverName        string `env:"DATABASE_DRIVER_NAME,default=postgres"`
	DatabaseDSN               string `env:"DATABASE_DSN"`
	OauthAccessTokenTTL       int    `env:"OAUTH_ACCESS_TOKEN_TTL,default=15"`
	OauthRefreshTokenTTL      int    `env:"OAUTH_REFRESH_TOKEN_TTL,default=2880"`
	MigrationsDir             string `env:"MIGRATIONS_DIR,default=migrations"`
	MigrationsTableName       string `env:"MIGRATIONS_TABLE_NAME,default=migration_versions"`
	ProfilerDir               string `env:"PROFILER_DIR,default=var/profiler"`
	ProfilerMiddlewareEnabled bool   `env:"PROFILER_MIDDLEWARE_ENABLED,default=true"`
	TemplateDir               string `env:"TEMPLATE_DIR,default=templates"`
}
