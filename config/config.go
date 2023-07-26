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
	RedisURI           string `env:"REDIS_URI,required=true"`
	RedisPassword      string `env:"REDIS_PASSWORD,required=true"`
	ElasticSearchURI   string `env:"ELASTICSEARCH_URI,required=true"`
	AwsAccessKeyID     string `env:"AWS_ACCESS_KEY_ID,required=true"`
	AwsSecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY,required=true"`
	AwsRegion          string `env:"AWS_REGION,required=true"`
	AwsS3Bucket        string `env:"AWS_S3_BUCKET,required=true"`
	MongoURI           string `env:"MONGO_URI,required=true"`
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
