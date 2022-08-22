package tgz

import (
	"compress/gzip"
	"fmt"
	"io/fs"
	"strings"
	"testing"
	"time"

	"github.com/martel-innovate/osmops/osmops/util/bytez"
)

// implements io.Reader
type contentBomb struct{}

func (x *contentBomb) Read(p []byte) (n int, err error) {
	return 1, fmt.Errorf("failed to read")
}

// implements fs.FileInfo
type bogusFile struct {
	forceRegular  bool
	size          int64
	modeCallCount int
}

func (x *bogusFile) Name() string {
	return ""
}

func (x *bogusFile) Size() int64 {
	if x.forceRegular {
		return x.size
	}
	return 0
}

func (x *bogusFile) Mode() fs.FileMode {
	if x.forceRegular {
		return 0 // regular file
	}
	if x.modeCallCount == 0 { // AddFileEntry checks if it's regular
		x.modeCallCount++
		return 0 // regular file
	}
	return fs.ModeIrregular
}

func (x *bogusFile) ModTime() time.Time {
	return time.Now()
}

func (x *bogusFile) IsDir() bool {
	return false
}

func (x *bogusFile) Sys() interface{} {
	return nil
}

func makeMemWriter(opts ...WriterOption) (Writer, *bytez.Buffer) {
	sink := bytez.NewBuffer()
	if len(opts) == 0 {
		opts = []WriterOption{WithBestSpeed()}
	}
	writer, _ := NewWriter("", sink, opts...)

	return writer, sink
}

func TestNewWriterAcceptEmptyBaseDir(t *testing.T) {
	sink := bytez.NewBuffer()
	if _, err := NewWriter("", sink); err != nil {
		t.Errorf("want: writer; got: %v", err)
	}
}

func TestNewWriterErrOnNilSink(t *testing.T) {
	if got, err := NewWriter("", nil); err == nil {
		t.Errorf("want error; got: %v", got)
	}
}

func withInvalidCompLevel() WriterOption {
	return func(opts *writerOpts) {
		opts.compressionLevel = gzip.HuffmanOnly - 1
	}
}

func TestNewWriterErrOnInvalidGzipCompLevel(t *testing.T) {
	sink := bytez.NewBuffer()
	if got, err := NewWriter("", sink, withInvalidCompLevel()); err == nil {
		t.Errorf("want error; got: %v", got)
	}
}

func TestAddEntryErrOnContentRead(t *testing.T) {
	writer, _ := makeMemWriter()
	if err := writer.AddEntry("foo", &contentBomb{}); err == nil {
		t.Errorf("want error; got: nil")
	}
}

func TestAddFileDoNothingOnNilFileInfo(t *testing.T) {
	writer, _ := makeMemWriter()
	if err := writer.AddFile("", "foo", nil); err != nil {
		t.Errorf("want no error; got: %v", err)
	}
}

func TestAddFileErrOnBogusFileInfo(t *testing.T) {
	writer, _ := makeMemWriter()
	if err := writer.AddFile("", "foo", &bogusFile{}); err == nil {
		t.Errorf("want error; got: nil")
	}
}

func TestAddFileErrOnWriteHeader(t *testing.T) {
	writer, _ := makeMemWriter()
	fileInfo := &bogusFile{forceRegular: true, size: -1}

	err := writer.AddFile("", "", fileInfo)
	if err == nil || !strings.HasPrefix(err.Error(), "archive/tar:") { // (*)
		t.Errorf("want: tar header error; got: %v", err)
	}

	// NOTE. Error type. We expect tar.headerError but the tar pkg doesn't
	// export it, so the best we can do is check for the error message.
}

func TestAddFileErrOnFileOpen(t *testing.T) {
	writer, _ := makeMemWriter()
	fileInfo := &bogusFile{forceRegular: true}

	err := writer.AddFile("", "", fileInfo)
	if _, ok := err.(*fs.PathError); !ok {
		t.Errorf("want: path error; got: %v", err)
	}
}

func TestHeaderWritingErr(t *testing.T) {
	wantErr := fmt.Errorf("foo")
	writer, _ := makeMemWriter(WithBestSpeed(), withErrHdrSetter(wantErr))
	content := bytez.NewBuffer()
	content.Write([]byte{1})

	got := writer.AddEntry("foo", content)
	if got != wantErr {
		t.Errorf("want error: %v; got: %v", wantErr, got)
	}
}
