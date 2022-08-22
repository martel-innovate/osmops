package pkgr

import (
	"io"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/martel-innovate/osmops/osmops/util/file"
	"github.com/martel-innovate/osmops/osmops/util/tgz"
)

func findTestDataDir(dataDirName string) file.AbsPath {
	_, thisFileName, _, _ := runtime.Caller(1)
	enclosingDir := filepath.Dir(thisFileName)
	testDataDir := filepath.Join(enclosingDir, "build_pkg_test_dir", dataDirName)
	p, _ := file.ParseAbsPath(testDataDir)

	return p
}

const wantOpenLdapNsChecksumContent = `
c122710acb043b99be209fefd9ae2032	openldap_ns/README.md
6cbc0db17616eff57c60efa0eb15ac76	openldap_ns/openldap_nsd.yaml
`

var wantOpenLdapNsPaths = []string{
	"openldap_ns/README.md", "openldap_ns/openldap_nsd.yaml",
	"openldap_ns/checksums.txt",
}

func TestPackOpenLdapNs(t *testing.T) {
	wantName := "openldap_ns"
	wantHash := "cae4506d23753ee95f21faf8c6f97eaa"
	verifyPackage(t, wantName, wantHash, wantOpenLdapNsChecksumContent,
		wantOpenLdapNsPaths)
}

const wantOpenLdapKnfChecksumContent = `
7044f64c16d4ef3eeef7f8668a4dc5a1	openldap_knf/openldap_vnfd.yaml
`

var wantOpenLdapKnfPaths = []string{
	"openldap_knf/openldap_vnfd.yaml", "openldap_knf/checksums.txt",
}

func TestPackOpenLdapKnf(t *testing.T) {
	wantName := "openldap_knf"
	wantHash := "95389ba9b38e9a76b66789217a178e75"
	verifyPackage(t, wantName, wantHash, wantOpenLdapKnfChecksumContent,
		wantOpenLdapKnfPaths)
}

const wantOpenLdapNestedChecksumContent = `
c122710acb043b99be209fefd9ae2032	openldap_nested/README.md
7044f64c16d4ef3eeef7f8668a4dc5a1	openldap_nested/knf/openldap_vnfd.yaml
6cbc0db17616eff57c60efa0eb15ac76	openldap_nested/openldap_nsd.yaml
`

var wantOpenLdapNestedPaths = []string{
	"openldap_nested/README.md", "openldap_nested/openldap_nsd.yaml",
	"openldap_nested/knf/openldap_vnfd.yaml",
	"openldap_nested/checksums.txt",
}

func TestPackOpenLdapNested(t *testing.T) {
	wantName := "openldap_nested"
	wantHash := "5c2c2e459d0997fbef089499a7976812"
	verifyPackage(t, wantName, wantHash, wantOpenLdapNestedChecksumContent,
		wantOpenLdapNestedPaths)
}

func verifyPackage(t *testing.T, wantName, wantHash, wantChecksum string,
	wantPaths []string) {
	source := findTestDataDir(wantName)
	epochStart := time.Unix(0, 0)                             // (*) see NOTE
	pkg, err := doPack(source, tgz.WithEntryTime(epochStart)) // (*) see NOTE
	if err != nil {
		t.Fatalf("want: no error; got: %v", err)
	}

	if pkg.Name != wantName {
		t.Errorf("want name: %s; got: %s", wantName, pkg.Name)
	}

	if pkg.Source.Directory().Value() != source.Value() {
		t.Errorf("want source: %v; got: %v", source, pkg.Source)
	}

	if pkg.Hash != wantHash {
		t.Errorf("want hash: %s; got: %s", wantHash, pkg.Hash)
	}

	gotPaths, gotChecksum := pathsAndChecksumFile(t, pkg.Data)
	wantChecksum = strings.TrimSpace(wantChecksum)
	gotChecksum = strings.TrimSpace(gotChecksum)
	if gotChecksum != wantChecksum {
		t.Errorf("want checksum: %s; got: %s", wantChecksum, gotChecksum)
	}
	sort.Strings(gotPaths)
	sort.Strings(wantPaths)
	if !reflect.DeepEqual(gotPaths, wantPaths) {
		t.Errorf("want paths: %v; got: %v", wantPaths, gotPaths)
	}
}

// NOTE. Reproducible package hashes.
// Even if the files that make up the package are always the same, their
// mod/change/access time can be different at different times---think of
// e.g. a fresh checkout, merging a branch, opening one of the files, etc.
// The tar entry tgz creates for each file includes the mod/change/access
// time. So to make our test reproducible we've got to make sure we always
// get the same hash, which is why we set those times to the epoch's start.

func pathsAndChecksumFile(t *testing.T, data io.ReadCloser) ([]string, string) {
	reader, err := tgz.NewReader(data)
	if err != nil {
		t.Fatalf("couldn't create tgz reader: %v", err)
	}
	defer reader.Close()

	paths := []string{}
	checksums := ""
	reader.IterateEntries(
		func(archivePath string, fi os.FileInfo, content io.Reader) error {
			paths = append(paths, archivePath)
			if strings.HasSuffix(archivePath, ChecksumFileName) {
				buf, _ := io.ReadAll(content)
				checksums = string(buf)
			}
			return nil
		})

	return paths, checksums
}
