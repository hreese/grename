package grename

import (
	"io"
	"os"
)

type Sink interface {
	Rename(string, string) error
	Done()
}

type sinkWriter struct {
	out      io.Writer
	sepNames string
	sepPairs string
}

func (s *sinkWriter) Rename(OldName string, NewName string) error {
	s.out.Write([]byte(OldName))
	s.out.Write([]byte(s.sepNames))
	s.out.Write([]byte(NewName))
	s.out.Write([]byte(s.sepPairs))
	return nil
}

func (s *sinkWriter) Done() {}

// MakeSinkToWriter returns a Sink that writes to io.Writer out with strings
// separated by sepNames and pairs separated by sepPairs.
func MakeSinkToWriter(out io.Writer, sepNames string, sepPairs string) Sink {
	return &sinkWriter{
		out,
		sepNames,
		sepPairs,
	}
}

var (
	StdoutSink = MakeSinkToWriter(os.Stdout, " → ", "\n")
	NULSink    = MakeSinkToWriter(os.Stdout, "\x00", "\x00")
)

//MakeSinkFileRenamer()
