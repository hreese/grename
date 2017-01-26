package grename

import (
	"io"
	"os"
)

// MakeSinkToWriter returns a Sink that writes to io.Writer out with strings
// separated by sepNames and pairs separated by sepPairs.
func MakeSinkToWriter(out io.Writer, sepNames string, sepPairs string) Sink {
	return func(inC <-chan Renamed) {
		for s := range inC {
			out.Write([]byte(s.Original))
			out.Write([]byte(sepNames))
			out.Write([]byte(s.Renamed))
			out.Write([]byte(sepPairs))
		}
	}
}

var (
	StdoutSink = MakeSinkToWriter(os.Stdout, " → ", "\n")
	NULSink    = MakeSinkToWriter(os.Stdout, "\x00", "\x00")
)
