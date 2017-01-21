package grename

import (
	"io"
)

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
