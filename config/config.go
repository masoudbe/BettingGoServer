package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Username: "postgres",
			Password: "123456",
			Name:     "Betting",
			Charset:  "utf8",
		},
	}
}
