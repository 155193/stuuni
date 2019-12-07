package services_test

import (
	"../services"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

//Before to run access_test install "github.com/google/go-cmp/cmp"

func TestAddAccess(t *testing.T) {
	services.Colview = "modules" //use other collection for tests
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range casesAddAccess {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if err := services.AddModule(bson.M{"_id": tc.Id,}); (err != nil) != tc.Thereiserror {
				t.Fatalf("%v", err)
			}
			if err := services.AddRole(tc.Id, bson.M{"_id": tc.Idaux,}); (err != nil) != tc.Thereiserror {
				t.Fatalf("%v", err)
			}
			if err := services.AddAccess(tc.Id, tc.Idaux, tc.Model); (err != nil) != tc.Thereiserror {
				t.Fatalf("%v", err)
			}
		})
	}
}

func TestUpdateAccess(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range casesUpdateAccess {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if err := services.UpdateAccess(tc.Idaux, tc.Id, tc.Idaux, tc.Order, tc.Position); (err != nil) != tc.Thereiserror {
				t.Fatalf("%v", err)
			}
			//if !tc.Thereiserror {
			//	accessM, err := services.GetAccess(tc.Id)
			//	if err != nil {
			//		t.Fatalf("%v", err)
			//	}
			//	fmt.Println(tc.Model)
			//	fmt.Println(accessM.ToBson())
			//	if !cmp.Equal(tc.Model, accessM.ToBson()) {
			//		t.Fatal("Error: service updateAccess not correct")
			//	}
			//}
		})
	}
}

func TestRemoveAccess(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range casesRemoveAccess {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if bol, err := services.RemoveAccess(tc.Id, tc.Idaux); bol == tc.Thereiserror {
				t.Fatalf("%v", err)
			}
		})
	}
}
