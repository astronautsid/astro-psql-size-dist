package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/mrexmelle/go-psql-size/config"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <app_config> <label_config>\n", os.Args[0])
	}

	cfg, err := config.ReadConfig(os.Args[1])

	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	fmt.Println("Connecting with connection string: ", connectionString)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT table_name, pg_total_relation_size(quote_ident(table_name)) FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE' ORDER BY table_name ASC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	tableInfo := []TableInfo{}
	for rows.Next() {
		var info TableInfo
		if err := rows.Scan(&info.Name, &info.Size); err != nil {
			continue
		}
		tableInfo = append(tableInfo, info)
	}

	for _, info := range tableInfo {
		fmt.Printf("table info: (name: %s, size: %d)\n", info.Name, info.Size)
	}

}
