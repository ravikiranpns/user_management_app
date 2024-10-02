package config

import "os"

// GetDatabaseDSN returns the Database connection string
func GetDatabaseDSN() string {
    return os.Getenv("DATABASE_DSN") // Set this in the environment or use a default string
}
