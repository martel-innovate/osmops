package tgz

import (
	"fmt"
	"path"
	"path/filepath"
	"reflect"
	"runtime"

	"io/ioutil"
	"os"
	"testing"

	"github.com/martel-innovate/osmops/osmops/util/file"
)

const ArchiveTestDirName = "archive_test_dir"

func findTestDataDir() file.AbsPath {
	_, thisFileName, _, _ := runtime.Caller(1)
	enclosingDir := filepath.Dir(thisFileName)
	testDataDir := filepath.Join(enclosingDir, ArchiveTestDirName)
	p, _ := file.ParseAbsPath(testDataDir)

	return p
}

func withTempDir(t *testing.T, do func(string)) {
	if tempDir, err := ioutil.TempDir("", "tgz-test"); err != nil {
		t.Errorf("couldn't create temp dir: %v", err)
	} else {
		defer os.RemoveAll(tempDir)
		defer os.Chmod(tempDir, 0700) // make sure you can remove it
		do(tempDir)
	}
}

func checkExtractedPaths(t *testing.T, sourceDir file.AbsPath, extractedDir string) {
	want, _ := file.ListPaths(sourceDir.Value())
	got, _ := file.ListPaths(extractedDir)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

func TestTgzThenExtract(t *testing.T) {
	withTempDir(t, func(tempDirPath string) {
		sourceDir := findTestDataDir()
		tarball, _ := file.ParseAbsPath(path.Join(tempDirPath, "test.tgz"))
		extractedDir := path.Join(tempDirPath, ArchiveTestDirName)

		MakeTarball(sourceDir, tarball)
		if err := ExtractTarball(tarball, tempDirPath); err != nil {
			t.Fatalf("want: extract; got: %v", err)
		}
		checkExtractedPaths(t, sourceDir, extractedDir)
	})
}

func checkFileContent(pathname string) error {
	name := path.Base(pathname)
	content, err := ioutil.ReadFile(pathname)
	if err != nil {
		return err
	}
	text := string(content)
	if name != text {
		return fmt.Errorf("path = %s; content = %s", pathname, text)
	}
	return nil
}

func checkExtractedFiles(t *testing.T, extractedDir string) {
	targetDir, _ := file.ParseAbsPath(extractedDir)
	scanner := file.NewTreeScanner(targetDir)
	es := scanner.Visit(func(node file.TreeNode) error {
		if !node.FsMeta.IsDir() {
			return checkFileContent(node.NodePath.Value())
		}
		return nil
	})
	if len(es) > 0 {
		t.Errorf("want no errors; got: %v", es)
	}
}

func TestTgzThenExtractContent(t *testing.T) {
	withTempDir(t, func(tempDirPath string) {
		sourceDir := findTestDataDir()
		tarball, _ := file.ParseAbsPath(path.Join(tempDirPath, "test.tgz"))
		extractedDir := path.Join(tempDirPath, ArchiveTestDirName)

		MakeTarball(sourceDir, tarball)
		if err := ExtractTarball(tarball, tempDirPath); err != nil {
			t.Fatalf("want: extract; got: %v", err)
		}
		checkExtractedFiles(t, extractedDir)
	})
}

func TestExtractTgzCreatedWithUnixTar(t *testing.T) {
	withTempDir(t, func(tempDirPath string) {
		sourceDir := findTestDataDir()
		tarball, _ := file.ParseAbsPath(sourceDir.Value() + ".tgz")
		extractedDir := path.Join(tempDirPath, ArchiveTestDirName)

		if err := ExtractTarball(tarball, tempDirPath); err != nil {
			t.Fatalf("want: extract; got: %v", err)
		}
		checkExtractedPaths(t, sourceDir, extractedDir)
		checkExtractedFiles(t, extractedDir)
	})
}
