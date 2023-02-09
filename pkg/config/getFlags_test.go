package config

import (
	"flag"
	"os"
	"testing"
)

func TestReadFlags(t *testing.T) {
	// Reset the flags before the test
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	// Test with default values

	ReadFlags()

	if exportCfg.Flags.RunAsDaemon != defaultRunDaemon {
		t.Errorf("Expected RunAsDaemon to be %t, got %t", defaultRunDaemon, exportCfg.Flags.RunAsDaemon)
	}
	if exportCfg.Flags.ConfigFile != defaultCfgFile {
		t.Errorf("Expected ConfigFile to be %s, got %s", defaultCfgFile, exportCfg.Flags.ConfigFile)
	}
	if exportCfg.Flags.UnlockAll != false {
		t.Errorf("Expected UnlockAll to be false, got %t", exportCfg.Flags.UnlockAll)
	}

	// Test with custom values
	os.Args = []string{os.Args[0], "-d", "-cfgFile=custom_file.conf", "-clear"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	ReadFlags()
	if exportCfg.Flags.RunAsDaemon != true {
		t.Errorf("Expected RunAsDaemon to be true, got %t", exportCfg.Flags.RunAsDaemon)
	}
	if exportCfg.Flags.ConfigFile != "custom_file.conf" {
		t.Errorf("Expected ConfigFile to be 'custom_file.conf', got %s", exportCfg.Flags.ConfigFile)
	}
	if exportCfg.Flags.UnlockAll != true {
		t.Errorf("Expected UnlockAll to be true, got %t", exportCfg.Flags.UnlockAll)
	}
}
