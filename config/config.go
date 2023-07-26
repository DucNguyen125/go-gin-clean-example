package config

import (
	env "github.com/Netflix/go-env"
)

type Environment struct {
	Constants
	CorsAllowOrigins   string `env:"CORS_ALLOW_ORIGINS,required=true"`
	PostgresqlHost     string `env:"POSTGRESQL_HOST,required=true"`
	PostgresqlPort     string `env:"POSTGRESQL_PORT,required=true"`
	PostgresqlUserName string `env:"POSTGRESQL_USERNAME,required=true"`
	PostgresqlPassword string `env:"POSTGRESQL_PASSWORD,required=true"`
	PostgresqlDatabase string `env:"POSTGRESQL_DATABASE,required=true"`
	// S3URI              string `env:"S3_URI, required=true"`
	// JwtSecretKey       string `env:"JWT_SECRET_KEY, required=true"`
	// FrontendUri        string `env:"FRONTEND_URI, required=true"`
	// JwtExpirationHour  int    `env:"JWT_EXPIRATION_HOUR, required=true"`
	// SendgridApiKey     string `env:"SENDGRID_API_KEY, required=true"`
	// MailFrom           string `env:"MAIL_FROM"`
}

func Load() (*Environment, error) {
	var environment Environment
	_, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		return nil, err
	}
	return &environment, nil
}
