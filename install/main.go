package main

import (
	"../app/db"
	"fmt"
)

//function for created and initial collections in mongo
func InitializeBd() {
	// creating indexes

}

/**
main function
*/
func main() {
	fmt.Println("installing ....")
	session, ctx, err := db.ConnectDB()
	if err != nil {
		panic(err)
	} else {
		defer session.Disconnect(ctx)
		_ = db.DropAndInitDataBase()
	}
	InitializeBd()
	fmt.Println("creating indexes---")
}
