package frappe


import (


	"github.com/jmoiron/sqlx"

	// postgres driver
	_ "github.com/lib/pq"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"

)


func (f *Frappe) initDB() {
	// Open the db connection and make a ping.
	f.Log.InfoF("Initializing database: %s", f.Config.Driver)
	db, err := sqlx.Connect(f.Config.Driver, f.Config.DSN)
	if err != nil {
		f.Log.CriticalF("error initializing DB. driver: %s, dsn: %s, message: %v", f.Config.Driver, f.Config.DSN, err)
		panic("Failed to initialize DB")
	}
	f.Log.NoticeF("Connected to database: %s", f.Config.Driver)

	f.Db = db
}
