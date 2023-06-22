package nest

import (
	"os"
	"time"

	"dotdev/nest/pkg/logger"
	"dotdev/nest/pkg/utils"

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

	// StaticPrefix stores the URL prefix used when serving static files
	StaticPrefix = "files"
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

type (
	// Config stores complete configuration
	Config struct {
		HTTP     HTTPConfig
		CORS     CORSConfig
		App      AppConfig
		Cache    CacheConfig
		Database DatabaseConfig
		Mail     MailConfig
	}

	// CORSConfig stores Cors configuration
	CORSConfig struct {
		Origin string `env:"CORS_ORIGIN"`
	}

	// HTTPConfig stores HTTP configuration
	HTTPConfig struct {
		Origin       string        `env:"CORS_ORIGIN"`
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

	// AppConfig stores application configuration
	AppConfig struct {
		Name        string      `env:"APP_NAME,default=Nest"`
		Environment Environment `env:"APP_ENV,default=local"`
		// THIS MUST BE OVERRIDDEN ON ANY LIVE ENVIRONMENTS
		EncryptionKey string        `env:"APP_ENCRYPTION_KEY,default=?E(G+KbPeShVmYq3t6w9z$C&F)J@McQf"`
		Timeout       time.Duration `env:"APP_TIMEOUT,default=20s"`
		PasswordToken struct {
			Expiration time.Duration `env:"APP_PASSWORD_TOKEN_EXPIRATION,default=60m"`
			Length     int           `env:"APP_PASSWORD_TOKEN_LENGTH,default=64"`
		}
		EmailVerificationTokenExpiration time.Duration `env:"APP_EMAIL_VERIFICATION_TOKEN_EXPIRATION,default=12h"`
	}

	// CacheConfig stores the cache configuration
	CacheConfig struct {
		Hostname     string `env:"CACHE_HOSTNAME,default=localhost"`
		Port         uint16 `env:"CACHE_PORT,default=6379"`
		Password     string `env:"CACHE_PASSWORD"`
		Database     int    `env:"CACHE_DB,default=0"`
		TestDatabase int    `env:"CACHE_DB_TEST,default=1"`
		Expiration   struct {
			StaticFile time.Duration `env:"CACHE_EXPIRATION_STATIC_FILE,default=4380h"`
			Page       time.Duration `env:"CACHE_EXPIRATION_PAGE,default=24h"`
		}
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

	// MailConfig stores the mail configuration
	MailConfig struct {
		Hostname    string `env:"MAIL_HOSTNAME"`
		Port        uint16 `env:"MAIL_PORT,default=25"`
		User        string `env:"MAIL_USER,default=admin"`
		Password    string `env:"MAIL_PASSWORD,default=admin"`
		FromAddress string `env:"MAIL_FROM_ADDRESS,default=admin@localhost"`
	}
)

// GetConfig loads and returns configuration
func GetConfig() (Config, error) {
	LoadEnv()

	var cfg Config
	err := envdecode.StrictDecode(&cfg)

	return cfg, err
}

var isEnvLoaded = false

// LoadEnv godoc
func LoadEnv() {
	if isEnvLoaded {
		return
	}

	isEnvLoaded = true

	dir, err := os.Getwd()
	utils.NoErrorOrFatal(err)

	if err := godotenv.Load(dir + "/.env"); err != nil {
		logger.Error("Error loading .env")
		logger.Error(err)
		return
	}

	logger.Log("==> Loaded .env file")
}

// NewConfig godoc
func NewConfig[T any]() T {
	var data T

	if err := envdecode.StrictDecode(&data); err != nil {
		logger.Error(err)
	}

	return data
}
