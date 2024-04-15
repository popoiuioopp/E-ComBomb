package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

type Env struct {
	DBHost string `mapstructure:"DB_HOST"`
	DBPort string `mapstructure:"DB_PORT"`
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`
	DBName string `mapstructure:"DB_NAME"`
}

func NewEnv() *Env {
	viper.AutomaticEnv() // Read from environment variables first

	viper.SetConfigFile(".env")
	err := viper.MergeInConfig() // Merge in the values from .env file, if present
	if err != nil {
		log.Printf("Warning: No .env file found or error in reading it: %v\n", err)
	}

	env := Env{}
	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &env
}

func NewMySQLDatabase(env *Env) *sql.DB {
	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass
	dbName := env.DBName

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func UpdateChangelog(db *sql.DB, filename string) error {
	query := "INSERT INTO database_changelog (filename, executed_by) VALUES (?, ?)"

	_, err := db.Exec(query, filename, "migration_runner")

	return err
}

func GetChangelog(db *sql.DB) ([]string, error) {
	var filenames []string
	query := "SELECT filename FROM database_changelog"

	results, err := db.Query(query)
	if err != nil {
		return []string{}, err
	}

	for results.Next() {
		var filename string

		err := results.Scan(&filename)
		if err != nil {
			return []string{}, err
		}

		filenames = append(filenames, filename)
	}

	return filenames, nil
}

// Execute a SQL file
func executeSQLFile(db *sql.DB, filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	_, err = db.Exec(string(content))
	if err != nil {
		return err
	}
	return nil
}

// Main migration function
func migrate(db *sql.DB, migrationsDir string) {
	appliedMigrationsMap := make(map[string]bool)

	appliedMigrations, err := GetChangelog(db)
	if err != nil {
		log.Fatalf("Error fetching changelog: %v", err)
	}

	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		log.Fatal(err)
	}

	// Sort files by name to ensure correct order
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for _, filename := range appliedMigrations {
		appliedMigrationsMap[filename] = true
	}

	for _, file := range files {
		if appliedMigrationsMap[file.Name()] {
			log.Printf("Skipping already applied migration: %s\n", file.Name())
			continue
		}

		log.Printf("Applying migration: %s\n", file.Name())
		err := executeSQLFile(db, migrationsDir+"/"+file.Name())
		if err != nil {
			log.Fatalf("Error applying migration %s: %v", file.Name(), err)
		}

		err = UpdateChangelog(db, file.Name())
		if err != nil {
			log.Fatalf("Error updating changelog for %s: %v", file.Name(), err)
		}
	}
}

// Helper function to read the manifest file
func readManifest(manifestFile string) ([]string, error) {
	var filenames []string
	data, err := os.ReadFile(manifestFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &filenames)
	if err != nil {
		return nil, err
	}
	return filenames, nil
}

func updateStoreProcs(db *sql.DB, storeProcDir, manifestFile string) {
	filenames, err := readManifest(manifestFile)
	if err != nil {
		log.Fatalf("Error reading manifest file: %v", err)
	}

	for _, filename := range filenames {
		fullPath := storeProcDir + "/" + filename
		log.Printf("Applying Stored Procedure: %s\n", filename)
		err := executeSQLFile(db, fullPath)
		if err != nil {
			log.Fatalf("Error applying Stored Procedure %s: %v", filename, err)
		}
	}
}

func main() {
	env := NewEnv()
	db := NewMySQLDatabase(env)
	defer db.Close()

	migrate(db, "migrations")
	updateStoreProcs(db, "store_procs", "stored_procs_order.json")
}
