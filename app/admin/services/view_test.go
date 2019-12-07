package services_test

import (
	"../services"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestAddView(t *testing.T) {
	services.Colview = "modules" //use other collection for tests
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesAddView {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if err := services.AddView(tc.Id, tc.Model); (err != nil) != tc.Thereiserror {
				t.Fatalf("%v", err)
			}
		})
	}
}

func TestUpdateView(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesUpdateView {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if err := services.UpdateView(tc.Id, tc.Nam, tc.Des, tc.Url, tc.Icon); (err != nil) != tc.Thereiserror {
				t.Fatalf("%v", err)
			}
			if !tc.Thereiserror {
				viewM, err := services.GetView(tc.Id)
				if err != nil {
					t.Fatalf("%v", err)
				}
				if !cmp.Equal(tc.Model, viewM.ToBsonM()) {
					t.Fatal("Error: service updateView not correct")
				}
			}
		})
	}
}

func TestViewOfUser(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesViewofUser {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if _, err := services.ViewOfUser(tc.Id); (err != nil) != tc.Thereiserror {
				t.Fatalf("%v", err)
			}
		})
	}
}

func TestRemoveView(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesRemoveView {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if bol, err := services.RemoveView(tc.Id); bol == tc.Thereiserror {
				t.Fatalf("%v", err)
			}
		})
	}
}
