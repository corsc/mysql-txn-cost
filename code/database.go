package code

import (
	"database/sql"

	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var dbOnce sync.Once

func getDb() *sql.DB {
	dbOnce.Do(func() {
		var err error
		db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/txncost?parseTime=true")
		if err != nil {
			panic("Failed to open database connection. Error:" + err.Error())
		}
		db.SetMaxIdleConns(8)
		db.SetMaxOpenConns(8)
		if db == nil {
			panic("db is nil")
		}
	})
	return db
}

func dropAndCreateTables() {
	db := getDb()

	_, _ = db.Exec(`DROP TABLE IF EXISTS orders`)
	_, _ = db.Exec(`CREATE TABLE orders (
  		id INT NOT NULL AUTO_INCREMENT,
  		created DATETIME NULL,
  		PRIMARY KEY (id)
  	)`)
	_, _ = db.Exec(`DROP TABLE IF EXISTS orderitems`)
	_, _ = db.Exec(`CREATE TABLE orderitems (
  		id INT NOT NULL AUTO_INCREMENT,
  		order_id INT NOT NULL,
  		added DATETIME NULL,
  		PRIMARY KEY (id)
  	)`)
}
