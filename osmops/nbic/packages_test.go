package nbic

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/martel-innovate/osmops/osmops/util/file"
)

func findTestDataDir(pkgDirName string) file.AbsPath {
	_, thisFileName, _, _ := runtime.Caller(1)
	enclosingDir := filepath.Dir(thisFileName)
	testDataDir := filepath.Join(enclosingDir, "packages_test_dir", pkgDirName)
	p, _ := file.ParseAbsPath(testDataDir)

	return p
}

func md5string(data []byte) string {
	hash := md5.Sum(data)
	return fmt.Sprintf("%x", hash)
}

func callCreateOrUpdatePackage(pkgDirName string) (*mockNbi, error) {
	nbi := newMockNbi()
	urls := newConn()
	nbic, _ := New(urls, usrCreds, nbi.exchange)
	pkgSrc := findTestDataDir(pkgDirName)

	return nbi, nbic.CreateOrUpdatePackage(pkgSrc)
}

func checkUploadedPackage(t *testing.T, mockNbi *mockNbi, req *http.Request,
	pkgDirName, osmPkgId string) {
	gotFilename := req.Header.Get("Content-Filename")
	gotHash := req.Header.Get("Content-File-MD5")
	gotPkgTgzData := mockNbi.packages[osmPkgId]

	wantFilename := fmt.Sprintf("%s.tar.gz", pkgDirName)
	if gotFilename != wantFilename {
		t.Errorf("want file: %s; got: %s", wantFilename, gotFilename)
	}
	wantHash := md5string(gotPkgTgzData)
	if gotHash != wantHash {
		t.Errorf("want hash: %s; got: %s", wantHash, gotHash)
	}
}

func checkUploadedPackageDesc(t *testing.T, mockNbi *mockNbi,
	pkgDirName, descFilePath, osmPkgId string) {
	descFile := findTestDataDir(pkgDirName).Join(descFilePath)
	gotDescData := mockNbi.packages[osmPkgId]
	wantDescData, err := os.ReadFile(descFile.Value())

	if err != nil {
		t.Fatalf("couldn't read file: %v", descFile)
	}
	if !reflect.DeepEqual(wantDescData, gotDescData) {
		t.Errorf("want desc: %v; got: %v", wantDescData, gotDescData)
	}
}

func checkUnsupportedPackageErr(t *testing.T, err error) {
	if err == nil {
		t.Fatalf("want err; got: nil")
	}
	if !strings.HasPrefix(err.Error(), "unsupported package type") {
		t.Errorf("want unsupported pkg err; got: %v", err)
	}
}

func runCreatePackageTest(t *testing.T, pkgDirName string) {
	mockNbi, err := callCreateOrUpdatePackage(pkgDirName)

	if err != nil {
		t.Errorf("want: create package; got: %v", err)
	}
	if len(mockNbi.exchanges) != 3 { // #1 = get token
		t.Fatalf("want: one req to lookup package, then one to create it; got: %d",
			len(mockNbi.exchanges)-1)
	}

	rr := mockNbi.exchanges[2]
	checkUploadedPackage(t, mockNbi, rr.req, pkgDirName, pkgDirName)
	if rr.res.StatusCode != http.StatusCreated {
		t.Errorf("want status: %d; got: %d",
			http.StatusCreated, rr.res.StatusCode)
	}
}

func runUpdatePackageTest(t *testing.T, pkgDirName, descFilePath, osmPkgId string) {
	mockNbi, err := callCreateOrUpdatePackage(pkgDirName)

	if err != nil {
		t.Errorf("want: update package; got: %v", err)
	}
	if len(mockNbi.exchanges) != 3 { // #1 = get token
		t.Fatalf("want: one req to lookup package, then one to update it; got: %d",
			len(mockNbi.exchanges)-1)
	}

	updateExchange := mockNbi.exchanges[2]
	checkUploadedPackageDesc(t, mockNbi, pkgDirName, descFilePath, osmPkgId)
	if updateExchange.res.StatusCode != http.StatusOK {
		t.Errorf("want update status: %d; got: %d",
			http.StatusOK, updateExchange.res.StatusCode)
	}
}

func runUpdatePackageTestNoDescErr(t *testing.T, pkgDirName string) {
	_, err := callCreateOrUpdatePackage(pkgDirName)
	if err == nil {
		t.Errorf("want: update package error; got: nil")
	}
	if !strings.HasPrefix(err.Error(), "no descriptor found") {
		t.Errorf("want: no desc error; got: %v", err)
	}
}

func runUpdatePackageTestManyDescErr(t *testing.T, pkgDirName string) {
	_, err := callCreateOrUpdatePackage(pkgDirName)
	if err == nil {
		t.Errorf("want: update package error; got: nil")
	}
	if !strings.HasPrefix(err.Error(), "found more than one potential descriptor") {
		t.Errorf("want: many desc error; got: %v", err)
	}
}

func TestCreateKnfPackage(t *testing.T) {
	runCreatePackageTest(t, "create_knf")
}

func TestCreateNsPackage(t *testing.T) {
	runCreatePackageTest(t, "create_ns")
}

func TestUpdateKnfPackage(t *testing.T) {
	osmPkgId := "4ffdeb67-92e7-46fa-9fa2-331a4d674137" // see vnfDescriptors
	runUpdatePackageTest(t, "openldap_knf", "openldap_vnfd.yaml", osmPkgId)
}

func TestUpdateNsPackage(t *testing.T) {
	osmPkgId := "aba58e40-d65f-4f4e-be0a-e248c14d3e03" // see nsDescriptors
	runUpdatePackageTest(t, "openldap_ns", "openldap_nsd.yaml", osmPkgId)
}

func TestUpdateKnfPackageNoDescErr(t *testing.T) {
	runUpdatePackageTestNoDescErr(t, "update_no_desc/openldap_knf")
}

func TestUpdateNsPackageNoDescErr(t *testing.T) {
	runUpdatePackageTestNoDescErr(t, "update_no_desc/openldap_ns")
}

func TestUpdateKnfPackageManyDescErr(t *testing.T) {
	runUpdatePackageTestManyDescErr(t, "update_many_desc/openldap_knf")
}

func TestUpdateNsPackageManyDescErr(t *testing.T) {
	runUpdatePackageTestManyDescErr(t, "update_many_desc/openldap_ns")
}

func TestPackErrOnSourceDirAccess(t *testing.T) {
	mockNbi, err := callCreateOrUpdatePackage("not-there_knf")

	if _, ok := err.(*file.VisitError); !ok {
		t.Errorf("want: visit error; got: %v", err)
	}
	if len(mockNbi.exchanges) > 0 {
		t.Errorf("want: no req to create or update package; got: %d",
			len(mockNbi.exchanges))
	}
}

func TestCreateUnsupportedPackage(t *testing.T) {
	mockNbi, err := callCreateOrUpdatePackage("unsupported")
	if len(mockNbi.exchanges) > 0 {
		t.Errorf("want: no req to create or update package; got: %d",
			len(mockNbi.exchanges))
	}
	checkUnsupportedPackageErr(t, err)
}
