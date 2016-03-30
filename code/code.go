package code

import (
	"time"

	"math/rand"
)

func v1() {
	db := getDb()

	tx, _ := db.Begin()

	result, _ := tx.Exec("INSERT INTO orders (created) VALUES (?)", time.Now())
	orderID, _ := result.LastInsertId()

	_, _ = tx.Exec("INSERT INTO orderItems (order_id, added) VALUES (?, ?)", orderID, time.Now())

	tx.Commit()
}

func v2() {
	db := getDb()

	result, _ := db.Exec("INSERT INTO orders (created) VALUES (?)", time.Now())
	orderID, _ := result.LastInsertId()

	_, _ = db.Exec("INSERT INTO orderItems (order_id, added) VALUES (?, ?)", orderID, time.Now())
}

func v3() {
	db := getDb()

	result, _ := db.Exec("INSERT INTO orders (created) VALUES (?)", time.Now())
	orderID, _ := result.LastInsertId()

	if isError(1) {
		_, _ = db.Exec("DELETE FROM orders WHERE id = ?", orderID)
		return
	}

	_, _ = db.Exec("INSERT INTO orderItems (order_id, added) VALUES (?, ?)", orderID, time.Now())
}

func isError(percentage int) bool {
	return rand.Intn(100) < percentage
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
