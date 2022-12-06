package main

import "fmt"

type PostgresConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (c PostgresConfig) Dialect() string {
	return "postgres"
}

// previously was
//  	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
//		"dbname=%s sslmode=disable password=%s",
//		host, port, userName, dbname, password)

func (c PostgresConfig) ConnectionInfo() string {
	if c.Password == "" {
		return fmt.Sprintf("host=%s port=%d user=%s "+
			"dbname=%s sslmode=disable",
			c.Host, c.Port, c.User, c.Name)
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s "+
		"dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Name, c.Password)
}

func DefaultPostgresConfigDev() PostgresConfig {
	return PostgresConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "aru",
		Password: "",
		Name:     "lenslocked_dev",
	}
}

func DefaultPostgresConfigTest() PostgresConfig {
	return PostgresConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "aru",
		Password: "",
		Name:     "lenslocked_test",
	}
}

type Config struct {
	Port    int    `json:"port"`
	Env     string `json:"env"`
	Pepper  string `json:"pepper"`
	HMACKey string `json:"hmac_key"`
}

func (c Config) IsProd() bool {
	return c.Env == "prod"
}

func DefaultConfig() Config {
	return Config{
		Port:    3000,
		Env:     "dev",
		Pepper:  "secret-secret-secret",
		HMACKey: "secret-hmac-hmacSecretKey",
	}
}
