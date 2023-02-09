package osUtil

import "testing"

func TestCheckFile(t *testing.T) {
	tests := []struct {
		filePath string
		expected bool
	}{
		{"/bin/sh", true},
		{"test2.txt", false},
		{"/nonexistent/file.txt", false},
	}

	for i, test := range tests {
		result := CheckFile(test.filePath)
		if result != test.expected {
			t.Errorf("Test %d: Expected %t but got %t for file path %s", i, test.expected, result, test.filePath)
		}
	}
}
