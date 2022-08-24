package frappe


import (


	"github.com/jmoiron/sqlx"


	// postgres driver
	_ "github.com/lib/pq"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"

)


func InitDB(driver, dsn string) *sqlx.DB {
	// Open the db connection and make a ping.
	Frappe.Log.InfoF("Initializing database: %s", driver)
	db, err := sqlx.Connect(driver, dsn)

	if err != nil {
		Frappe.Log.CriticalF("error initializing DB. driver: %s, dsn: %s, message: %v", driver, dsn, err)
		panic("Failed to initialize DB")
	}
	Frappe.Log.NoticeF("Connected to database: %s", driver)

	return db
}
