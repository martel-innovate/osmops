package tgz

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/martel-innovate/osmops/osmops/util/file"
)

// ExtractTarball extracts the files in the given tarball to the specified
// directory, taking care of creating intermediate directories as needed.
//
// If you pass the empty string for destDirPath, ExtractTarball preserves
// the original archive paths, even if they're absolute. For example, if
// "/d/f" is the path of file f in the archive, ExtractTarball will try
// creating a directory "/d" if it doesn't exist and then put f in there.
// With an empty destDirPath, ExtractTarball resolves relative archive paths
// with respect to the current directory. For example, if "d/f" is the path
// of file f in the archive, ExtractTarball will try creating a directory
// "./d" if it doesn't exist and then put f in there.
//
// On the other hand, if you specify a destDirPath (either absolute or
// relative to the current directory), ExtractTarball recreates the directory
// structure of the archived files entirely in destDirPath by interpreting
// all archive paths (even absolute ones) relative to destDirPath. For example,
// if "/d/f" is the path of file f in the archive, ExtractTarball will try
// creating a directory "destDirPath/d" if it doesn't exist and then put f
// in there. The same happens to relative paths. For example, if "d/f" is
// the path of file f in the archive, ExtractTarball will try creating a
// directory "destDirPath/d" if it doesn't exist and then put f in there.
func ExtractTarball(tarballPath file.AbsPath, destDirPath string) error {
	source, err := os.Open(tarballPath.Value())
	if err != nil {
		return err
	}

	reader, err := NewReader(source)
	if err != nil {
		return err
	}

	return reader.IterateEntries(makeEntryReader(destDirPath))
}

func makeEntryReader(destDirPath string) EntryReader {
	return func(archivePath string, fi os.FileInfo, content io.Reader) error {
		if fi.IsDir() {
			return nil
		}

		targetPath := filepath.Join(destDirPath, archivePath)
		if err := ensureDirs(fi, targetPath); err != nil {
			return err
		}

		fd, err := os.OpenFile(targetPath,
			os.O_CREATE|os.O_TRUNC|os.O_WRONLY, fi.Mode())
		if err != nil {
			return err
		}
		defer fd.Close()

		_, err = io.Copy(fd, content)
		return err
	}
}

func ensureDirs(fi os.FileInfo, targetPath string) error {
	if fi.IsDir() {
		return os.MkdirAll(targetPath, fi.Mode())
	}
	enclosingDir := filepath.Dir(targetPath)
	return os.MkdirAll(enclosingDir, fs.ModePerm|fs.ModeDir)
}
