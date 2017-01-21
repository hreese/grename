package grename

import (
	"bufio"
	"bytes"
	"io"
)

// MakeSourceFromStrings returns a Source for []string-arrays like argv
func MakeSourceFromStrings(instrings ...string) Source {
	return func() <-chan Renamed {
		outC := make(chan Renamed)
		go func() {
			for _, instring := range instrings {
				outC <- Renamed{instring, instring}
			}
			close(outC)
		}()
		return outC
	}
}

// MakeSourceFromStrings returns a Source for strings read from an io.Reader
// (like os.Stdin). String separation is determined by splitFunc (bufio.Scanlines
// and ScanNUL).
func MakeSourceFromScanner(input io.Reader, splitFunc bufio.SplitFunc) Source {
	return func() <-chan Renamed {
		outC := make(chan Renamed)
		go func() {
			scanner := bufio.NewScanner(input)
			scanner.Split(splitFunc)
			for scanner.Scan() {
                token := scanner.Text()
				outC <- Renamed{token, token}
			}
			close(outC)
		}()
		return outC
	}
}

// ScanNUL tokenizes strings by splitting them at NUL bytes
func ScanNUL(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\x00'); i >= 0 {
		// complete NUL terminated string
		return i + 1, dropNUL(data[0:i]), nil
	}
	if atEOF {
		return len(data), dropNUL(data), nil
	}
	// Request more data
	return 0, nil, nil
}

// dropNUL removes a terminala NUL byte from data
func dropNUL(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\x00' {
		return data[0 : len(data)-1]
	}
	return data
}
