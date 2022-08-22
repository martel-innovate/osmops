package tgz

import (
	"archive/tar"
	"compress/gzip"
	"path"
	"time"
)

// tarHeaderSetter is a function to tweak tar headers the Writer puts in
// the archive. For each archive entry, you can make Writer call this
// function (or a list of them) to modify or add fields to the entry header.
// Setters get chained (in a Either monad of sorts) and there's always an
// initial setter to specify the base header for each file---see below.
type tarHeaderSetter func(archivePath string, hdr *tar.Header) error

type writerOpts struct {
	baseDirName      string
	compressionLevel int
	setHeaderFields  tarHeaderSetter
}

func (opts *writerOpts) chainHdrSetter(set tarHeaderSetter) {
	fst, snd := opts.setHeaderFields, set
	opts.setHeaderFields = func(archivePath string, hdr *tar.Header) error {
		if err := fst(archivePath, hdr); err != nil {
			return err
		}
		return snd(archivePath, hdr)
	}
}

type WriterOption func(opts *writerOpts)

func baseWriterOpts(baseDirName string) *writerOpts {
	return &writerOpts{
		baseDirName:      baseDirName,
		compressionLevel: gzip.BestCompression,
		setHeaderFields: func(archivePath string, hdr *tar.Header) error {
			hdr.Name = path.Join(baseDirName, archivePath)
			hdr.Typeflag = tar.TypeReg
			hdr.Format = tar.FormatPAX
			return nil
		},
	}
}

func makeWriterCfg(baseDirName string, opts ...WriterOption) *writerOpts {
	cfg := baseWriterOpts(baseDirName)
	for _, setting := range opts {
		if setting != nil {
			setting(cfg)
		}
	}
	return cfg
}

// Use gzip's default compression level when writing the archive.
func WithDefaultCompression() WriterOption {
	return func(opts *writerOpts) {
		opts.compressionLevel = gzip.DefaultCompression
	}
}

// Use the highest gzip compression level when writing the archive.
func WithBestCompression() WriterOption {
	return func(opts *writerOpts) {
		opts.compressionLevel = gzip.BestCompression
	}
}

// Use the lowest gzip compression level when writing the archive.
func WithBestSpeed() WriterOption {
	return func(opts *writerOpts) {
		opts.compressionLevel = gzip.BestSpeed
	}
}

// Set the access, change and mod time of each tar header to the specified
// time point.
func WithEntryTime(when time.Time) WriterOption {
	return func(opts *writerOpts) {
		opts.chainHdrSetter(func(archivePath string, hdr *tar.Header) error {
			hdr.AccessTime = when
			hdr.ChangeTime = when
			hdr.ModTime = when
			return nil
		})
	}
}
