package bytez

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

const dataSize = 4 * 1024

func makeData() []byte {
	data := make([]byte, dataSize)
	for k := 0; k < dataSize; k++ {
		data[k] = byte(k % 256)
	}
	return data
}

func checkData(t *testing.T, data []byte) {
	if len(data) != dataSize {
		t.Errorf("want size: %d; got: %d", dataSize, len(data))
	}
	for k := 0; k < len(data); k++ {
		want := byte(k % 256)
		if data[k] != want {
			t.Errorf("[%d] want: %d; got: %d", k, want, data[k])
		}
	}
}

func writeAll(dest io.WriteCloser) {
	defer dest.Close()
	src := bytes.NewBuffer(makeData())
	io.Copy(dest, src)
}

func readAll(src io.ReadCloser) []byte {
	defer src.Close()
	data, _ := io.ReadAll(src)
	return data
}

func TestWriteThenRead(t *testing.T) {
	buf := NewBuffer()
	writeAll(buf)
	got := readAll(buf)
	checkData(t, got)
}

func TestBytesBeforeAnyRead(t *testing.T) {
	buf := NewBuffer()
	buf.Write([]byte{1, 2, 3})

	got := buf.Bytes()
	want := []byte{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v; got: %v", want, got)
	}

	got, _ = io.ReadAll(buf)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v; got: %v", want, got)
	}
}

func TestBytesAfterRead(t *testing.T) {
	buf := NewBuffer()
	buf.Write([]byte{1, 2, 3})

	firstTwo := make([]byte, 2)
	buf.Read(firstTwo)

	want := []byte{1, 2}
	if !reflect.DeepEqual(firstTwo, want) {
		t.Errorf("want: %v; got: %v", want, firstTwo)
	}

	got := buf.Bytes()
	want = []byte{3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v; got: %v", want, got)
	}
}
