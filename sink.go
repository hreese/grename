package grename

import (
    "io"
)

func MakeSinkFromWriter(out io.Writer, sepNames string, sepPairs string) Sink {
    return func(inC <-chan Renamed) {
        go func() {
            for s := range inC {
                out.Write([]byte(s.original))
                out.Write([]byte(sepNames))
                out.Write([]byte(s.renamed))
                out.Write([]byte(sepPairs))
            }
        }()
    }
}
