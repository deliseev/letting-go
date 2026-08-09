package examples

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// Constraints: examples must stay narrow and regions must stay
// short enough to read as a single unit on the page.
const (
	maxLineLen     = 72
	maxRegionLines = 35
)

// region is one // region: ... // endregion: ... span, successfully
// paired.
type region struct {
	file      string // path relative to examples/
	name      string
	startLine int // 1-based, line of the "// region:" marker
	lines     int // code lines strictly between the markers
}

// parsedFile is the result of scanning one .go file for regions.
type parsedFile struct {
	path     string
	regions  []region
	problems []string // pairing/nesting/duplicate violations, "path:line: message"
}

func lintableFiles(t *testing.T) []string {
	t.Helper()
	var files []string
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, "_test.go") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return files
}

func parseFile(t *testing.T, path string) parsedFile {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	pf := parsedFile{path: path}
	lines := strings.Split(string(data), "\n")

	var open *region
	seen := map[string]bool{}
	problem := func(line int, format string, args ...any) {
		msg := fmt.Sprintf(format, args...)
		pf.problems = append(pf.problems, fmt.Sprintf("%s:%d: %s", path, line, msg))
	}

	for i, line := range lines {
		lineNo := i + 1
		trimmed := strings.TrimSpace(line)

		switch {
		case strings.HasPrefix(trimmed, "// region: "):
			name := strings.TrimPrefix(trimmed, "// region: ")
			if open != nil {
				problem(lineNo, "region %q opened before region %q was closed", name, open.name)
				continue
			}
			if seen[name] {
				problem(lineNo, "region %q defined more than once in this file", name)
			}
			open = &region{file: path, name: name, startLine: lineNo}

		case strings.HasPrefix(trimmed, "// endregion: "):
			name := strings.TrimPrefix(trimmed, "// endregion: ")
			switch {
			case open == nil:
				problem(lineNo, "endregion %q has no matching region", name)
			case open.name != name:
				problem(lineNo, "endregion %q does not match open region %q", name, open.name)
			default:
				open.lines = lineNo - open.startLine - 1
				pf.regions = append(pf.regions, *open)
				seen[name] = true
				open = nil
			}
		}
	}
	if open != nil {
		problem(open.startLine, "region %q has no matching endregion", open.name)
	}
	return pf
}

// TestLineLength enforces the ≤72 char line limit for example files
// (test infrastructure such as *_test.go is exempt: it never renders
// on the page).
func TestLineLength(t *testing.T) {
	for _, path := range lintableFiles(t) {
		data, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		for i, line := range strings.Split(string(data), "\n") {
			if n := len([]rune(line)); n > maxLineLen {
				t.Errorf("%s:%d: line is %d chars, want <= %d", path, i+1, n, maxLineLen)
			}
		}
	}
}

// TestRegionStructure enforces that every // region: has exactly one
// matching // endregion: with the same name, regions don't nest, and
// no name is defined twice in the same file.
func TestRegionStructure(t *testing.T) {
	for _, path := range lintableFiles(t) {
		pf := parseFile(t, path)
		for _, p := range pf.problems {
			t.Error(p)
		}
	}
}

// TestRegionSize enforces the ≤35 line limit per region.
func TestRegionSize(t *testing.T) {
	for _, path := range lintableFiles(t) {
		pf := parseFile(t, path)
		for _, r := range pf.regions {
			if r.lines > maxRegionLines {
				t.Errorf("%s:%d: region %q is %d lines, want <= %d",
					r.file, r.startLine, r.name, r.lines, maxRegionLines)
			}
		}
	}
}

// regionRefRE matches a {{< code ... >}} shortcode call and
// captures its attributes so file/region can be pulled out
// regardless of the order they're written in.
var regionRefRE = regexp.MustCompile(`\{\{<\s*code\s+([^>]*?)\s*/?>\}\}`)
var regionAttrRE = regexp.MustCompile(`(\w+)\s*=\s*"([^"]*)"`)

// contentRegionRefs walks content/ (a sibling of examples/) and
// returns the set of "file#region" pairs referenced by {{< code >}}
// shortcode calls.
func contentRegionRefs(t *testing.T) map[string]bool {
	t.Helper()
	refs := map[string]bool{}
	root := filepath.Join("..", "content")
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		for _, call := range regionRefRE.FindAllStringSubmatch(string(data), -1) {
			attrs := map[string]string{}
			for _, m := range regionAttrRE.FindAllStringSubmatch(call[1], -1) {
				attrs[m[1]] = m[2]
			}
			if attrs["file"] != "" && attrs["region"] != "" {
				refs[attrs["file"]+"#"+attrs["region"]] = true
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return refs
}

// TestNoOrphanedRegions enforces that every region defined under
// examples/ is referenced by at least one {{< code >}} call in
// content/ — a region nobody includes is dead weight nothing catches
// if it silently rots.
func TestNoOrphanedRegions(t *testing.T) {
	refs := contentRegionRefs(t)
	for _, path := range lintableFiles(t) {
		if strings.HasPrefix(path, "testdata"+string(filepath.Separator)) {
			continue // fixtures, not course content
		}
		pf := parseFile(t, path)
		for _, r := range pf.regions {
			key := r.file + "#" + r.name
			if !refs[key] {
				t.Errorf("%s:%d: region %q is not referenced by any {{< code >}} call in content/",
					r.file, r.startLine, r.name)
			}
		}
	}
}
