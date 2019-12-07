package services_test

import (
	"../services"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestAddModule(t *testing.T) {
	//services.Colview = "modules" //use other collection for tests
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesAddModule {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if err := services.AddModule(tc.Model); (err != nil) != tc.Thereiserror {
				t.Fatalf("%v", err)
			}
		})
	}
}

func TestUpdateModule(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesUpdateModule {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if err := services.UpdateModule(tc.Id,tc.Model); (err != nil) != tc.Thereiserror {
				t.Fatalf("%v", err)
			}
			if !tc.Thereiserror {
				ModuleM, err := services.GetModule(tc.Id)
				if err != nil {
					t.Fatalf("%v", err)
				}
				if !cmp.Equal(tc.Model, ModuleM.ToBson()) {
					t.Fatal("Error: service update Module not correct")
				}
			}
		})
	}
}

func TestGetModules(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesGetModules {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if _, err := services.GetModules(); (err != nil) != tc.Thereiserror {
				t.Fatalf("%v", err)
			}
		})
	}
}

func TestModulesOfUser(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesModulesofUser {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if _, err := services.ModulesOfUser(tc.IdList); (err != nil) != tc.Thereiserror {
				t.Fatalf("%v", err)
			}
		})
	}
}

func TestRemoveModule(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesRemoveModule {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if bol, err := services.RemoveModule(tc.Id); bol == tc.Thereiserror {
				t.Fatalf("%v", err)
			}
		})
	}
}
