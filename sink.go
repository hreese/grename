package grename

import (
	"io"
	"os"
)

type FileRenameOp struct {
	OldName string
	NewName string
}

type Sink struct {
	Input chan FileRenameOp
	Done  chan bool
}

// MakeSinkToWriter returns a Sink that writes to io.Writer out with strings
// separated by sepNames and pairs separated by sepPairs.
func MakeSinkToWriter(out io.Writer, sepNames string, sepPairs string) Sink {
	sink := Sink{
		make(chan FileRenameOp, defaultQueueLength),
		make(chan bool),
	}

	go func() {
		for s := range sink.Input {
			out.Write([]byte(s.OldName))
			out.Write([]byte(sepNames))
			out.Write([]byte(s.NewName))
			out.Write([]byte(sepPairs))
		}
		sink.Done <- true
	}()

	return sink
}

var (
	StdoutSink = MakeSinkToWriter(os.Stdout, " → ", "\n")
	NULSink    = MakeSinkToWriter(os.Stdout, "\x00", "\x00")
)
