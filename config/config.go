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
		Objects: []string{"branch", "staff", "staffTariff", "sale", "staffTransaction"},
	}
}

type ConfigPostgres struct {
	Environment string // debug, test, release

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
	RedisHost        string
	RedisPort        string
	RedisPassword    string
	RedisDatabase    string

	Port string

	PostgresMaxConnections int32
}

// Branch implements storage.StorageI.
func (ConfigPostgres) Branch() storage.BranchesI {
	panic("unimplemented")
}

// Sale implements storage.StorageI.
func (ConfigPostgres) Sale() storage.SalesI {
	panic("unimplemented")
}

// Staff implements storage.StorageI.
func (ConfigPostgres) Staff() storage.StaffsI {
	panic("unimplemented")
}

// StaffTarif implements storage.StorageI.
func (ConfigPostgres) StaffTarif() storage.StaffTariffsI {
	panic("unimplemented")
}

// StaffTransaction implements storage.StorageI.
func (ConfigPostgres) StaffTransaction() storage.StaffTransactionI {
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
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "project_db"))

	config.PostgresMaxConnections = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 30))

	config.RedisHost = cast.ToString(getOrReturnDefaultValue("REDIS_HOST", "localhost"))
	config.RedisPort = cast.ToString(getOrReturnDefaultValue("REDIS_PORT", 6379))
	config.RedisPassword = cast.ToString(getOrReturnDefaultValue("REDIS_PASSWORD", ""))
	config.RedisDatabase = cast.ToString(getOrReturnDefaultValue("REDIS_DATABASE", 0))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return defaultValue
}
