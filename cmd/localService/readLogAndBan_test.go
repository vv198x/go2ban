package localService

import (
	"bytes"
	"context"
	"github.com/vv198x/go2ban/cmd/firewall"
	"github.com/vv198x/go2ban/config"
	"github.com/vv198x/go2ban/pkg/storage"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func Test_serviceWork_checkLogAndBlock(t *testing.T) {
	firewall.ExportFirewall = &firewall.Mock{}
	// In-memory cache
	countFailsMap := storage.NewCountersMap()
	endBytesMap := storage.NewStorageMap()
	testFile := "/tmp/test.log"
	testData := `
Feb 20 11:09:41 fedora sudo[2365]: Failed password 1.2.3.4 ion closed for user root
Feb 20 11:09:41 fedora sudo[2365]: Failed password 1.2.3.4 ion closed for user root`
	config.Get().ServiceFails = 2
	testMapFile := "/tmp/test_map.file"

	type fields struct {
		Name   string
		FindSt [][]byte
	}
	type args struct {
		ctx           context.Context
		sysFile       string
		countFailsMap storage.SyncMap
		endBytesMap   storage.SyncMap
	}
	defaultArg := args{
		ctx:           context.Background(),
		countFailsMap: countFailsMap,
		endBytesMap:   endBytesMap,
	}
	tests := []struct {
		name   string
		file   string
		inLog  string
		fields fields
		args   args //Find string log
	}{
		{"Valid file", "/etc/fstab", "", fields{Name: "Test 1", FindSt: stringToBytes("test")}, defaultArg},
		{"Not valid file", "/etc/fstab1", "can't open log file", fields{Name: "Test 2", FindSt: stringToBytes("test")}, defaultArg},
		{"Bloking ip in test file", testFile, "Mock firewall blocked ip: 1.2.3.4", fields{Name: "Test 3", FindSt: stringToBytes("Failed password")}, defaultArg},
	}

	file, err := os.Create(testFile)
	if err != nil {
		t.Fatalf("Error creating test config file: %v", err)
	}
	defer file.Close()

	_, err = file.Write([]byte(testData))
	if err != nil {
		t.Fatalf("Error writing test config file: %v", err)
	}

	err = file.Sync()
	if err != nil {
		t.Fatalf("Error syncing test config file: %v", err)
	}

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &serviceWork{
				Name:   tt.fields.Name,
				FindSt: tt.fields.FindSt,
			}
			s.checkLogAndBlock(tt.args.ctx, tt.file, tt.args.countFailsMap, tt.args.endBytesMap)
			time.Sleep(time.Millisecond * 200)
			if !strings.Contains(buf.String(), tt.inLog) {
				t.Errorf("Log = %v, don't have = %v", buf.String(), tt.inLog)
			}
		})
	}

	t.Run("Test map read and write", func(t *testing.T) {
		if err = endBytesMap.WriteToFile(testMapFile); err != nil {
			t.Errorf("Don`t save map")
		}

		endBytesMap = storage.NewStorageMap()

		if err = endBytesMap.ReadFromFile(testMapFile); err != nil {
			t.Errorf("Don`t read map")
		}
	})

	t.Run("Test key map and end byte", func(t *testing.T) {
		key := tests[2].fields.Name + tests[2].file
		fileSize := len(testData)
		if endByte := endBytesMap.Load(key); endByte != int64(fileSize) {
			t.Errorf("Wrong size in map. Got %v | Want %v", endByte, fileSize)
		}
	})

	err = os.Remove(testFile)
	if err != nil {
		t.Fatalf("Failed to remove file: %v", err)
	}
}

func stringToBytes(s string) [][]byte {
	bytes := make([][]byte, len(s))
	for i := range s {
		bytes[i] = []byte{s[i]}
	}
	return bytes
}
