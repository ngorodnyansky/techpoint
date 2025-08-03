package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestGetView(t *testing.T) {
	entries, err := os.ReadDir("./tests_data")
	if err != nil {
		t.Fatal(err)
	}

	var in *bufio.Reader
	for _, entry := range entries {
		if !strings.HasSuffix(entry.Name(), ".a") {
			file, err := os.Open("./tests_data/" + entry.Name())
			if err != nil {
				t.Error(err)
				continue
			}
			defer file.Close()

			in = bufio.NewReader(file)
			got := GetHexagons(in)

			data, err := os.ReadFile("./tests_data/" + entry.Name() + ".a")
			if err != nil {
				t.Error(err)
				continue
			}

			want := strings.Split(strings.TrimSpace(string(data)), "\n")
			for i := 0; i < len(got); i++ {
				gotLine := strings.TrimSpace(string(got[i]))
				wantLine := strings.TrimSpace(want[i])
				if gotLine != wantLine {
					t.Errorf("Test %s failed at line %d:\ngot:  '%s'\nwant: '%s'",
						entry.Name(), i+1, gotLine, wantLine)
					return
				}
			}
		}
	}
}
