package examples

import (
	"bytes"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestBrokenExamples verifies every *_broken.go file still fails to
// compile with exactly the message recorded in its sibling .want
// file. The .want holds the raw stderr a reader sees in their
// terminal — not a paraphrase — so a Go release that reworded a
// diagnostic, or a stray edit to .want, fails this test instead of
// letting the course quote an error the compiler no longer produces.
func TestBrokenExamples(t *testing.T) {
	var files []string
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(path, "_broken.go") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(files) == 0 {
		t.Skip("no *_broken.go files yet")
	}

	for _, file := range files {
		t.Run(file, func(t *testing.T) {
			src, err := os.ReadFile(file)
			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Contains(src, []byte("//go:build broken")) {
				t.Fatalf("missing //go:build broken constraint")
			}

			wantPath := strings.TrimSuffix(file, ".go") + ".want"
			want, err := os.ReadFile(wantPath)
			if err != nil {
				t.Fatalf("no matching %s: every broken example needs one (%v)", filepath.Base(wantPath), err)
			}

			cmd := exec.Command("go", "build", "-tags", "broken", "-o", os.DevNull, file)
			var stderr bytes.Buffer
			cmd.Stderr = &stderr
			if runErr := cmd.Run(); runErr == nil {
				t.Fatalf("compiled without error — example is no longer broken, fix it or drop the tag")
			}

			got := strings.TrimSpace(stderr.String())
			wantStr := strings.TrimSpace(string(want))
			if got != wantStr {
				t.Fatalf("compiler output does not match %s\n--- want ---\n%s\n--- got ---\n%s",
					filepath.Base(wantPath), wantStr, got)
			}
		})
	}
}
