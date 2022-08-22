package cfg

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	u "github.com/martel-innovate/osmops/osmops/util"
	"github.com/martel-innovate/osmops/osmops/util/file"
)

func findTestDataDir(dirIndex int) file.AbsPath {
	_, thisFileName, _, _ := runtime.Caller(1)
	enclosingDir := filepath.Dir(thisFileName)
	testDataDirName := fmt.Sprintf("test_%d", dirIndex)
	testDataDir := filepath.Join(enclosingDir, "store_test_dir",
		testDataDirName)
	p, _ := file.ParseAbsPath(testDataDir)

	return p
}

func TestInstantiateStoreWithFullConfig(t *testing.T) {
	var err error
	repoRootDir := findTestDataDir(1)
	s, err := NewStore(repoRootDir)

	if err != nil {
		t.Fatalf("want: new store; got: %v", err)
	}

	if !reflect.DeepEqual(repoRootDir, s.RepoRootDirectory()) {
		t.Errorf("want: %v; got: %v", repoRootDir, s.RepoRootDirectory())
	}

	wantTargetDir := repoRootDir.Join("deploy.me")
	if !reflect.DeepEqual(wantTargetDir, s.RepoTargetDirectory()) {
		t.Errorf("want: %v; got: %v", wantTargetDir, s.RepoTargetDirectory())
	}

	ext, _ := u.NewNonEmptyStr(".ops.yaml")
	wantExts := []u.NonEmptyStr{ext}
	if !reflect.DeepEqual(wantExts, s.OpsFileExtensions()) {
		t.Errorf("want: %v; got: %v", wantExts, s.OpsFileExtensions())
	}

	wantCreds := &OsmConnection{
		Hostname: "host.ie:8008", Project: "boetie", User: "vans", Password: "*",
	}
	if !reflect.DeepEqual(wantCreds, s.OsmConnection()) {
		t.Errorf("want: %v; got: %v", wantCreds, s.OsmConnection())
	}

}

func TestInvalidRepoRootDir(t *testing.T) {
	repoRootDir := findTestDataDir(0)
	if s, err := NewStore(repoRootDir); err == nil {
		t.Errorf("want: no store on invalid root dir; got: %v", s)
	}
}

func TestNoConfigFile(t *testing.T) {
	repoRootDir := findTestDataDir(2)
	if s, err := NewStore(repoRootDir); err == nil {
		t.Errorf("want: no store if no config file found; got: %v", s)
	}
}

func TestNoTargetDirOnFS(t *testing.T) {
	repoRootDir := findTestDataDir(3)
	if s, err := NewStore(repoRootDir); err == nil {
		t.Errorf("want: no store if no target dir exists; got: %v", s)
	}
}

func TestNoConnectionFileOnFS(t *testing.T) {
	repoRootDir := findTestDataDir(4)
	if s, err := NewStore(repoRootDir); err == nil {
		t.Errorf("want: no store if no connection file exists; got: %v", s)
	}
}

func TestDefaultFileExtensions(t *testing.T) {
	repoRootDir := findTestDataDir(5)
	if s, err := NewStore(repoRootDir); err != nil {
		t.Fatalf("want: new store; got: %v", err)
	} else {
		wantExts := DefaultOpsFileExtensions()
		if !reflect.DeepEqual(wantExts, s.OpsFileExtensions()) {
			t.Errorf("want: %v; got: %v", wantExts, s.OpsFileExtensions())
		}
	}
}

func TestNoRepoPkgRootDir(t *testing.T) {
	repoRootDir := findTestDataDir(1)
	store, _ := NewStore(repoRootDir)

	got, err := store.RepoPkgDirectories()
	if err != nil {
		t.Fatalf("want: empty list; got error: %v", err)
	}
	if got == nil || len(got) != 0 {
		t.Errorf("want: empty list; got: %v", got)
	}
}

func TestRepoPkgRootDirWithNoSubdirs(t *testing.T) {
	repoRootDir := findTestDataDir(5)
	store, _ := NewStore(repoRootDir)

	got, err := store.RepoPkgDirectories()
	if err != nil {
		t.Fatalf("want: empty list; got error: %v", err)
	}
	if got == nil || len(got) != 0 {
		t.Errorf("want: empty list; got: %v", got)
	}
}

func TestRepoPkgRootDirWithSubdirs(t *testing.T) {
	repoRootDir := findTestDataDir(6)
	store, _ := NewStore(repoRootDir)
	want := []file.AbsPath{
		repoRootDir.Join("deploy.me/osm-pkgs/p1"),
		repoRootDir.Join("deploy.me/osm-pkgs/p2"),
	}

	got, err := store.RepoPkgDirectories()
	if err != nil {
		t.Fatalf("want: %v; got error: %v", want, err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v; got: %v", want, got)
	}
}
