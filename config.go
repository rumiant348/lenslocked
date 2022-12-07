package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type PostgresConfig interface {
	Dialect() string
	ConnectionInfo() string
}

type postgresConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (c postgresConfig) Dialect() string {
	return "postgres"
}

func (c postgresConfig) ConnectionInfo() string {
	if c.Password == "" {
		return fmt.Sprintf("host=%s port=%d user=%s "+
			"dbname=%s sslmode=disable",
			c.Host, c.Port, c.User, c.Name)
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s "+
		"dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Name, c.Password)
}

type envPostgresConfig struct {
}

func (c envPostgresConfig) Dialect() string {
	return "postgres"
}

func (c envPostgresConfig) ConnectionInfo() string {
	connectionInfo := os.Getenv("DATABASE_URL")
	if connectionInfo == "" {
		panic("Empty env var DATABASE_URL")
	}
	return connectionInfo
}

func DefaultPostgresConfigDev() PostgresConfig {
	return postgresConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "aru",
		Password: "",
		Name:     "lenslocked_dev",
	}
}

func DefaultPostgresConfigTest() PostgresConfig {
	return postgresConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "aru",
		Password: "",
		Name:     "lenslocked_test",
	}
}

type Config struct {
	Port     int            `json:"port"`
	Env      string         `json:"env"`
	Pepper   string         `json:"pepper"`
	HMACKey  string         `json:"hmac_key"`
	Database PostgresConfig `json:"database"`
}

func (c Config) IsProd() bool {
	return c.Env == "prod"
}

func DefaultConfig() Config {
	return Config{
		Port:     3000,
		Env:      "dev",
		Pepper:   "secret-secret-secret",
		HMACKey:  "secret-hmac-hmacSecretKey",
		Database: DefaultPostgresConfigDev(),
	}
}

func LoadConfig(isProd bool) Config {

	if isProd {
		return loadConfigProd()
	}

	f, err := os.Open(".config.json")
	if err != nil {
		fmt.Println("Using the default config")
		return DefaultConfig()
	}
	var c Config
	dec := json.NewDecoder(f)
	err = dec.Decode(&c)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully loaded .config.json")
	return c
}

func loadConfigProd() Config {
	port, err := strconv.Atoi(os.Getenv("port"))
	if err != nil {
		panic(err)
	}
	return Config{
		Port:     port,
		Env:      os.Getenv("env"),
		Pepper:   os.Getenv("pepper"),
		HMACKey:  os.Getenv("hmac_key"),
		Database: envPostgresConfig{},
	}
}
