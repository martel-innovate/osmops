package pkgr

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/martel-innovate/osmops/osmops/util/file"
)

func TestPackErrOnSourceDirAccess(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "pkgr-test")
	if err != nil {
		t.Fatalf("couldn't create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	sourceDir, _ := file.ParseAbsPath(tempDir)
	contentFile := path.Join(tempDir, "content")

	os.WriteFile(contentFile, []byte{}, 0200) // Pack's visitor can't open it

	_, err = Pack(sourceDir)
	if _, ok := err.(*file.VisitError); !ok {
		t.Errorf("want: visit error; got: %v", err)
	}
}

func TestWritePackageDataErrOnNilSink(t *testing.T) {
	srcDir, _ := file.ParseAbsPath("no/where")
	pkgSrc := newPkgSrc(srcDir)
	if err := writePackageData(pkgSrc, nil); err == nil {
		t.Errorf("want: nil error; got: no error")
	}
}
