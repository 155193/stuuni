package services_test

import (
	"../../db"
	"../services"
	"testing"
)

func TestAddHistoryEvent(t *testing.T) {
	teardownTestCase := db.SetupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesAddHistoryEvent {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := db.SetupSubTest(t)
			defer teardownSubTest(t)
			if err := services.AddHistoryEvent(tc.Model); (err != nil) != tc.Thereiserror {
				t.Fatalf("%v", err)
			}
			if tc.Thereiserror {
				if err := services.RemoveHistoryEvent(tc.Id); err != nil {
					t.Fatalf("%v", err)
				}
			}
		})
	}
}

func TestGetHistoryEvents(t *testing.T) {
	teardownTestCase := db.SetupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesGetHistoryEvents {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := db.SetupSubTest(t)
			defer teardownSubTest(t)
			if _, err := services.GetHistoryEvents(); (err != nil) != tc.Thereiserror {
				t.Fatalf("%v", err)
			}
		})
	}
}
