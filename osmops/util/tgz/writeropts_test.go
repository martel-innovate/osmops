package tgz

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"path"
	"testing"
	"time"
)

func checkCompressionLevel(t *testing.T, cfg *writerOpts, wantLevel int) {
	if cfg.compressionLevel != wantLevel {
		t.Errorf("want comp: %d; got: %d", wantLevel, cfg.compressionLevel)
	}
}

func checkBaseHdrFields(t *testing.T, cfg *writerOpts) *tar.Header {
	hdr := &tar.Header{}
	cfg.setHeaderFields("some/file", hdr)

	wantName := path.Join(cfg.baseDirName, "some/file")
	if hdr.Name != wantName {
		t.Errorf("want name: %s; got: %s", wantName, hdr.Name)
	}
	if hdr.Typeflag != tar.TypeReg {
		t.Errorf("want type reg; got: %v", hdr.Typeflag)
	}
	if hdr.Format != tar.FormatPAX {
		t.Errorf("want pax format; got: %v", hdr.Format)
	}

	return hdr
}

func TestBaseWriterCfg(t *testing.T) {
	got := makeWriterCfg("baseDir")
	checkCompressionLevel(t, got, gzip.BestCompression)
	checkBaseHdrFields(t, got)
}

func TestDefaultCompressionOpt(t *testing.T) {
	got := makeWriterCfg("baseDir", WithDefaultCompression())
	checkCompressionLevel(t, got, gzip.DefaultCompression)
	checkBaseHdrFields(t, got)
}

func TestBestCompressionOpt(t *testing.T) {
	got := makeWriterCfg("baseDir", WithBestCompression())
	checkCompressionLevel(t, got, gzip.BestCompression)
	checkBaseHdrFields(t, got)
}

func TestSpeedCompressionOpt(t *testing.T) {
	got := makeWriterCfg("baseDir", WithBestSpeed())
	checkCompressionLevel(t, got, gzip.BestSpeed)
	checkBaseHdrFields(t, got)
}

func TestEntryTimeOpt(t *testing.T) {
	epochStart := time.Unix(0, 0)
	cfg := makeWriterCfg("baseDir", WithEntryTime(epochStart))
	hdr := checkBaseHdrFields(t, cfg)

	if !hdr.AccessTime.Equal(epochStart) {
		t.Errorf("want access time: %v; got: %v", epochStart, hdr.AccessTime)
	}
	if !hdr.ChangeTime.Equal(epochStart) {
		t.Errorf("want change time: %v; got: %v", epochStart, hdr.ChangeTime)
	}
	if !hdr.ModTime.Equal(epochStart) {
		t.Errorf("want mod time: %v; got: %v", epochStart, hdr.ModTime)
	}
}

func TestMakeWriterCfgIgnoreNilSettings(t *testing.T) {
	got := makeWriterCfg("baseDir", nil)
	if got == nil {
		t.Fatalf("want: config; got: nil")
	}
	checkBaseHdrFields(t, got)
}

func withErrHdrSetter(err error) WriterOption {
	return func(opts *writerOpts) {
		opts.chainHdrSetter(func(archivePath string, hdr *tar.Header) error {
			return err
		})
	}
}

func withHdrSetter(set func(hdr *tar.Header)) WriterOption {
	return func(opts *writerOpts) {
		opts.chainHdrSetter(func(archivePath string, hdr *tar.Header) error {
			set(hdr)
			return nil
		})
	}
}

func TestChainSettersBailOutOnFirstErr(t *testing.T) {
	wantErr := fmt.Errorf("foo")
	dontWantHdrName := "this setter shouldn't be called"
	nameOverride := func(hdr *tar.Header) {
		hdr.Name = dontWantHdrName
	}
	cfg := makeWriterCfg("baseDir",
		withErrHdrSetter(wantErr), withHdrSetter(nameOverride))
	hdr := &tar.Header{}

	if err := cfg.setHeaderFields("some/file", hdr); err != wantErr {
		t.Errorf("want err: %v; got: %v", wantErr, err)
	}
	if hdr.Name == dontWantHdrName {
		t.Errorf(
			"want: name setter not called b/c of previous setter err; got: called")
	}
}
