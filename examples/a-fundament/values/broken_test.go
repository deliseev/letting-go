package values

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestBrokenCases runs every case under testdata/broken/<case>/ and
// checks its output against want.txt. Each case has its own go.mod
// (testdata is ignored by go tooling, so these never join a normal
// go build ./... / go test ./...). A case's cmd file names the
// command to run (default "go build ./"); an empty expect-success
// file flips the expectation from "must fail" to "must succeed" —
// see examples/a-fundament/values/testdata/broken/escape-analysis.
func TestBrokenCases(t *testing.T) {
	root := filepath.Join("testdata", "broken")
	entries, err := os.ReadDir(root)
	if err != nil {
		t.Fatal(err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		caseName := entry.Name()
		caseDir := filepath.Join(root, caseName)

		t.Run(caseName, func(t *testing.T) {
			wantPath := filepath.Join(caseDir, "want.txt")
			want, err := os.ReadFile(wantPath)
			if err != nil {
				t.Fatalf("no want.txt for case %q: %v", caseName, err)
			}

			cmd := "go build ./"
			if data, err := os.ReadFile(filepath.Join(caseDir, "cmd")); err == nil {
				cmd = strings.TrimSpace(string(data))
			}

			expectSuccess := false
			if _, err := os.Stat(filepath.Join(caseDir, "expect-success")); err == nil {
				expectSuccess = true
			}

			c := exec.Command("sh", "-c", cmd)
			c.Dir = caseDir
			var output bytes.Buffer
			c.Stdout = &output
			c.Stderr = &output
			runErr := c.Run()

			if expectSuccess && runErr != nil {
				t.Fatalf("case %q: %q failed: %v\n%s", caseName, cmd, runErr, &output)
			}
			if !expectSuccess && runErr == nil {
				t.Fatalf("case %q: %q succeeded, want failure\n%s", caseName, cmd, &output)
			}

			got := strings.TrimSpace(output.String())
			wantStr := strings.TrimSpace(string(want))
			if got != wantStr {
				t.Fatalf("case %q: output does not match want.txt\n--- want ---\n%s\n--- got ---\n%s",
					caseName, wantStr, got)
			}
		})
	}
}
