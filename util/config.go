package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	DBProduct           string        `mapstructure:"DB_FOODIO"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	SecretKey           string        `mapstructure:"SECRET_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefeshTokenDuration time.Duration `mapstructure:"REFESH_TOKEN_DURATION"`
	MaxConns            int32         `mapstructure:"MAX_CONNS"`
	MinConns            int32         `mapstructure:"MIN_CONNS"`
	MaxConnIdleTime     time.Duration `mapstructure:"MAX_CONN_IDLE_TIME"`
	MaxConnLifetime     time.Duration `mapstructure:"MAX_CONN_LIFE_TIME"`
	FrontEndUrl         string        `mapstructure:"FRONTEND_URL"`
	EmailAdmin          string        `mapstructure:"EMAIL_ADMIN"`
	SMTPHost            string        `mapstructure:"SMTP_HOST"`
	SMTPPort            string        `mapstructure:"SMTP_PORT"`
	SMTPUsername        string        `mapstructure:"SMTP_USERNAME"`
	SMTPPassword        string        `mapstructure:"SMTP_PASSWORD"`
	APKLink             string        `mapstructure:"APK_LINK"`
	ResetPass           time.Duration `mapstructure:"RESET_PASS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
