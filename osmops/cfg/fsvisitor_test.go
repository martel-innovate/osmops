package cfg

import (
	"fmt"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"github.com/martel-innovate/osmops/osmops/util/file"
)

type processor struct {
	received []*KduNsActionFile
}

func (p *processor) Process(file *KduNsActionFile) error {
	p.received = append(p.received, file)
	if file.Content.Kdu.Name == "k3" {
		return fmt.Errorf("k3")
	}
	return nil
}

func buildScanner(t *testing.T) *KduNsActionRepoScanner {
	var err error
	repoRootDir := findTestDataDir(6)

	store, err := NewStore(repoRootDir)
	if err != nil {
		t.Fatalf("want: new store; got: %v", err)
	}

	return NewKduNsActionRepoScanner(store)
}

func TestVisit(t *testing.T) {
	scanner := buildScanner(t)
	visitor := &processor{}
	errors := scanner.Visit(visitor)

	errorFileNames := []string{}
	for _, e := range errors {
		if ve, ok := e.(*file.VisitError); ok {
			name := filepath.Base(ve.AbsPath)
			errorFileNames = append(errorFileNames, name)
		}
	}
	sort.Strings(errorFileNames)
	wantErrorFileNames := []string{"k1.ops.yaml", "k3.ops.yaml"}
	if !reflect.DeepEqual(wantErrorFileNames, errorFileNames) {
		t.Errorf("want error files: %s; got: %s",
			wantErrorFileNames, errorFileNames)
	}

	visited := []string{}
	for _, r := range visitor.received {
		visited = append(visited, r.Content.Kdu.Name)
	}
	sort.Strings(visited)
	wantVisited := []string{"k2", "k3"}
	if !reflect.DeepEqual(wantVisited, visited) {
		t.Errorf("want visited: %s; got: %s", wantVisited, visited)
	}
}

func TestVisitFileIOReadError(t *testing.T) {
	scanner := buildScanner(t)
	scanner.readFile = func(path string) ([]byte, error) {
		return nil, fmt.Errorf("can't read file: %v", path)
	}
	visitor := &processor{}

	errors := scanner.Visit(visitor)
	if len(errors) != 3 {
		t.Errorf("want: IO err on k1, k2 and k3 paths; got: %v", errors)
	}
	if len(visitor.received) != 0 {
		t.Errorf("want: no ops files visited; got: %v", visitor.received)
	}
}
