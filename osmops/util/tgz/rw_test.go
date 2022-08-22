package tgz

import (
	"fmt"
	"io"
	"os"
	"path"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/martel-innovate/osmops/osmops/util/bytez"
	"github.com/martel-innovate/osmops/osmops/util/file"
)

func writeArchive(sink io.WriteCloser) []error {
	errs := []error{}

	writer, err := NewWriter("", sink, WithBestSpeed())
	if err != nil {
		errs = append(errs, err)
		return errs
	}
	defer writer.Close()

	sourceDir := findTestDataDir()
	scanner := file.NewTreeScanner(sourceDir)
	es := scanner.Visit(writer.Visitor())
	errs = append(errs, es...)

	if err := writer.AddEntry("extra", strings.NewReader("extra")); err != nil {
		errs = append(errs, err)
	}

	return errs
}

func writeArchiveAndCreateReader(t *testing.T) Reader {
	buf := bytez.NewBuffer()
	es := writeArchive(buf)
	if len(es) > 0 {
		t.Fatalf("couldn't write archive: %v", es)
	}

	reader, err := NewReader(buf)
	if err != nil {
		t.Fatalf("couldn't create reader: %v", err)
	}

	return reader
}

func checkEntryContent(archivePath string, fi os.FileInfo, entry io.Reader) error {
	name := path.Base(archivePath)
	contentBytes, err := io.ReadAll(entry)
	if err != nil {
		return err
	}
	text := string(contentBytes)
	if name != text {
		return fmt.Errorf("path = %s; content = %s", archivePath, text)
	}
	return nil
}

func checkArchivePaths(t *testing.T, got []string) {
	want := []string{
		"d1/f2", "d1/f3", "d2/d3/f6", "d2/f4", "d2/f5", "extra", "f1",
	}

	sort.Strings(want)
	sort.Strings(got)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

func TestWriteThenReadContent(t *testing.T) {
	reader := writeArchiveAndCreateReader(t)
	if err := reader.IterateEntries(checkEntryContent); err != nil {
		t.Errorf("entry content should be the same as entry name: %v", err)
	}
}

func TestWriteThenReadPaths(t *testing.T) {
	reader := writeArchiveAndCreateReader(t)

	paths := []string{}
	_ = reader.IterateEntries(
		func(archivePath string, fi os.FileInfo, entry io.Reader) error {
			paths = append(paths, archivePath)
			return nil
		})

	checkArchivePaths(t, paths)
}
