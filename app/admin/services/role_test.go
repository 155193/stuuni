package services_test

import (
	"../services"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestAddRole(t *testing.T) {
	//services.Colview = "modules" //use other collection for tests
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesAddRole {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if err := services.AddRole(tc.Id, tc.Model); (err != nil) != tc.Thereiserror {
				t.Fatalf("%v", err)
			}
		})
	}
}

func TestUpdateRole(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesUpdateRole {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if err := services.UpdateRole(tc.Id, tc.Nam, tc.Des); (err != nil) != tc.Thereiserror {
				t.Fatalf("%v", err)
			}
			if !tc.Thereiserror {
				RoleM, err := services.GetRole(tc.Id)
				if err != nil {
					t.Fatalf("%v", err)
				}
				if !cmp.Equal(tc.Model, RoleM.ToBsonM()) {
					t.Fatal("Error: service updateView not correct")
				}
			}
		})
	}
}

func TestGetRolesOfUser(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesGetRolesofUser {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if _,err := services.GetRolesOfUser(tc.IdList); (err != nil) != tc.Thereiserror {
				t.Fatalf("%v", err)
			}
		})
	}
}

func TestRemoveRole(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesRemoveRole {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if bol, err := services.RemoveRole(tc.Id); bol == tc.Thereiserror {
				t.Fatalf("%v", err)
			}
		})
	}
}