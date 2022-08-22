package pkgr

import (
	"io/fs"
	"io/ioutil"
	"os"
	"testing"

	"github.com/martel-innovate/osmops/osmops/util/file"
)

var md5stringFixtures = []struct {
	input []byte
	want  string // go's md5 should be the same as md5sum
}{
	// $ touch empty; md5sum empty
	{nil, "d41d8cd98f00b204e9800998ecf8427e"},
	{[]byte{}, "d41d8cd98f00b204e9800998ecf8427e"},
	// $ echo -n 1 | md5sum
	{[]byte{49}, "c4ca4238a0b923820dcc509a6f75849b"},
	// $ echo -n 12 | md5sum
	{[]byte{49, 50}, "c20ad4d76fe97759aa27a0c99bff6710"},
}

func TestMd5string(t *testing.T) {
	for k, d := range md5stringFixtures {
		got := md5string(d.input)
		if got != d.want {
			t.Errorf("[%d] want: %s; got: %s", k, d.want, got)
		}
	}
}

func TestComputeChecksumFileAccessErr(t *testing.T) {
	fd, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("couldn't create temp file: %v", err)
	}
	defer os.Remove(fd.Name())
	fd.Close()
	os.Chmod(fd.Name(), 0200) // computeChecksum can't read it

	filePath, err := file.ParseAbsPath(fd.Name())
	if err != nil {
		t.Fatalf("couldn't create temp file: %v", err)
	}
	_, err = computeChecksum(filePath)
	if _, ok := err.(*fs.PathError); !ok {
		t.Errorf("want: path error; got: %v", err)
	}
}
