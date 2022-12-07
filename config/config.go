package config

import (
	"fmt"
	"os"
	"strconv"
)

func GetConfig() Config {
	if checkIfProdFromEnv() {
		port, err := strconv.Atoi(os.Getenv("port"))
		if err != nil {
			panic(err)
		}
		fmt.Println("Running with prod env")
		return Config{
			Env:            "prod",
			Pepper:         os.Getenv("pepper"),
			HMACKey:        os.Getenv("hmac_key"),
			Port:           port,
			connectionInfo: getConnectionInfoFromEnv(),
		}
	}
	fmt.Println("Running with dev env")
	return Config{
		Env:            "dev",
		Pepper:         "secret-secret-secret",
		HMACKey:        "secret-hmac-hmacSecretKey",
		Port:           3000,
		connectionInfo: getConnectionInfoDev(),
	}
}

type Config struct {
	Env            string
	Pepper         string
	HMACKey        string
	Port           int
	connectionInfo string
}

func checkIfProdFromEnv() bool {
	return os.Getenv("env") == "prod"
}

func (c *Config) IsProd() bool {
	return c.Env == "prod"
}

func (c *Config) Dialect() string {
	return "postgres"
}

func (c *Config) ConnectionInfo() string {
	return c.connectionInfo
}

func getConnectionInfoFromEnv() string {
	connectionInfo := os.Getenv("DATABASE_URL")
	if connectionInfo == "" {
		panic("Empty env var DATABASE_URL")
	}
	return fmt.Sprintf("%s?sslmode=disable", connectionInfo)
}

type postgresConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func getConnectionInfoDev() string {

	// todo: read conf from json
	p := postgresConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "aru",
		Password: "",
		Name:     "lenslocked_dev",
	}

	if p.Password == "" {
		return fmt.Sprintf("host=%s port=%d user=%s "+
			"dbname=%s sslmode=disable",
			p.Host, p.Port, p.User, p.Name)
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s "+
		"dbname=%s sslmode=disable",
		p.Host, p.Port, p.User, p.Name, p.Password)
}
