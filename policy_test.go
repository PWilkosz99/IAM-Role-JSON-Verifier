package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestVerifyPolicy_ValidPolicy(t *testing.T) {
	jsonData := []byte(`{"PolicyDocument":{"Statement":[{"Resource":"example-resource"}]}}`)

	result := verifyPolicy(jsonData)

	if !result {
		t.Error("Expected policy to be valid, but got invalid")
	}
}

func TestVerifyPolicy_InvalidPolicyWithAsterisk(t *testing.T) {
	jsonData := []byte(`{"PolicyDocument":{"Statement":[{"Resource":"*"}]}}`)

	result := verifyPolicy(jsonData)

	if result {
		t.Error("Expected policy to be invalid, but got valid")
	}
}

func TestVerifyPolicy_InvalidPolicyWithNoResource(t *testing.T) {
	jsonData := []byte(`{"PolicyDocument":{"Statement":[{}]}}`)

	result := verifyPolicy(jsonData)

	if !result {
		t.Error("Expected policy to be valid, but got invalid")
	}
}

func TestVerifyPolicy_EmptyJSON(t *testing.T) {
	jsonData := []byte(`{}`)

	result := verifyPolicy(jsonData)

	if !result {
		t.Error("Expected policy to be valid, but got invalid")
	}
}

func TestVerifyPolicy_InvalidJSON(t *testing.T) {
	jsonData := []byte(`invalid-json`)

	result := verifyPolicy(jsonData)

	if result {
		t.Error("Expected policy to be invalid, but got valid")
	}
}

func TestVerifyPolicy_TestValidJSONFile1(t *testing.T) {
	filePath := filepath.Join("testdata", "policy-1.json")
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Error reading test JSON file: %v", err)
	}

	result := verifyPolicy(jsonData)

	if !result {
		t.Error("Expected policy to be valid, but got invalid")
	}
}

func TestVerifyPolicy_TestInvalidJSONFile2(t *testing.T) {
	filePath := filepath.Join("testdata", "policy-2.json")
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Error reading test JSON file: %v", err)
	}

	result := verifyPolicy(jsonData)

	if result {
		t.Error("Expected policy to be invalid, but got valid")
	}
}

func TestVerifyPolicy_TestInvalidJSONFile3(t *testing.T) {
	filePath := filepath.Join("testdata", "policy-3.json")
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Error reading test JSON file: %v", err)
	}

	result := verifyPolicy(jsonData)

	if result {
		t.Error("Expected policy to be invalid, but got valid")
	}
}

func TestVerifyPolicy_TestInvalidJSONFile4(t *testing.T) {
	filePath := filepath.Join("testdata", "policy-4.json")
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Error reading test JSON file: %v", err)
	}

	result := verifyPolicy(jsonData)

	if result {
		t.Error("Expected policy to be invalid, but got valid")
	}
}
