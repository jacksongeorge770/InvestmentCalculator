package main

import (
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	config "github.com/jacksongeorge770/InvestmentCalculator/config"
	"github.com/jacksongeorge770/InvestmentCalculator/db"
	"github.com/pressly/goose"
)

func main() {
	// Check if the required environment variables are set
	if config.Envs.DBUser == "" || config.Envs.DBPassword == "" || config.Envs.DBAddress == "" || config.Envs.DBName == "" {
		log.Fatal("Missing database credentials in environment variables")
	}

	// Connect to MySQL database
	cfg := mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress + ":" + config.Envs.DBPort,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := db.MysqlStorage(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close() // Ensure the database connection is closed when done

	// Check the command line argument
	if len(os.Args) < 2 {
		log.Fatal("Missing migration command (up or down)")
	}

	cmd := os.Args[1]

	// Set the migrations directory
	migrationsDir := "cmd/migrate/migration"

	// Configure goose to use MySQL
	goose.SetDialect("mysql")

	// Execute the migration command
	switch cmd {
	case "up":
		err = goose.Up(db, migrationsDir)
		if err != nil {
			log.Fatalf("Failed to apply migrations: %v", err)
		}
		log.Println("Migrations applied successfully!")
	case "down":
		err = goose.Down(db, migrationsDir)
		if err != nil {
			log.Fatalf("Failed to revert migrations: %v", err)
		}
		log.Println("Migrations reverted successfully!")
	default:
		log.Fatal("Invalid command. Use 'up' or 'down'.")
	}
}
