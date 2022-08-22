package file

import (
	"errors"
	"fmt"
	"io/fs"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"testing"
)

func findTestDataDir(dirIndex int) AbsPath {
	_, thisFileName, _, _ := runtime.Caller(1)
	enclosingDir := filepath.Dir(thisFileName)
	testDataDirName := fmt.Sprintf("test_%d", dirIndex)
	testDataDir := filepath.Join(enclosingDir, "visitor_test_dir",
		testDataDirName)
	p, _ := ParseAbsPath(testDataDir)

	return p
}

func assertPathInvariants(t *testing.T, node TreeNode) {
	if filepath.IsAbs(node.RelPath) {
		t.Errorf("want: rel path; got: %s", node.RelPath)
	}

	nodePath := node.NodePath.Value()
	joinedPath := path.Join(node.RootPath.Value(), node.RelPath)
	if nodePath != joinedPath {
		t.Errorf("want: root + rel = node; got: root + rel = %s, node = %s",
			joinedPath, nodePath)
	}

	nodeName := path.Base(nodePath)
	if node.FsMeta.Name() != nodeName {
		t.Errorf("want: node name = fs name; got: node name = %s, fs name = %s",
			nodeName, node.FsMeta.Name())
	}
}

func TestPathInvariants(t *testing.T) {
	for k := 1; k < 3; k++ {
		targetDir := findTestDataDir(k)
		scanner := NewTreeScanner(targetDir)
		es := scanner.Visit(func(node TreeNode) error {
			assertPathInvariants(t, node)
			return nil
		})
		if len(es) > 0 {
			t.Errorf("want: no errors; got: %v", es)
		}
	}
}

func TestNilVisitor(t *testing.T) {
	targetDir := findTestDataDir(1)
	scanner := NewTreeScanner(targetDir)
	es := scanner.Visit(nil)
	if len(es) != 1 {
		t.Errorf("want: nil visitor error; got: %v", es)
	}
}

func TestNonExistentTargetDir(t *testing.T) {
	targetDir := findTestDataDir(0)
	scanner := NewTreeScanner(targetDir)
	es := scanner.Visit(func(node TreeNode) error {
		return nil
	})
	if len(es) != 1 {
		t.Errorf("want: non-existent target dir error; got: %v", es)
	}
}

func TestVisitCollectWalkError(t *testing.T) {
	targetDir := findTestDataDir(1)
	scanner := NewTreeScanner(targetDir).(*scanner)
	es := []error{}

	fn := scanner.visitAllAndCollectErrors(
		func(n TreeNode) error {
			return nil
		}, &es)
	err := fmt.Errorf("I/O error while scanning the dir tree.")
	var info fs.FileInfo
	fn("/pa/th", info, err)

	if len(es) != 1 {
		t.Errorf("want: one error; got: %v", es)
	}
	want := &VisitError{AbsPath: "/pa/th", Err: err}
	if got, ok := es[0].(*VisitError); !ok || want.Error() != got.Error() {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

func TestCollectAllErrors(t *testing.T) {
	targetDir := findTestDataDir(2)
	scanner := NewTreeScanner(targetDir)
	es := scanner.Visit(func(node TreeNode) error {
		return fmt.Errorf("%s", node.NodePath.Value())
	})
	if len(es) != 10 {
		t.Errorf("want: one error for each node; got: %v", es)
	}
	for _, e := range es {
		if ve, ok := e.(*VisitError); !ok {
			t.Errorf("want: VisitError; got: %v", e)
		} else {
			originalErrMgs := ve.Unwrap().Error()
			if ve.AbsPath != originalErrMgs {
				t.Errorf("want: %s; got: %s", originalErrMgs, ve.AbsPath)
			}

			want := fmt.Sprintf("%s: %s", ve.AbsPath, ve.AbsPath)
			if want != e.Error() {
				t.Errorf("want: %s; got: %s", want, e.Error())
			}
		}
	}
}

func assertVisitedPaths(t *testing.T, dirIndex int, want []string) {
	visitedPaths := []string{}
	targetDir := findTestDataDir(dirIndex)
	scanner := NewTreeScanner(targetDir)
	es := scanner.Visit(func(node TreeNode) error {
		visitedPaths = append(visitedPaths, node.RelPath)
		return nil
	})

	if len(es) != 0 {
		t.Errorf("want: no errors; got: %v", es)
	}

	sort.Strings(visitedPaths)
	if !reflect.DeepEqual(want, visitedPaths) {
		t.Errorf("want: %v; got: %v", want, visitedPaths)
	}
}

func TestVisitFlatDir(t *testing.T) {
	want := []string{"", "f1", "f2"}
	assertVisitedPaths(t, 1, want)
}

func TestVisitDirTree(t *testing.T) {
	want := []string{
		"",
		"d1", "d1/f2", "d1/f3",
		"d2", "d2/d3", "d2/d3/f6", "d2/f4", "d2/f5",
		"f1",
	}
	assertVisitedPaths(t, 2, want)
}

func TestVisitErrorStringRepr(t *testing.T) {
	e := VisitError{AbsPath: "p", Err: fmt.Errorf("e")}
	want := "p: e"
	if e.Error() != want {
		t.Errorf("want: %s; got: %s", want, e)
	}
}

func TestVisitErrorUnwrapping(t *testing.T) {
	cause := fmt.Errorf("cause")
	e := VisitError{AbsPath: "p", Err: cause}
	got := errors.Unwrap(e)
	if cause != got {
		t.Errorf("want: %v; got: %v", cause, got)
	}
}
