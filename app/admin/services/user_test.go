package services_test

import (
	"../services"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestAddUser(t *testing.T) {
	//services.Coluser = "users" //use other collection for tests
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesAddUser {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if err := services.AddUser(tc.Model); (err != nil) != tc.Thereiserror {
				t.Fatalf("%v", err)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesUpdateUser {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if err := services.UpdateUser(tc.Id, tc.Model); (err != nil) != tc.Thereiserror {
				t.Fatalf("%v", err)
			}
			if !tc.Thereiserror {
				ModuleM, err := services.GetUserById(tc.Id)
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

func TestAddRole2User(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesAddRole2User {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if bol, err := services.AddRole2User(tc.Id, tc.Idaux); bol == tc.Thereiserror {
				t.Fatalf("%v", err)
			}
		})
	}
}

func TestRemoveRole2User(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesRemoveRole2User {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if bol, err := services.RemoveRole2User(tc.Id, tc.Idaux); bol == tc.Thereiserror {
				t.Fatalf("%v", err)
			}
		})
	}
}

func TestGetUserById(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesGetuserbyId {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if _, err := services.GetUserById(tc.Id); (err != nil) != tc.Thereiserror {
				t.Fatalf("%v", err)
			}
		})
	}
}

func TestGetUserByDNI(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesGetUserbyDni {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if _, err := services.GetUserByDNI(tc.Dni); (err != nil) != tc.Thereiserror {
				t.Fatalf("%v", err)
			}
		})
	}
}

func TestGetUsers(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesGetUsers {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if _, err := services.GetUsers(); (err != nil) != tc.Thereiserror {
				t.Fatalf("%v", err)
			}
		})
	}
}

func TestLogin(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesLogin {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if _, err := services.Login(tc.Dni, tc.Dni); (err != nil) != tc.Thereiserror {
				t.Fatalf("%v", err)
			}
		})
	}
}

func TestLoginT(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesLoginT {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if _, err := services.LoginT(tc.Dni); (err != nil) != tc.Thereiserror {
				t.Fatalf("%v", err)
			}
		})
	}
}

func TestCreateTokenString(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CaseCreateTokenString {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if _, err := services.CreateTokenString(tc.Dni); (err != nil) != tc.Thereiserror {
				t.Fatalf("%v", err)
			}
		})
	}
}

func TestDecodeTokenString(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesDecodeTokenString {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if _, err := services.DecodeTokenString(tc.Dni); (err != nil) != tc.Thereiserror {
				t.Fatalf("%v", err)
			}
		})
	}
}

func TestRefreshPasswordUser(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesRefreshPasswordUser {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if _, err := services.RefreshPasswordUser(tc.Id); (err != nil) != tc.Thereiserror {
				t.Fatalf("%v", err)
			}
		})
	}
}

func TestUpdatePasswordUser(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesUpdatePasswordUser {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if _, err := services.UpdatePasswordUser(tc.Id, tc.Dni, tc.Icon); (err != nil) != tc.Thereiserror {
				t.Fatalf("%v", err)
			}
		})
	}
}

func TestRemoveUser(t *testing.T) {
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)
	for _, tc := range CasesRemoveUser {
		t.Run(tc.Name, func(t *testing.T) {
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)
			if bol, err := services.RemoveUser(tc.Id); bol == tc.Thereiserror {
				t.Fatalf("%v", err)
			}
		})
	}
}
