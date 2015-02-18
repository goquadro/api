package main

import (
	"testing"

	"gopkg.in/mgo.v2/bson"
)

var testUser User

func TestsanitizeUrl(t *testing.T) {
	u := "http://www.goquadro.com"
	if _, err := sanitizeUrl(u); err != nil {
		t.Error("Address not correctly parsed")
	}
}

func TestAddDocument(t *testing.T) {
	d := Document{
		Url:   "http://github.com/goquadro",
		Title: "Goquadro on Github",
		Tags:  []string{"goquadro", "github", "sourcecode"},
	}
	err := testUser.AddDocument(&d)
	if err != nil {
		t.Error(err)
	}
}

func TestUserRegister(t *testing.T) {
	if !bson.IsObjectIdHex(testUser.ID.String()) {
		t.Errorf("Test user has not been persisted (no bson ID)")
	}
	u1 := User{
		Username: "testuser",
		Name:     "Big Lebowski",
		URL:      "http://www.goquadro.com",
		Email:    "spam@goquadro.com",
	}
	u2 := User{
		Username: testUser.Username,
		Name:     testUser.Name,
		URL:      testUser.URL,
		Email:    testUser.Email,
	}

	itemProperties := [...][2]string{
		[2]string{u1.Username, u2.Username},
		[2]string{u1.Name, u2.Name},
		[2]string{u1.URL, u2.URL},
		[2]string{u1.Email, u2.Email},
	}
	for _, comp := range itemProperties {
		if comp[0] != comp[1] {
			t.Errorf("Expected %v, got %v", comp[0], comp[1])
		}
	}
}

func init() {
	_ = testUser.Register(User{
		Username: "TestUser",
		Name:     "Big Lebowski",
		URL:      "http://www.goquadro.com",
		Email:    "spam@goquadro.com",
	})
}
