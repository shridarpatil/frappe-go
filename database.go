package frappe


import (
	"log"
	"fmt"

	"github.com/jmoiron/sqlx"


	// postgres driver
	_ "github.com/lib/pq"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"

)


func InitDB(driver, dsn string) *sqlx.DB {
	// Open the db connection and make a ping.
	fmt.Println("\033[33m Initializing database: \033[0m", driver)
	db, err := sqlx.Connect(driver, dsn)

	if err != nil {
		log.Fatalf("error initializing DB. driver: %s, dsn: %s, message: %v", driver, dsn, err)
		panic("Failed to connec to database")
	}
	fmt.Println("\033[;32m Connected to database: \033[0m", driver,)

	return db
}
