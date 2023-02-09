package osUtil

import (
	"os"
	"testing"
)

func TestReadWriteStrsFile(t *testing.T) {
	filePath := "/tmp/test.txt"

	// Test write and read a single line to the file
	expectedLines := []string{"test"}
	err := WriteStrsFile(expectedLines, filePath)
	if err != nil {
		t.Fatalf("Failed to write file: %v", err)
	}

	lines, err := ReadStsFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if len(lines) != len(expectedLines) {
		t.Fatalf("Expected %d lines, got %d", len(expectedLines), len(lines))
	}

	for i, line := range lines {
		if line != expectedLines[i] {
			t.Fatalf("Expected line %d to be %q, got %q", i, expectedLines[i], line)
		}
	}

	// Test write and read multiple lines to the file
	expectedLines = []string{"test1", "test2", "test3"}
	err = WriteStrsFile(expectedLines, filePath)
	if err != nil {
		t.Fatalf("Failed to write file: %v", err)
	}

	lines, err = ReadStsFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if len(lines) != len(expectedLines) {
		t.Fatalf("Expected %d lines, got %d", len(expectedLines), len(lines))
	}

	for i, line := range lines {
		if line != expectedLines[i] {
			t.Fatalf("Expected line %d to be %q, got %q", i, expectedLines[i], line)
		}
	}

	// Clean up the file after the test
	err = os.Remove(filePath)
	if err != nil {
		t.Fatalf("Failed to remove file: %v", err)
	}
}
