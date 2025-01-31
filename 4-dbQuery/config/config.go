package config

import (
    "fmt"
    "os"
    "strconv"

    "github.com/joho/godotenv"
)

type DatabaseConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
}

func LoadEnvConfig() (*DatabaseConfig, error) {
    if err := godotenv.Load(); err != nil {
        return nil, fmt.Errorf("Error loading .env file")
    }

    config := &DatabaseConfig{
        Host:     os.Getenv("DB_HOST"),
        User:     os.Getenv("DB_USER"),
        Password: os.Getenv("DB_PASSWORD"),
        DBName:   os.Getenv("DB_NAME"),
    }

    portStr := os.Getenv("DB_PORT")
    if portStr == "" {
        config.Port = 5432
        return config, nil
    }

    port, err := strconv.Atoi(portStr)
    if err != nil {
        return nil, fmt.Errorf("invalid DB_PORT value in .env: %v", err)
    }

    config.Port = port
    return config, nil
}

func GetConnectionString() (string, error) {
    config, err := LoadEnvConfig()
    if err != nil {
        return "", err
    }

    return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        config.Host, config.Port, config.User, config.Password, config.DBName), nil
}

