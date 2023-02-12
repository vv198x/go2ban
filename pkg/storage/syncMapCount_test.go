package storage

import (
	"os"
	"testing"
)

func TestCounters(t *testing.T) {
	c := NewCountersMap()

	c.Increment("key1")
	if c.Load("key1") != 1 {
		t.Errorf("Expected value for key1 to be 1, got %d", c.Load("key1"))
	}

	c.Increment("key1")
	if c.Load("key1") != 2 {
		t.Errorf("Expected value for key1 to be 2, got %d", c.Load("key1"))
	}
}

func TestCountersNotUse(t *testing.T) {
	c := NewCountersMap()
	c.Save("counter1", 100)
	c.Save("counter2", 200)
	if c.Load("counter1") != 0 {
		t.Fatalf("Save work")
	}

	// Save the map to file
	fileMap := "/tmp/test_map.json"
	err := c.WriteToFile(fileMap)
	if err != nil {
		t.Fatalf("Error saving the map to file: %v", err)
	}

	// Read the map from file into a new map
	c2 := NewCountersMap()
	err = c2.ReadFromFile(fileMap)
	if err != nil {
		t.Fatalf("Error reading the map from file: %v", err)
	}

	// Clean up the test file
	err = os.Remove(fileMap)
	if err == nil {
		t.Fatalf("File created")
	}
}
