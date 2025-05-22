package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jacksongeorge770/InvestmentCalculator/cmd/api"
	"github.com/jacksongeorge770/InvestmentCalculator/config"
	"github.com/jacksongeorge770/InvestmentCalculator/db"
)

func main() {

	// Configure MySQL connection
	cfg := mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	// Connect to RDS
	db, err := db.MysqlStorage(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close() // Ensure connection is closed on exit

	// Initialize storage (ping and migrations)
	initStorage(db)

	// Start API server
	server := api.NewAPIServer(fmt.Sprintf(":%s", config.Envs.Port), db)
	if err := server.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
