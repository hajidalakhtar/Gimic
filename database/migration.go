
package database

import (
	"Gimic/model"
)

// Migrate : migrate the table
func Migrate() {

	db := Connect()
	defer db.Close()

	db.AutoMigrate(&model.User{})
}