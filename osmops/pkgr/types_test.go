package pkgr

import (
	"io/fs"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/martel-innovate/osmops/osmops/util/file"
)

func TestFileHashFailedLookup(t *testing.T) {
	srcDir, _ := file.ParseAbsPath("no/where")
	pkgSrc := newPkgSrc(srcDir)
	got := pkgSrc.FileHash("not/there")
	if got != "" {
		t.Errorf("want: empty; got: %s", got)
	}
}

func TestFileContentReadsAllFileIntoMem(t *testing.T) {
	source := findTestDataDir("openldap_nested")
	targetFile := source.Join("knf/openldap_vnfd.yaml")

	want, err := os.ReadFile(targetFile.Value())
	if err != nil {
		t.Fatalf("couldn't read file: %v", targetFile)
	}

	pkg, err := Pack(source)
	if err != nil {
		t.Fatalf("couldn't pack: %v; error: %v", source, err)
	}

	archiveFilePath := ""
	for _, p := range pkg.Source.SortedFilePaths() {
		if strings.HasSuffix(targetFile.Value(), p) {
			archiveFilePath = p
			break
		}
	}
	got, err := pkg.Source.FileContent(archiveFilePath)
	if err != nil {
		t.Fatalf("want: read file; got: %v", err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

func TestFileContentErrWhenUsingPathNotRetrievedFromSortedFilePaths(t *testing.T) {
	source := findTestDataDir("openldap_nested")
	pkg, err := Pack(source)
	if err != nil {
		t.Fatalf("couldn't pack: %v; error: %v", source, err)
	}

	// missing pkg dir; SortedFilePaths would've returned:
	// - openldap_nested/knf/openldap_vnfd.yaml
	_, err = pkg.Source.FileContent("knf/openldap_vnfd.yaml")
	if _, ok := err.(*fs.PathError); !ok {
		t.Errorf("want: path error; got: %v", err)
	}

	// SortedFilePaths never returns empty path
	_, err = pkg.Source.FileContent("")
	if _, ok := err.(*fs.PathError); !ok {
		t.Errorf("want: path error; got: %v", err)
	}
}
