package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/mrexmelle/go-psql-size/config"
	"github.com/mrexmelle/go-psql-size/info"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatalf("Usage: %s <app_config> <app_rule> <result_file>\n", os.Args[0])
	}

	cfg, err := config.ReadConfig(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	fmt.Println("Connecting with: ", connectionString)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT table_name, pg_total_relation_size(quote_ident(table_name)) FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE' ORDER BY table_name ASC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	tableInfo := []info.Table{}
	for rows.Next() {
		var it info.Table
		if err := rows.Scan(&it.Name, &it.Size); err != nil {
			continue
		}
		tableInfo = append(tableInfo, it)
	}

	fmt.Printf("Table count: %d\n", len(tableInfo))

	rule, err := config.ReadRule(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Rule count: %d\n", len(rule.Labels))

	labelInfo := []info.Label{}
	for _, r := range rule.Labels {
		groupedTableInfo := []info.Table{}
		for _, tableName := range r.Tables {
			size, err := info.GetSizeOfTableName(tableInfo, tableName)
			if err != nil {
				size = 0
			}
			groupedTableInfo = append(groupedTableInfo, info.Table{Name: tableName, Size: size})
			tableInfo = info.Delete(tableInfo, tableName)
		}
		labelInfo = append(labelInfo, info.Label{Name: r.Name, Tables: groupedTableInfo})
	}

	labelInfo = append(labelInfo, info.Label{Name: "unknown", Tables: tableInfo})

	for _, li := range labelInfo {
		fmt.Printf("Label info: (name: %s, table count: %d, total size: %d)\n", li.Name, len(li.Tables), li.CountSize())
	}

	info.WriteToFile(os.Args[3], labelInfo)
}
