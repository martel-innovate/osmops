// Traversal of the repo target directory tree to process the content of any
// OSM GitOps files found in there.
//
package cfg

import (
	"io/fs"
	"io/ioutil"
	"strings"

	u "github.com/martel-innovate/osmops/osmops/util"
	"github.com/martel-innovate/osmops/osmops/util/file"
)

// KduNsActionFile is the data passed to the OSM GitOps file visitor.
type KduNsActionFile struct {
	FilePath file.AbsPath
	Content  *KduNsAction
}

// KduNsActionProcessor is a file visitor that is given, in turn, the content
// of each OSM GitOps file found in the target directory.
type KduNsActionProcessor interface {
	// Do something with the current OSM GitOps file, possibly returning an
	// error if something goes wrong.
	Process(file *KduNsActionFile) error
}

// KduNsActionRepoScanner has methods to let visitors process OSM GitOps files
// found while traversing the target directory.
type KduNsActionRepoScanner struct {
	targetDir file.AbsPath
	fileExt   []u.NonEmptyStr
	readFile  func(string) ([]byte, error) // (*)

	// (*) added for testability, so we can sort of mock stuff
}

// NewKduNsActionRepoScanner instantiates a KduNsActionRepoScanner to
// traverse the target directory configured in the given Store.
func NewKduNsActionRepoScanner(store *Store) *KduNsActionRepoScanner {
	return &KduNsActionRepoScanner{
		targetDir: store.RepoTargetDirectory(),
		fileExt:   store.OpsFileExtensions(),
		readFile:  ioutil.ReadFile,
	}
}

// Visit scans the repo's OSM Ops target directory recursively, calling the
// specified visitor with the content of each OSM Git Ops file found.
// For now the only kind of Git Ops file OSM Ops can process is a file
// containing KduNsAction YAML. Any I/O errors that happen while traversing
// the target directory tree get collected in the returned error buffer as
// VisitErrors. Ditto for I/O errors that happen when reading or validating
// a Git Ops file as well as any error returned by the visitor.
func (k *KduNsActionRepoScanner) Visit(visitor KduNsActionProcessor) []error {
	scanner := file.NewTreeScanner(k.targetDir)
	return scanner.Visit(func(node file.TreeNode) error {
		if !k.isGitOpsFile(node.FsMeta) {
			return nil
		}
		return k.visitFile(node.NodePath, visitor)
	})
}

func (k *KduNsActionRepoScanner) isGitOpsFile(info fs.FileInfo) bool {
	if !info.IsDir() {
		for _, ext := range k.fileExt {
			name := strings.ToLower(info.Name())
			if strings.HasSuffix(name, ext.Value()) {
				return true
			}
		}
	}
	return false
}

func (k *KduNsActionRepoScanner) visitFile(absPath file.AbsPath,
	visitor KduNsActionProcessor) error {
	var err error
	file := &KduNsActionFile{FilePath: absPath}

	yaml, err := k.readFile(absPath.Value())
	if err != nil {
		return err
	}
	content, err := readKduNsAction(yaml)
	if err != nil {
		return err
	}
	file.Content = content

	return visitor.Process(file)
}
