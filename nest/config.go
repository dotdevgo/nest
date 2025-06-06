package nest

import (
	"os"
	"time"

	"dotdev/logger"

	"github.com/defval/di"
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

const (
	// TemplateDir stores the name of the directory that contains templates
	TemplateDir = "templates"

	// TemplateExt stores the extension used for the template files
	TemplateExt = ".gohtml"

	// StaticDir stores the name of the directory that will serve static files
	StaticDir = "static"
)

type Environment string

const (
	EnvLocal      Environment = "local"
	EnvTest       Environment = "test"
	EnvDevelop    Environment = "dev"
	EnvStaging    Environment = "staging"
	EnvQA         Environment = "qa"
	EnvProduction Environment = "prod"
)

// SwitchEnvironment sets the environment variable used to dictate which environment the application is
// currently running in.
// This must be called prior to loading the configuration in order for it to take effect.
func SwitchEnvironment(env Environment) {
	if err := os.Setenv("APP_ENV", string(env)); err != nil {
		logger.Fatal(err)
	}
}

// DBConnectFunc defines a database connection initialization function.
type ContainerFactoryFn func(providers ...di.Option) (*di.Container, error)

type (
	// Config stores complete configuration
	Config struct {
		App      AppConfig
		CORS     CORSConfig
		Cache    CacheConfig
		HTTP     HTTPConfig
		Database DatabaseConfig

		ContainerFactory ContainerFactoryFn
	}

	// AppConfig stores application configuration
	AppConfig struct {
		//Name        string      `env:"APP_NAME,default=Nest"`
		Timeout time.Duration `env:"APP_TIMEOUT,default=20s"`

		Environment Environment `env:"APP_ENV,default=local"`
		// THIS MUST BE OVERRIDDEN ON ANY LIVE ENVIRONMENTS
		EncryptionKey string `env:"APP_ENCRYPTION_KEY,default=?E(G+KbPeShVmYq3t6w9z$C&F)J@McQf"`
		// PasswordToken struct {
		// 	Expiration time.Duration `env:"APP_PASSWORD_TOKEN_EXPIRATION,default=60m"`
		// 	Length     int           `env:"APP_PASSWORD_TOKEN_LENGTH,default=64"`
		// }
		// EmailVerificationTokenExpiration time.Duration `env:"APP_EMAIL_VERIFICATION_TOKEN_EXPIRATION,default=12h"`
	}

	// CORSConfig stores Cors configuration
	CORSConfig struct {
		Origin string `env:"CORS_ORIGIN"`
	}

	// HTTPConfig stores HTTP configuration
	HTTPConfig struct {
		Hostname     string        `env:"HTTP_HOSTNAME"`
		Port         uint16        `env:"HTTP_PORT,default=1333"`
		ReadTimeout  time.Duration `env:"HTTP_READ_TIMEOUT,default=5s"`
		WriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT,default=10s"`
		IdleTimeout  time.Duration `env:"HTTP_IDLE_TIMEOUT,default=2m"`
		TLS          struct {
			Enabled     bool   `env:"HTTP_TLS_ENABLED,default=false"`
			Certificate string `env:"HTTP_TLS_CERTIFICATE"`
			Key         string `env:"HTTP_TLS_KEY"`
		}
	}

	// CacheConfig stores the cache configuration
	CacheConfig struct {
		Hostname     string `env:"CACHE_HOSTNAME,default=localhost"`
		Port         uint16 `env:"CACHE_PORT,default=6379"`
		Password     string `env:"CACHE_PASSWORD"`
		Database     int    `env:"CACHE_DB,default=0"`
		TestDatabase int    `env:"CACHE_DB_TEST,default=1"`
		// Expiration   struct {
		// 	StaticFile time.Duration `env:"CACHE_EXPIRATION_STATIC_FILE,default=4380h"`
		// 	Page       time.Duration `env:"CACHE_EXPIRATION_PAGE,default=24h"`
		// }
	}

	// DatabaseConfig stores the database configuration
	DatabaseConfig struct {
		Hostname     string `env:"DB_HOSTNAME,default=localhost"`
		Port         uint16 `env:"DB_PORT,default=5432"`
		User         string `env:"DB_USER,default=admin"`
		Password     string `env:"DB_PASSWORD,default=admin"`
		Database     string `env:"DB_NAME,default=app"`
		TestDatabase string `env:"DB_NAME_TEST,default=app_test"`
	}
)

var cfg Config

// GetConfig loads and returns configuration
func GetConfig() Config {
	return cfg
}

var isEnvLoaded = false

// loadEnvironment godoc
func loadEnvironment() error {
	if isEnvLoaded {
		return nil
	}

	isEnvLoaded = true

	dir, err := os.Getwd()
	logger.FatalOnError(err)

	if err := godotenv.Load(dir + "/.env"); err != nil {
		// logger.Warn(err)
		return err
	}

	if err := envdecode.StrictDecode(&cfg); err != nil {
		return err
	}

	// logger.FatalOnError(envdecode.StrictDecode(&cfg))

	// logger.Info("==> Config loaded")

	return nil
}

// NewConfig godoc
// func NewConfig[T any]() T {
// 	var data T

// 	if err := envdecode.StrictDecode(&data); err != nil {
// 		logger.Error(err)
// 	}

// 	return data
// }
