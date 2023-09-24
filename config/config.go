package config

import (
	"fmt"
	"new_project/storage"
	"os"

	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Limit   int
	Page    int
	Methods []string
	Objects []string
}

const (
	SuccessStatus = iota + 1
	CancelStatus
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"

	TimeExpiredAt = time.Hour * 720
)

func Load() *Config {
	return &Config{
		Limit:   10,
		Page:    1,
		Methods: []string{"create", "update", "get", "getAll", "delete"},
		Objects: []string{"branch", "category", "product", "comingTable", "comingTableProduct"},
	}
}

type ConfigPostgres struct {
	Environment string // debug, test, release

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	Port string

	PostgresMaxConnections int32
}

// Branch implements storage.StorageI.
func (ConfigPostgres) Branch() storage.BranchesI {
	panic("unimplemented")
}

// Category implements storage.StorageI.
func (ConfigPostgres) Category() storage.CategoriesI {
	panic("unimplemented")
}

// Product implements storage.StorageI.
func (ConfigPostgres) Product() storage.ProductsI {
	panic("unimplemented")
}

// ComingTable implements storage.StorageI.
func (ConfigPostgres) ComingTable() storage.ComingTablesI {
	panic("unimplemented")
}

// ComingTableProduct implements storage.StorageI.
func (ConfigPostgres) ComingTableProduct() storage.ComingTableProductsI {
	panic("unimplemented")
}

// Remaining implements storage.StorageI.
func (ConfigPostgres) Remaining() storage.RemainingsI {
	panic("unimplemented")
}

// Load ...
func LoadP() ConfigPostgres {
	if err := godotenv.Load("./.env"); err != nil {
		fmt.Println("No .env file found")
	}

	config := ConfigPostgres{}

	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", DebugMode))
	config.Port = cast.ToString(getOrReturnDefaultValue("PORT", "8080"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "postgres"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "exams"))

	config.PostgresMaxConnections = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 30))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return defaultValue
}
