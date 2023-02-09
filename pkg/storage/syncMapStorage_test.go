package storage

import (
	"os"
	"testing"
)

func TestStorageMap(t *testing.T) {
	sm := NewStorageMap()
	sm.Save("counter1", 100)
	sm.Save("counter2", 200)

	// Save the map to file
	fileMap := "/tmp/test_map.json"
	err := sm.WriteToFile(fileMap)
	if err != nil {
		t.Fatalf("Error saving the map to file: %v", err)
	}

	// Read the map from file into a new map
	sm2 := NewStorageMap()
	err = sm2.ReadFromFile(fileMap)
	if err != nil {
		t.Fatalf("Error reading the map from file: %v", err)
	}

	// Check if the values are the same
	v1 := sm2.Load("counter1")
	if v1 != 100 {
		t.Fatalf("Expected value 100, got %d", v1)
	}

	v2 := sm2.Load("counter2")
	if v2 != 200 {
		t.Fatalf("Expected value 200, got %d", v2)
	}

	// Clean up the test file
	err = os.Remove(fileMap)
	if err != nil {
		t.Fatalf("Failed to remove file: %v", err)
	}
}
