package main_test

import (
	"./db"
	"fmt"
)

func init() {
	fmt.Println("-----------init------------")
	if _, _, err := db.ConnectDB(); err == nil {

	}

}

//func TestMain(t *testing.M) {
//
//	fmt.Println("-----start tests")
//
//	//session, err := db.ConnectDB()
//	////adminDB.DB = session
//	////secDB.DB = session
//	////camDB.DB = session
//	////csDB.DB = session
//	////syncDB.DB = session
//	//
//	//if err != nil {
//	//	panic(err)
//	//}
//	//
//	//defer session.Close()
//	////
//	////// Optional. Switch the session to a monotonic behavior.
//	//session.SetMode(mgo.Monotonic, true)
//	fmt.Println("-----tests finish")
//}
