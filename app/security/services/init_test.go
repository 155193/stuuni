package services_test

import (
	"../../db"
)

func init() {
	if db.Session == nil {
		if _, _, err := db.ConnectDB(); err != nil {
		}
	}
}
