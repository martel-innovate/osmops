package engine

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/martel-innovate/osmops/osmops/cfg"
	"github.com/martel-innovate/osmops/osmops/util/file"
)

func findTestDataDir(dirIndex int) file.AbsPath {
	_, thisFileName, _, _ := runtime.Caller(1)
	enclosingDir := filepath.Dir(thisFileName)
	testDataDirName := fmt.Sprintf("test_%d", dirIndex)
	testDataDir := filepath.Join(enclosingDir, "reconcile_test_dir",
		testDataDirName)
	p, _ := file.ParseAbsPath(testDataDir)

	return p
}

func TestReconcileFailOnInvalidRootDir(t *testing.T) {
	logger := newLogCollector()

	if _, err := New(newCtx(logger), ""); err == nil {
		t.Errorf("want: error; got: nil")
	}

	if got := logger.countEntries(); got != 1 {
		t.Fatalf("want: 1; got: %d", got)
	}
	if got := logger.msgAt(0); got != engineInitErrMsg {
		t.Errorf("want: %s; got: %s", engineInitErrMsg, got)
	}
}

func TestReconcileFailOnInvalidOpsConfig(t *testing.T) {
	logger := newLogCollector()
	repoRootDir := findTestDataDir(1)

	if _, err := New(newCtx(logger), repoRootDir.Value()); err == nil {
		t.Errorf("want: error; got: nil")
	}

	if got := logger.countEntries(); got != 1 {
		t.Fatalf("want: 1; got: %d", got)
	}
	if got := logger.msgAt(0); got != engineInitErrMsg {
		t.Errorf("want: %s; got: %s", engineInitErrMsg, got)
	}
	if got, ok := logger.errAt(0).(*fs.PathError); !ok {
		t.Errorf("want: path error; got: %v", got)
	}
}

func TestReconcileDoNothingIfNoOsmGitOpsFileNorOsmPkgFound(t *testing.T) {
	logger := newLogCollector()
	repoRootDir := findTestDataDir(2)

	engine, err := New(newCtx(logger), repoRootDir.Value())
	if err != nil {
		t.Errorf("want: engine; got: %v", err)
	}

	engine.Reconcile()
	if got := logger.countEntries(); got != 0 {
		t.Fatalf("want: 0; got: %d", got)
	}
}

func TestReconcileProcessOsmGitOpsFiles(t *testing.T) {
	logger := newLogCollector()
	repoRootDir := findTestDataDir(3)
	mockNbic := newMockNbicWorkflow()

	engine, err := New(newCtx(logger), repoRootDir.Value())
	if err != nil {
		t.Errorf("want: engine; got: %v", err)
	}

	engine.nbic = mockNbic
	engine.Reconcile()

	if mockNbic.hasProcessedKdu("k1") {
		t.Errorf("want: skip k1 (invalid content); got: processed")
	}

	if data := mockNbic.dataFor("k2"); data == nil {
		t.Errorf("want: process k2; got: not processed")
	} else {
		if data.KduParams != nil {
			t.Errorf("want: nil; got: %+v", data.KduParams)
		}
	}

	if !mockNbic.hasProcessedKdu("k3") {
		t.Errorf("want: process k3; got: not processed")
	} else {
		got := mockNbic.lookupParam("k3", "replicaCount")
		if got != "3" {
			t.Errorf(`want: "3"; got: %s`, got)
		}
	}

	if got := logger.countEntries(); got != 4 {
		t.Errorf("want: 4; got: %d", got)
	}
	want := []string{"k2.ops.yaml", "k3.ops.yaml"}
	if got := logger.sortProcessedFileNames(); !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v; got: %v", want, got)
	}
	want = []string{"k1.ops.yaml", "k2.ops.yaml"}
	// k2: simulated processing error, see mockCreateOrUpdate
	if got := logger.sortErrorFileNames(); !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

func TestNewNbicFailOnInvalidHostAndPort(t *testing.T) {
	config := &cfg.OsmConnection{
		Hostname: "missing.port", // never happens b/c of yaml validation
		User:     "u",
		Password: "*",
		Project:  "p",
	}
	if _, err := newNbic(config); err == nil {
		t.Errorf("want: error; got: nil")
	}
}

func TestReconcileProcessNoOsmGitOpsFilesOnPackageErr(t *testing.T) {
	logger := newLogCollector()
	repoRootDir := findTestDataDir(4)
	mockNbic := newMockNbicWorkflow()
	engine, _ := New(newCtx(logger), repoRootDir.Value())
	engine.nbic = mockNbic

	engine.Reconcile()

	wantProcessedPkgs := []string{"p2"}
	if !reflect.DeepEqual(mockNbic.processedPkgNames, wantProcessedPkgs) {
		t.Errorf("want processed pkgs: %v; got: %v", wantProcessedPkgs,
			mockNbic.processedPkgNames)
	}
	if mockNbic.hasProcessedKdus() {
		t.Errorf("want: skip kdus b/c of prev pkg errors; got: some processed")
	}
}

func TestReconcileProcessStopOnRootPackageDirAccessErr(t *testing.T) {
	logger := newLogCollector()
	repoRootDir := findTestDataDir(5)
	mockNbic := newMockNbicWorkflow()
	engine, _ := New(newCtx(logger), repoRootDir.Value())
	engine.nbic = mockNbic

	pkgsDir := repoRootDir.Join("deploy.me/osm-pkgs")
	os.Chmod(pkgsDir.Value(), 0200) // processPackages can't scan it
	defer os.Chmod(pkgsDir.Value(), 0755)

	engine.Reconcile()

	wantProcessedPkgs := []string{}
	if !reflect.DeepEqual(mockNbic.processedPkgNames, wantProcessedPkgs) {
		t.Errorf("want processed pkgs: %v; got: %v", wantProcessedPkgs,
			mockNbic.processedPkgNames)
	}
	if mockNbic.hasProcessedKdus() {
		t.Errorf("want: skip kdus b/c of prev pkg errors; got: some processed")
	}
}

func TestReconcileProcessPackagesAndOsmGitOpsFiles(t *testing.T) {
	logger := newLogCollector()
	repoRootDir := findTestDataDir(5)
	mockNbic := newMockNbicWorkflow()
	engine, _ := New(newCtx(logger), repoRootDir.Value())
	engine.nbic = mockNbic

	engine.Reconcile()

	wantProcessedPkgs := []string{"p2", "p3"}
	if !reflect.DeepEqual(mockNbic.processedPkgNames, wantProcessedPkgs) {
		t.Errorf("want processed pkgs: %v; got: %v", wantProcessedPkgs,
			mockNbic.processedPkgNames)
	}

	if mockNbic.hasProcessedKdu("k1") {
		t.Errorf("want: skip k1 (invalid content); got: processed")
	}
	if !mockNbic.hasProcessedKdu("k2") {
		t.Errorf("want: process k2; got: not processed")
	}
	if !mockNbic.hasProcessedKdu("k3") {
		t.Errorf("want: process k3; got: not processed")
	}
}
