package connect

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/viper"
	// Driver for cloudsqlpostgres
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
)

// Connect connects to cloudSql postgres instance
func Connect() (*sql.DB, error) {
	host := viper.GetString("google.cloudsql.host")
	name := viper.GetString("google.cloudsql.name")
	user := viper.GetString("google.cloudsql.user")
	pass := viper.GetString("google.cloudsql.pass")
	dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
		host, name, user, pass)
	db, err := sql.Open("cloudsqlpostgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error while connecting to cloudsql: %v", err)
	}
	log.Printf("Connected to cloudsql %q", host)
	return db, nil
}
