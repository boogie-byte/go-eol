package eol

import (
	"bytes"
	"errors"
	"testing"
)

type errByteReader struct{}

func (*errByteReader) ReadByte() (byte, error) {
	return 0, errors.New("test error")
}

func runTest(t *testing.T, s, exp string) {
	t.Helper()

	b := bytes.NewBufferString(s)
	eol, err := Detect(b)
	if err != nil {
		t.Fatal(err)
	}

	if eol != exp {
		t.Fatalf("EOL should be '%s', but got '%s' instead", exp, eol)
	}
}

func TestNoEOL(t *testing.T) {
	runTest(t, "foo", DEFAULT)
}

func TestEmpty(t *testing.T) {
	runTest(t, "", DEFAULT)
}

func TestCRLF(t *testing.T) {
	runTest(t, "foo\r\nbar", WINDOWS)
}

func TestLF(t *testing.T) {
	runTest(t, "foo\nbar", POSIX)
}

func TestLFCR(t *testing.T) {
	runTest(t, "foo\n\rbar", POSIX)
}

func TestCR(t *testing.T) {
	runTest(t, "foo\rbar", DEFAULT)
}

func TestError(t *testing.T) {
	b := &errByteReader{}
	eol, err := Detect(b)
	if err == nil {
		t.Fatal("error should not be nil")
	}

	if eol != DEFAULT {
		t.Fatalf("EOL should be '%s', but got '%s' instead", DEFAULT, eol)
	}
}
