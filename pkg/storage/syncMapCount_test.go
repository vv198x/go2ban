package storage

import "testing"

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
