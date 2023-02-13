package storage

import (
	"os"
	"testing"
)

func TestCounters(t *testing.T) {
	c := NewSyncMap()

	c.Increment("key1")
	if c.Load("key1") != 1 {
		t.Errorf("Expected value for key1 to be 1, got %d", c.Load("key1"))
	}

	c.Increment("key1")
	if c.Load("key1") != 2 {
		t.Errorf("Expected value for key1 to be 2, got %d", c.Load("key1"))
	}
}

func TestSaveAndRead(t *testing.T) {
	sm := NewSyncMap()
	sm.Save("counter1", 100)
	sm.Save("counter2", 200)

	// Save the map to file
	fileMap := "/tmp/test_map.json"
	err := sm.WriteToFile(fileMap)
	if err != nil {
		t.Fatalf("Error saving the map to file: %v", err)
	}

	// Read the map from file into a new map
	sm2 := NewSyncMap()
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
