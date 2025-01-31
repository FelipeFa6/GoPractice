package config

import (
    "fmt"
    "os"
    "strconv"

    "github.com/joho/godotenv"
)

// DatabaseConfig holds the configuration for the database connection.
type DatabaseConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
}

// LoadConfig loads the environment variables from the .env file.
func LoadConfig() (*DatabaseConfig, error) {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        return nil, fmt.Errorf("Error loading .env file")
    }

    // Retrieve database configuration from environment variables
    config := &DatabaseConfig{
        Host:     os.Getenv("DB_HOST"),
        User:     os.Getenv("DB_USER"),
        Password: os.Getenv("DB_PASSWORD"),
        DBName:   os.Getenv("DB_NAME"),
    }

    // Attempt to retrieve and parse DB_PORT (if it's set)
    portStr := os.Getenv("DB_PORT")
    if portStr == "" {
        config.Port = 5432 // default port if not set
    } else {
        port, err := strconv.Atoi(portStr) // Convert string to int
        if err != nil {
            return nil, fmt.Errorf("Invalid DB_PORT value in .env: %v", err)
        }
        config.Port = port
    }

    return config, nil
}

// GetDBConnectionString returns the formatted database connection string.
func GetDBConnectionString() (string, error) {
    config, err := LoadConfig()
    if err != nil {
        return "", err
    }

    return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        config.Host, config.Port, config.User, config.Password, config.DBName), nil
}

