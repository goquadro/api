package main

import "testing"

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
	err := testUser.AddDocument(d)
	if err != nil {
		t.Error(err)
	}
}

func TestUserRegister(t *testing.T) {
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
	if u1 != u2 {
		t.Errorf("Expected %v, got %v", u1, testUser)
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
