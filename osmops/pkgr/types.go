package pkgr

import (
	"io"
	"os"
	"path"
	"path/filepath"
	"sort"

	"github.com/martel-innovate/osmops/osmops/util/bytez"
	"github.com/martel-innovate/osmops/osmops/util/file"
)

// Package holds the data that make up an OSM package.
type Package struct {
	// The package name. Conventionally this is the same as the name of
	// the directory containing the source files. We also follow this
	// convention.
	Name string
	// Metadata about the package source files.
	Source PackageSource
	// Gzipped tar stream containing the package source files plus a
	// checksum file.
	//
	// Each source file is archived at a path "r/p" where r is the name of
	// the directory containing the package source files and p is the file's
	// path relative to r. For example, if a package source directory "my-pkg"
	// contained a file "d/f", f's path in the archive would be "my-pkg/d/f".
	//
	// The stream also contains a checksum file at path "r/checksums.txt",
	// where r is the name of the directory containing the package source
	// files. This file has an MD5 hash entry in correspondence of each file
	// found in the package source directory and subdirectories. Each entry
	// is a text line starting with the MD5 of the file, followed by a tab
	// and then the path of the file in the archive. Here's an example:
	//
	//     c122710acb043b99be209fefd9ae2032	my-pkg/README.md
	//     7044f64c16d4ef3eeef7f8668a4dc5a1	my-pkg/knf/vnfd.yaml
	//     6cbc0db17616eff57c60efa0eb15ac76	my-pkg/nsd.yaml
	//
	Data io.ReadCloser
	// MD5 hash of the whole gzipped tar stream.
	Hash string
}

func makePackage(src PackageSource, data *bytez.Buffer) *Package {
	return &Package{
		Name:   src.DirectoryName(),
		Source: src,
		Data:   data,
		Hash:   md5string(data.Bytes()),
	}
}

// PackageSource provides metadata about an OSM package's source files
// as well as their content.
type PackageSource interface {
	// The root directory containing the package source files.
	Directory() file.AbsPath
	// The name of the root directory.
	DirectoryName() string
	// Relative paths of the files in the package source directory and
	// subdirectories. Each path is prefixed by the package source directory's
	// name. For example, if "my-pkg" is the root and there's a file "f" at
	// "d/f", the corresponding path returned by this method is "my-pkg/d/f".
	// This method returns the paths sorted in alphabetical order.
	SortedFilePaths() []string
	// Lookup the MD5 hash of a source file in the package.
	// The filePath argument must be one of the paths returned by
	// SortedFilePaths.
	FileHash(filePath string) string
	// FileContent returns the bytes that make up the specified file in the
	// package.
	// The filePath argument must be one of the paths returned by
	// SortedFilePaths.
	FileContent(filePath string) ([]byte, error)
}

type pkgSrc struct {
	srcDir        file.AbsPath
	srcDirName    string
	pathToHashMap map[string]string
}

func newPkgSrc(srcDir file.AbsPath) *pkgSrc {
	return &pkgSrc{
		srcDir:        srcDir,
		srcDirName:    path.Base(srcDir.Value()),
		pathToHashMap: make(map[string]string),
	}
}

func (p *pkgSrc) Directory() file.AbsPath {
	return p.srcDir
}

func (p *pkgSrc) DirectoryName() string {
	return p.srcDirName
}

func (p *pkgSrc) SortedFilePaths() []string {
	keys := make([]string, 0, len(p.pathToHashMap))
	for k := range p.pathToHashMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys
}

func (p *pkgSrc) FileHash(filePath string) string {
	if hash, ok := p.pathToHashMap[filePath]; ok {
		return hash
	}
	return ""
}

func (p *pkgSrc) addFileHash(node file.TreeNode) error {
	if !node.FsMeta.Mode().IsRegular() {
		return nil
	}
	hash, err := computeChecksum(node.NodePath)
	if err == nil {
		baseNamePlusPath := path.Join(p.srcDirName, node.RelPath)
		p.pathToHashMap[baseNamePlusPath] = hash
	}
	return err
}

func (p *pkgSrc) FileContent(filePath string) ([]byte, error) {
	filePathFromSrcDir, _ := filepath.Rel(p.srcDirName, filePath)
	absPath := p.srcDir.Join(filePathFromSrcDir)
	return os.ReadFile(absPath.Value())
}
