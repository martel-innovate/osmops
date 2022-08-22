package file

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

// TreeNode holds information about a filesystem node traversed by a
// TreeScanner.
type TreeNode struct {
	// The absolute path of the target directory being scanned.
	RootPath AbsPath
	// The absolute path of the node currently being visited.
	NodePath AbsPath
	// The path, relative to RootPath, of the node currently being visited.
	// It'll be the empty string for the root node, i.e. the target directory,
	// whereas it'll be the path to the current node from the target directory
	// for any other node, e.g. "some/dir", "some/dir/file", etc.
	// Also, for each visited node, including the target directory, you always
	// have: RootPath + RelPath = NodePath.
	RelPath string
	// Filesystem metadata about the node currently being visited.
	FsMeta fs.FileInfo
}

// Visitor is a function the TreeScanner calls on traversing each node
// in a given directory tree.
type Visitor func(TreeNode) error

// TreeScanner traverses a directory tree calling a visitor on each node.
type TreeScanner interface {
	// Visit scans a given target directory recursively, calling the
	// specified visitor on each filesystem node in the directory tree.
	// Any I/O errors that happen while traversing the target directory
	// tree get collected in the returned error buffer as VisitErrors.
	// Ditto for any error returned by the visitor.
	Visit(v Visitor) []error
}

// VisitError wraps any error that happened while traversing the target
// directory with an additional path to indicate where the error happened.
type VisitError struct {
	AbsPath string
	Err     error
}

// Error implements the standard error interface.
func (e VisitError) Error() string {
	return fmt.Sprintf("%s: %v", e.AbsPath, e.Err)
}

// Unwrap implements Go's customary error unwrapping.
func (e VisitError) Unwrap() error { return e.Err }

type scanner struct {
	targetDir AbsPath
}

// NewTreeScanner returns a TreeScanner to traverse the specified directory.
func NewTreeScanner(targetDir AbsPath) TreeScanner {
	return &scanner{targetDir: targetDir}
}

func (s *scanner) Visit(v Visitor) []error {
	es := []error{}
	if v != nil {
		filepath.Walk(s.targetDir.Value(), // (*)
			s.visitAllAndCollectErrors(v, &es))
	} else {
		es = appendVisitError(s.targetDir.Value(), fmt.Errorf("nil visitor"),
			es)
	}
	return es

	// (*) b/c targetDir is absolute, so is the path parameter passed to
	// the lambda returned by visitAllAndCollectErrors---see Walk docs.
}

func (s *scanner) visitAllAndCollectErrors(
	visit Visitor, acc *[]error) filepath.WalkFunc {
	return func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			*acc = appendVisitError(path, err, *acc)
			return nil
		}
		node := TreeNode{
			RootPath: s.targetDir,
			NodePath: AbsPath{data: path}, // see note above about targetDir
			RelPath:  extractRelPath(s.targetDir.Value(), path),
			FsMeta:   info,
		}
		if err := visit(node); err != nil {
			*acc = appendVisitError(path, err, *acc)
		}
		return nil
	}
}

func extractRelPath(base, node string) string {
	sep := string(filepath.Separator)
	pathFromBase := strings.TrimPrefix(node, base)
	return strings.TrimPrefix(pathFromBase, sep)
}

func appendVisitError(path string, err error, errors []error) []error {
	visitError := &VisitError{AbsPath: path, Err: err}
	return append(errors, visitError)
}
