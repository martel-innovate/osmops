package tgz

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

// EntryReader processes an entry in a tar archive.
// The entry is at archivePath and has an associated file metadata whereas
// the content should only be read if the entry is a regular file.
type EntryReader func(
	archivePath string, fi os.FileInfo, content io.Reader) error

// Reader calls an EntryReader on each entry in a tar archive.
type Reader interface {
	// IterateEntries calls process on each tar entry.
	// Regardless of errors, IterateEntries closes the archive stream,
	// making the Reader unusable.
	IterateEntries(process EntryReader) error
	// Close releases all archive stream resources, making the Reader
	// unusable. Subsequent calls have no effect.
	Close()
}

type rdr struct {
	source        io.ReadCloser
	deflateStream *gzip.Reader
	archive       *tar.Reader
	closed        bool
}

// NewReader creates a Reader to process entries contained in the given
// gzip-compressed tar archive.
func NewReader(source io.ReadCloser) (Reader, error) {
	if source == nil {
		return nil, fmt.Errorf("nil source")
	}
	deflateStream, err := gzip.NewReader(source)
	if err != nil {
		return nil, err
	}
	return &rdr{
		source:        source,
		deflateStream: deflateStream,
		archive:       tar.NewReader(deflateStream),
		closed:        false,
	}, nil
}

func (r *rdr) Close() {
	if r.closed {
		return
	}
	r.source.Close()
	r.deflateStream.Close()
	r.closed = true
}

func (r *rdr) IterateEntries(process EntryReader) error {
	defer r.Close()

	if process == nil {
		return fmt.Errorf("nil entry reader")
	}
	if r.closed {
		return fmt.Errorf("closed reader")
	}

	return r.forEachEntry(process)
}

func (r *rdr) forEachEntry(process EntryReader) error {
	for {
		header, err := r.archive.Next()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = process(header.Name, header.FileInfo(), r.archive)
		if err != nil {
			return err
		}
	}
}
