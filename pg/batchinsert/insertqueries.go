package batchinsert

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/lib/pq"
)

// Bulk insert as described here: https://stackoverflow.com/a/48070387/518469
func bulkInsertWithQuery(records []record, db *sql.DB) (time.Duration, error) {
	valueStrings := make([]string, 0, len(records))
	valueArgs := make([]interface{}, 0, len(records)*2)
	for i, record := range records {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2))
		valueArgs = append(valueArgs, record.time)
		valueArgs = append(valueArgs, record.ref)
	}
	stmt := fmt.Sprintf("INSERT INTO runs (time, ref) VALUES %s", strings.Join(valueStrings, ","))

	start := time.Now()
	_, err := db.Exec(stmt, valueArgs...)
	end := time.Now()
	return end.Sub(start), err
}

// Bulk insert as described here: https://godoc.org/github.com/lib/pq#hdr-Bulk_imports
func bulkInsertWithDriver(records []record, db *sql.DB) (time.Duration, error) {
	var duration time.Duration
	start := time.Now()
	txn, err := db.Begin()
	if err != nil {
		return duration, fmt.Errorf("error while starting transactions")
	}

	stmt, err := txn.Prepare(pq.CopyIn("runs", "time", "ref"))
	if err != nil {
		return duration, fmt.Errorf("error while preparing statement: %v", err)
	}

	for _, record := range records {
		_, err = stmt.Exec(record.time, record.ref)
		if err != nil {
			return duration, fmt.Errorf("error while performing insertions: %v", err)
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		return duration, fmt.Errorf("error while executing final statement: %v", err)
	}

	err = stmt.Close()
	if err != nil {
		return duration, fmt.Errorf("error while closing statement: %v", err)
	}

	err = txn.Commit()
	if err != nil {
		return duration, fmt.Errorf("error while committing transactions: %v", err)
	}
	end := time.Now()
	return end.Sub(start), nil
}
