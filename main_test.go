package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

type version struct {
	major, minor, patch int
}

// Only one version 0.0.1 not 10.10.10
func newVersion(name string) *version {
	s := regexp.MustCompile(`(?m)(\d)\.(\d)\.(\d)`).FindStringSubmatch(name)

	v := &version{}

	if len(s) >= 4 {
		v.major, _ = strconv.Atoi(s[1]) //no lint
		v.minor, _ = strconv.Atoi(s[2]) //no lint
		v.patch, _ = strconv.Atoi(s[3]) //no lint
	} else {
		return nil
	}

	return v
}

func TestVersionInChangeLog(t *testing.T) {
	t.Parallel()
	changeLogFile := "change.log"

	shaOut, err := exec.Command("git", "rev-list", "--tags", "--max-count=1").Output()
	if err != nil {
		t.Error(string(shaOut))
		t.Fatal("git rev-list: ", err)
	}

	sha := strings.TrimSuffix(string(shaOut), "\n")

	out, err := exec.Command("git", "describe", "--tags", "--always", sha).Output()
	if err != nil {
		t.Error(string(out))
		t.Fatal("git describe: ", err)
	}

	tag := newVersion(string(out))
	if tag == nil {
		t.Fatal("not found version")
	}

	//Read change log
	f, err := os.Open(changeLogFile)
	if err != nil {
		t.Fatal("read change log")
	}
	defer f.Close()

	_, err = os.Stat(changeLogFile)
	if err != nil {
		t.Fatal("not found change log")
	}

	var buf bytes.Buffer
	_, err = buf.ReadFrom(f)
	if err != nil {
		t.Fatal("read change log")
	}

	allVersions := regexp.MustCompile(`\n(\d)\.(\d)\.(\d)`).FindAllString(buf.String(), -1)

	// Last version in change log
	code := newVersion(allVersions[len(allVersions)-1])
	if code == nil {
		t.Fatal("not found version in changelog")
	}

	switch {
	case code.major > tag.major: // 2.*.* > 1.*.*
		break
	case code.major < tag.major: // 1.*.* < 2.*.* bad
		t.Error("Bad MAJOR version")
	case code.minor > tag.minor: // 1.1.* > 1.0.*
		break
	case code.minor < tag.minor: // 1.0.* < 1.1.* bad
		t.Error("Bad MINOR version")
	case code.patch < tag.patch: // 1.0.0 < 1.0.1 bad
		t.Error("Bad PATCH version")
	}

}

func TestStartMain(t *testing.T) {
	configFile := "/tmp/config.main"
	os.Args = []string{"cmd", fmt.Sprintf("--cfgFile=%s", configFile)}

	file, err := os.Create(configFile)
	if err != nil {
		t.Fatalf("Error creating test config file: %v", err)
	}
	defer file.Close()

	_, err = file.Write([]byte(`firewall=iptables`))
	if err != nil {
		t.Fatalf("Error writing test config file: %v", err)
	}

	err = file.Sync()
	if err != nil {
		t.Fatalf("Error syncing test config file: %v", err)
	}

	main()

	err = os.Remove(configFile)
	if err != nil {
		t.Fatalf("Failed to remove file: %v", err)
	}
}
