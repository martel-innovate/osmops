package file

import (
	"io/fs"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
)

// TODO. The path tests will probably fail on Windows since we're using
// Unix paths. We could use filepath.Join to make most of them platform
// independent but I'm not sure how to make absolute paths though...

var invalidPathFixtures = []string{"", " ", "\n", "\t "}

func TestInvalidPath(t *testing.T) {
	for k, d := range invalidPathFixtures {
		if err := IsStringPath(d); err == nil {
			t.Errorf("[%d] want: invalid; got: valid", k)
		}
	}
}

var parsePathFixtures = []struct {
	in   string
	want string
	rel  bool
}{
	{"/a/b/s", "/a/b/s", false}, {"r/e/l", "/r/e/l", true},
}

func TestParsePath(t *testing.T) {
	for k, d := range parsePathFixtures {
		if p, err := ParseAbsPath(d.in); err != nil {
			t.Errorf("[%d] want: valid parse; got: %v", k, err)
		} else {
			if !d.rel && d.want != p.Value() {
				t.Errorf("[%d] want: %s; got: %s", k, d.want, p.Value())
			}
			if d.rel && !strings.HasSuffix(p.Value(), d.want) {
				t.Errorf("[%d] want suffix: %s; got: %s", k, d.want, p.Value())
			}
		}
	}
}

var joinPathFixtures = []struct {
	base string
	rel  string
	want string
}{
	{"/a", "", "/a"}, {"/a/", " ", "/a"}, {"/a", "\t", "/a"},
	{"/a/", "b ", "/a/b"}, {"/a", "b\n", "/a/b"}, {"/a/b", "//c", "/a/b/c"},
}

func TestJoinPath(t *testing.T) {
	for k, d := range joinPathFixtures {
		if base, err := ParseAbsPath(d.base); err != nil {
			t.Errorf("[%d] want: valid parse; got: %v", k, err)
		} else {
			joined := base.Join(d.rel)
			if joined.Value() != d.want {
				t.Errorf("[%d] want: %s; got: %s", k, d.want, joined)
			}
		}
	}
}

func TestIsDir(t *testing.T) {
	if pwd, err := ParseAbsPath("."); err != nil {
		t.Errorf("want: valid parse; got: %v", err)
	} else {
		if err := pwd.IsDir(); err != nil {
			t.Errorf("want: pwd is a directory; got: %v", err)
		}

		notThere := pwd.Join("notThere")
		if err := notThere.IsDir(); err == nil {
			t.Errorf("want: not a directory; got directory: %v", notThere)
		}

		if tempFile, err := ioutil.TempFile("", "prefix"); err != nil {
			t.Errorf("couldn't create temp file: %v", err)
		} else {
			defer os.Remove(tempFile.Name())

			if tf, err := ParseAbsPath(tempFile.Name()); err != nil {
				t.Errorf("want: valid temp file parse; got: %v", err)
			} else {
				if err := tf.IsDir(); err == nil {
					t.Errorf("want: not a dir; got dir: %v", tf)
				}
			}
		}
	}
}

func assertListPaths(t *testing.T, dirIndex int, want []string) {
	got, err := ListPaths(findTestDataDir(dirIndex).Value())
	if len(err) != 0 {
		t.Fatalf("want: %v; got: %v", want, err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

func TestListPathsOfFlatDir(t *testing.T) {
	want := []string{"f1", "f2"}
	assertListPaths(t, 1, want)
}

func TestListPathsOfDirTree(t *testing.T) {
	want := []string{
		"d1", "d1/f2", "d1/f3",
		"d2", "d2/d3", "d2/d3/f6", "d2/f4", "d2/f5",
		"f1",
	}
	assertListPaths(t, 2, want)
}

func TestListPathsErrorOnInvalidTargetDir(t *testing.T) {
	got, err := ListPaths("")
	if err == nil {
		t.Errorf("want error; got: %v", got)
	}
}

func TestListSubDirectoryNamesWhenNoSubdirs(t *testing.T) {
	flatDir := findTestDataDir(1)
	got, err := ListSubDirectoryNames(flatDir.Value())
	if err != nil {
		t.Fatalf("want: empty list; got error: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("want: empty list; got: %v", got)
	}
}

func TestListSubDirectoryNamesWithDirTree(t *testing.T) {
	dirTree := findTestDataDir(2)
	want := []string{"d1", "d2"}

	got, err := ListSubDirectoryNames(dirTree.Value())
	if err != nil {
		t.Fatalf("want: %v; got error: %v", want, err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

func TestListSubDirectoryNamesScanDirErr(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "file-test")
	if err != nil {
		t.Fatalf("couldn't create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	os.Chmod(tempDir, 0200) // ListSubDirectoryNames can't scan it

	got, err := ListSubDirectoryNames(tempDir)
	if _, ok := err.(*fs.PathError); !ok {
		t.Errorf("want: path access error; got: %v", err)
	}
	if got == nil {
		t.Errorf("want: empty names list; got: nil")
	}
	if len(got) != 0 {
		t.Errorf("want: empty names list; got: %v", got)
	}
}
