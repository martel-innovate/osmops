package tgz

import (
	"io/fs"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/martel-innovate/osmops/osmops/util/file"
)

func makeTarballPath(t *testing.T, tempDirPath string) file.AbsPath {
	tarball, err := file.ParseAbsPath(path.Join(tempDirPath, "test.tgz"))
	if err != nil {
		t.Fatalf("couldn't build tarball pathname: %v", err)
	}
	return tarball
}

func writeBogusTarball(t *testing.T, tempDirPath string) file.AbsPath {
	tarball := makeTarballPath(t, tempDirPath)

	data := []byte{1, 2, 3}
	if err := os.WriteFile(tarball.Value(), data, os.ModePerm); err != nil {
		t.Fatalf("couldn't write tarball: %v", err)
	}

	return tarball
}

func writeFlatTarball(t *testing.T, tempDirPath string) file.AbsPath {
	tarball := makeTarballPath(t, tempDirPath)

	dest, err := os.OpenFile(tarball.Value(),
		os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		t.Fatalf("couldn't open tarball for writing: %v", err)
	}

	writer, err := NewWriter("", dest, WithBestSpeed())
	if err != nil {
		t.Fatalf("couldn't create tarball writer: %v", err)
	}
	defer writer.Close()

	if err = writer.AddEntry("foo", strings.NewReader("bar")); err != nil {
		t.Fatalf("couldn't write tarball entry: %v", err)
	}

	return tarball
}

func TestExtractTarballFileOpenErr(t *testing.T) {
	withTempDir(t, func(tempDirPath string) {
		tarball := writeBogusTarball(t, tempDirPath)
		os.Chmod(tarball.Value(), 0200) // can't read

		err := ExtractTarball(tarball, tempDirPath)
		if _, ok := err.(*fs.PathError); !ok {
			t.Errorf("want: file open error; got: %v", err)
		}
	})
}

func TestExtractTarballMalformedFileErr(t *testing.T) {
	withTempDir(t, func(tempDirPath string) {
		tarball := writeBogusTarball(t, tempDirPath)

		err := ExtractTarball(tarball, tempDirPath)
		wantMgs := "unexpected EOF"
		if err == nil || err.Error() != wantMgs { // (*)
			t.Errorf("want: malformed file error; got: %v", err)
		}
	})

	// NOTE. Error type. The gzip pkg returns a generic errors.errorString,
	// so the best we can do is check for the error message.
}

func TestExtractTarballWriteEntryErr(t *testing.T) {
	withTempDir(t, func(tempDirPath string) {
		tarball := writeFlatTarball(t, tempDirPath)
		os.Chmod(tempDirPath, 0500) // ExtractTarball can't write in here

		err := ExtractTarball(tarball, tempDirPath)
		if _, ok := err.(*fs.PathError); !ok {
			t.Errorf("want: file entry write error; got: %v", err)
		}
	})
}

func TestExtractTarballCreateDestDirErr(t *testing.T) {
	withTempDir(t, func(tempDirPath string) {
		destDirPath := path.Join(tempDirPath, "dest")
		tarball := writeFlatTarball(t, tempDirPath)
		os.Chmod(tempDirPath, 0500) // ExtractTarball can't write in here

		err := ExtractTarball(tarball, destDirPath)
		if _, ok := err.(*fs.PathError); !ok {
			t.Errorf("want: file entry write error; got: %v", err)
		}
	})
}

func TestEnsureDir(t *testing.T) {
	withTempDir(t, func(tempDirPath string) {
		dirPath := path.Join(tempDirPath, "foo")
		if err := os.Mkdir(dirPath, fs.ModePerm|fs.ModeDir); err != nil {
			t.Fatalf("couldn't create %s: %v", dirPath, err)
		}

		fi, err := os.Stat(dirPath)
		if err != nil {
			t.Fatalf("couldn't stat %s: %v", dirPath, err)
		}

		if err = ensureDirs(fi, dirPath); err != nil {
			t.Errorf("want: no error; got: %v", err)
		}
	})
}
