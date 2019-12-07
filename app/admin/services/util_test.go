package services_test

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

//Global variables for nested documents
var idModule, _ = primitive.ObjectIDFromHex("5a3ed8f22c1e492f45cd2842")
var idRole, _ = primitive.ObjectIDFromHex("5a3ed8f22c1e492f45cd2843")
var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjVhM2VkOWYxMmMxZTQ5MmY0NWNkMjg0OCJ9.1fkwsoey7rtjw9-hl5r0GMXwwovJgJHGqr7Bt_hhxCI"

/* function for setup test case */
func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")
	return func(t *testing.T) {
		t.Log("teardown test case")
	}
}

/* function for setup Sub test */
func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("setup sub test")
	return func(t *testing.T) {
		t.Log("teardown sub test")
	}
}
