package tgz

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/martel-innovate/osmops/osmops/util/bytez"
)

func TestNewReaderErrOnNilSource(t *testing.T) {
	if got, err := NewReader(nil); err == nil {
		t.Errorf("want error; got: %v", got)
	}
}

func TestNewReaderErrOnNonGzipSource(t *testing.T) {
	src := bytez.NewBuffer()
	src.Write([]byte{1, 2, 3})

	if got, err := NewReader(src); err == nil {
		t.Errorf("want error; got: %v", got)
	}
}

func TestIterateEntriesErrOnNilProcess(t *testing.T) {
	reader := writeArchiveAndCreateReader(t)
	if err := reader.IterateEntries(nil); err == nil {
		t.Errorf("want error; got: nil")
	}
}

func TestIterateEntriesErrOnClosedReader(t *testing.T) {
	reader := writeArchiveAndCreateReader(t)
	noop := func(archivePath string, fi os.FileInfo, content io.Reader) error {
		return nil
	}

	reader.Close()
	if err := reader.IterateEntries(noop); err == nil {
		t.Errorf("want error; got: nil")
	}
}

func TestIterateEntriesProcessErr(t *testing.T) {
	reader := writeArchiveAndCreateReader(t)
	bomb := func(archivePath string, fi os.FileInfo, content io.Reader) error {
		return fmt.Errorf("error at: %s", archivePath)
	}

	if err := reader.IterateEntries(bomb); err == nil {
		t.Errorf("want error; got: nil")
	}
}

func makeBrokenArchiveReader(t *testing.T) Reader {
	sink := bytez.NewBuffer()
	writer, err := NewWriter("", sink, WithBestSpeed())
	if err != nil {
		t.Fatalf("couldn't create writer: %v", err)
	}
	if err := writer.AddEntry("foo", strings.NewReader("bar")); err != nil {
		t.Fatalf("couldn't write entry: %v", err)
	}
	writer.Close()

	data, err := io.ReadAll(sink)
	if err != nil {
		t.Fatalf("couldn't read tgz data: %v", err)
	}
	data[50] = 50 // make reading of tar header fail

	source := bytez.NewBuffer()
	source.Write(data)

	reader, err := NewReader(source)
	if err != nil {
		t.Fatalf("couldn't create reader: %v", err)
	}
	return reader
}

func TestIterateEntriesErrOnBrokenArchive(t *testing.T) {
	reader := makeBrokenArchiveReader(t)
	noop := func(archivePath string, fi os.FileInfo, content io.Reader) error {
		return nil
	}

	if err := reader.IterateEntries(noop); err == nil {
		t.Errorf("want error; got: nil")
	}
}
