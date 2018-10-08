package batchinsert

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/yb172/experiments/pg/connect"
)

// RunBatchInsert inserts given amount of records
func RunBatchInsert(count int) error {
	db, err := connect.Connect()
	if err != nil {
		return fmt.Errorf("error while connecting to db: %v", err)
	}
	initTable(db)

	recordsCount, err := findCount(db)
	if err != nil {
		return fmt.Errorf("error while finding count: %v", err)
	}
	log.Printf("Current record count: %v", recordsCount)

	records := generateRecords(count)
	log.Print("Start inserting records")
	duration, err := bulkInsertWithDriver(records, db)
	if err != nil {
		return fmt.Errorf("error while doing batch insert: %v", err)
	}
	log.Printf("Insertion completed. Took %v", duration)

	recordsCount, err = findCount(db)
	if err != nil {
		return fmt.Errorf("error while finding count: %v", err)
	}
	log.Printf("Record count after update: %v", recordsCount)
	return nil
}

func initTable(db *sql.DB) error {
	var ctx = context.Background()
	createTableQuery := `CREATE TABLE IF NOT EXISTS runs(
		id serial primary key,
		time DATE,
		ref VARCHAR(128)
	)`
	_, err := db.ExecContext(ctx, createTableQuery)
	if err != nil {
		return fmt.Errorf("error while creating table: %v", err)
	}
	return nil
}

type record struct {
	time time.Time
	ref  string
}

func generateRecords(count int) []record {
	var records []record
	for i := 0; i < count; i++ {
		records = append(records, record{time.Now(), time.Now().Format(time.RFC1123Z)})
	}
	return records
}

func findCount(db *sql.DB) (int, error) {
	var count int
	err := db.QueryRowContext(context.Background(), "SELECT COUNT(*) FROM runs").Scan(&count)
	if err != nil {
		return count, fmt.Errorf("error while querying count: %v", err)
	}
	return count, nil
}
