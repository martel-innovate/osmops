package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type AbsPath struct{ data string }

func (d AbsPath) Value() string {
	return d.data
}

func IsStringPath(value interface{}) error {
	s, _ := value.(string)
	_, err := ParseAbsPath(s)
	return err
}

func ParseAbsPath(path string) (AbsPath, error) {
	path = strings.TrimSpace(path) // (*)
	if len(path) == 0 {
		return AbsPath{},
			errors.New("must be a non-empty, non-whitespace-only string")
	}
	if p, err := filepath.Abs(path); err != nil {
		return AbsPath{}, err
	} else {
		return AbsPath{data: p}, nil
	}

	// (*) Abs doesn't trim space, e.g. Abs('/a/b ') == '/a/b '.
}

func (d AbsPath) Join(relativePath string) AbsPath {
	rest := strings.TrimSpace(relativePath) // (1)
	return AbsPath{
		data: filepath.Join(d.Value(), rest), // (2)
	}

	// (1) Join doesn't trim space, e.g. Join("/a", "/b ") == "/a/b "
	// (2) In principle this is wrong since we don't know if relativePath
	// is a valid path according to the FS we're running on. (Join doesn't
	// check that.) So we could potentially return an inconsistent AbsPath.
	// Go's standard lib is quite weak in the handling of abstract paths,
	// i.e. independent of OS, so this is the best we can do. See e.g.
	// - https://stackoverflow.com/questions/35231846
}

func (d AbsPath) IsDir() error {
	if f, err := os.Stat(d.Value()); err != nil {
		return err
	} else {
		if !f.IsDir() {
			return fmt.Errorf("not a directory: %v", d.Value())
		}
	}
	return nil
}

// ListPaths collects, recursively, the paths of all the directories and
// files inside dirPath. Each collected path is relative to dirPath, so
// for example, if dirPath = "b" and f is a file at "b/d/f", then "d/f"
// gets returned. ListPaths sorts the returned paths in alphabetical
// order.
func ListPaths(dirPath string) ([]string, []error) {
	visitedPaths := []string{}
	errs := []error{}

	targetDir, err := ParseAbsPath(dirPath)
	if err != nil {
		errs = append(errs, err)
		return visitedPaths, errs
	}

	scanner := NewTreeScanner(targetDir)
	es := scanner.Visit(func(node TreeNode) error {
		if node.RelPath != "" {
			visitedPaths = append(visitedPaths, node.RelPath)
		}
		return nil
	})

	errs = append(errs, es...)
	sort.Strings(visitedPaths)

	return visitedPaths, errs
}

// ListSubDirectoryNames returns the names of any directory found just
// below dirPath. ListSubDirectoryNames sorts the returned names in
// alphabetical order. Also, ListSubDirectoryNames will return an empty
// list if an error happens.
func ListSubDirectoryNames(dirPath string) ([]string, error) {
	dirs := []string{}
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return dirs, err
	}

	for _, e := range entries {
		if e.IsDir() {
			dirs = append(dirs, e.Name())
		}
	}

	sort.Strings(dirs)

	return dirs, nil
}
