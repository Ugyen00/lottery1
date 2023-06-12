package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	postgres_host     = "dpg-chqrtc7dvk4goeu8of9g-a.singapore-postgres.render.com"
	postgres_port     = 5432
	postgres_user     = "postgres_admin"
	postgres_password = "XkTGi9dbW2ro5Tbi6BVIH3zysU1ooCqs"
	postgres_dbname   = "my_db_rzx2"
)

var Db *sql.DB

func init() {
	db_info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", postgres_host, postgres_port, postgres_user, postgres_password, postgres_dbname)

	var err error
	Db, err = sql.Open("postgres", db_info)

	if err != nil {
		panic(err)
	} else {
		log.Println("Database successfully configured")
	}
}
