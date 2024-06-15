package main

import (
	"encoding/json"
	"os"
	"testing"
)

func TestLoadData(t *testing.T) {
	// Prepare test data
	testData := Data{
		Users: []User{
			{Username: "testuser"},
		},
	}
	testFilename := "test_data.json"
	bytes, _ := json.Marshal(testData)
	os.WriteFile(testFilename, bytes, 0644)
	defer os.Remove(testFilename)

	// Run the test
	data, err := loadData(testFilename)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(data.Users) != 1 || data.Users[0].Username != "testuser" {
		t.Fatalf("expected user 'testuser', got %v", data.Users)
	}
}

func TestSaveData(t *testing.T) {
	// Prepare test data
	testData := Data{
		Users: []User{
			{Username: "testuser"},
		},
	}
	testFilename := "test_save_data.json"
	defer os.Remove(testFilename)

	// Run the test
	err := saveData(testFilename, testData)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Verify the data was saved correctly
	data, err := loadData(testFilename)
	if err != nil {
		t.Fatalf("expected no error reading file, got %v", err)
	}

	if len(data.Users) != 1 || data.Users[0].Username != "testuser" {
		t.Fatalf("expected user 'testuser', got %v", data.Users)
	}
}

func TestFindUser(t *testing.T) {
	// Prepare test data
	testData := Data{
		Users: []User{
			{Username: "testuser"},
		},
	}

	// Run the test
	user, index := testData.findUser("testuser")
	if user == nil || index != 0 {
		t.Fatalf("expected to find user 'testuser' at index 0, got user %v, index %d", user, index)
	}

	// Test for a non-existent user
	user, index = testData.findUser("nonexistent")
	if user != nil || index != -1 {
		t.Fatalf("expected not to find user 'nonexistent', got user %v, index %d", user, index)
	}
}