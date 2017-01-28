package grename

import (
	"io"
	"os"
)

type Sink func(string, string) error


// MakeSinkToWriter returns a Sink that writes to io.Writer out with strings
// separated by sepNames and pairs separated by sepPairs.
func MakeSinkToWriter(out io.Writer, sepNames string, sepPairs string) Sink {
    return func(OldName string, NewName string) error {
        out.Write([]byte(OldName))
        out.Write([]byte(sepNames))
        out.Write([]byte(NewName))
        out.Write([]byte(sepPairs))
        return nil
    }
}

var (
	StdoutSink = MakeSinkToWriter(os.Stdout, " → ", "\n")
	NULSink    = MakeSinkToWriter(os.Stdout, "\x00", "\x00")
)

//func MakeSinkFileRenamer() Sink {
//}
