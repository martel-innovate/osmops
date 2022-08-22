package bytez

import (
	"bytes"
)

// Buffer wraps a bytes.Buffer to implement io.ReadCloser and io.WriteCloser.
type Buffer struct {
	data *bytes.Buffer
}

// NewBuffer returns a new empty memory buffer.
func NewBuffer() *Buffer {
	return &Buffer{data: new(bytes.Buffer)}
}

// Write implements io.Writer.
func (buf *Buffer) Write(p []byte) (n int, err error) {
	return buf.data.Write(p)
}

// Read implements io.Reader.
func (buf *Buffer) Read(p []byte) (n int, err error) {
	return buf.data.Read(p)
}

// Close implements io.Closer.
func (buf *Buffer) Close() error {
	return nil
}

// Bytes returns a slice holding the unread portion of the buffer.
func (buf *Buffer) Bytes() []byte {
	return buf.data.Bytes()
}
