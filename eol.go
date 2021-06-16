// Package eol implements end-of-line sequence detection.
// It only supports Windows (CRLF) and POSIX (LF) sequences.
package eol

import (
	"io"
)

const (
	WINDOWS = "\r\n" // Windows end-of-line sequence
	POSIX   = "\n"   // POSIX end-of-line sequence
)

// Detect takes an io.ByteReader and reads it until
// a known end-of-line sequence is met. If no EOLs
// were met, a default sequence for current GOOS is
// returned.
func Detect(r io.ByteReader) (eol string, err error) {
	eol = DEFAULT
	var currByte, prevByte byte
	for {
		currByte, err = r.ReadByte()
		if err == io.EOF {
			err = nil
			return
		} else if err != nil {
			return
		}

		if currByte == '\n' {
			if prevByte == '\r' {
				eol = WINDOWS
				return
			} else {
				eol = POSIX
				return
			}
		}

		prevByte = currByte
	}
}
