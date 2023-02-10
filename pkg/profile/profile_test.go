package proFile

import (
	"os"
	"testing"
)

func TestStart(t *testing.T) {
	modeTests := []struct {
		mode   string
		result string
	}{
		{"cpu", "cpu"},
		{"mem", "mem"},
		{"mutex", "mutex"},
		{"block", "block"},
		{"", ""},
	}

	for _, tt := range modeTests {
		pPROF := Start(tt.mode)
		if pPROF == nil && tt.mode != "" {
			t.Errorf("Expected profile to start for mode %s, but got nil", tt.mode)
		}
		if pPROF != nil {
			pPROF.Stop()
			file, err := os.Open(pprofPath + "/" + tt.mode + ".pprof")
			if err != nil {
				t.Errorf("Expected pprof file to be created for mode %s, but got error: %v", tt.mode, err)
			}
			file.Close()
			os.Remove(pprofPath + "/" + tt.mode + ".pprof")
		}
	}
}
