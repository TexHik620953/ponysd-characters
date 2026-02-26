package appconfig

import "github.com/ilyakaznacheev/cleanenv"

type AppConfig struct {
	HttpAddr      string `env:"HTTP_ADDR" env-default:":8080"`
	JWTSecret     string `env:"JWT_SECRET" env-default:"123"`
	SessionSecret string `env:"SESSION_SECRET" env-default:"123"`

	DatabaseDSN string `env:"DATABASE_DSN" env-default:"postgres://postgres:HermanFuLLer@localhost:5432/general"`

	RedisAddr string `env:"REDIS_ADDR" env-default:"localhost:6379"`
	RedisPwd  string `env:"REDIS_PWD" env-default:"HermanFuLLer"`
	RedisDB   int    `env:"REDIS_DB" env-default:"0"`

	MinioAddr   string `env:"MINIO_ADDR" env-default:"127.0.0.1:9000"`
	MinioKey    string `env:"MINIO_KEY" env-default:"admin"`
	MinioSecret string `env:"MINIO_SECRET" env-default:"HermanFuLLer"`
}

func LoadAppConfig() (*AppConfig, error) {
	var cfg AppConfig
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
