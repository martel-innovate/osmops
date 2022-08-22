package tgz

import (
	"io"
	"os"
	"path"

	"github.com/martel-innovate/osmops/osmops/util/file"
)

// WriteFileArchive collects all the files in sourceDir (and its sub-dirs)
// and writes them to a gzip tar archive.
// Each file gets written to the archive at path "b/r" where b is sourceDir's
// base name and r is the file's path relative to sourceDir. For example, if
// sourceDir = "my/source" contains a file "my/source/d/f", that file gets
// archived at "source/d/f". The archive bytes get written to the give sink
// stream.
func WriteFileArchive(sourceDir file.AbsPath, sink io.WriteCloser) error {
	archiveBaseDirName := path.Base(sourceDir.Value())
	scanner := file.NewTreeScanner(sourceDir)
	writer, err := NewWriter(archiveBaseDirName, sink, WithBestCompression())
	if err != nil {
		return err
	}

	defer writer.Close()
	if es := scanner.Visit(writer.Visitor()); len(es) > 0 {
		return es[0]
	}
	return nil
}

// MakeTarball collects all the files in sourceDir (and its sub-dirs) and
// writes them to a gzip tar archive file. The archive file is created with
// 0644 permissions at the path specified by tarballPath.
// Each file in sourceDir gets written to the archive at path "b/r" where b
// is sourceDir's base name and r is the file's path relative to sourceDir.
// For example, if sourceDir = "my/source" contains a file "my/source/d/f",
// that file gets archived at "source/d/f".
func MakeTarball(sourceDir, tarballPath file.AbsPath) error {
	dest, err := os.OpenFile(tarballPath.Value(),
		os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	return WriteFileArchive(sourceDir, dest)
}
