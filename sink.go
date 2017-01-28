package grename

import (
	"io"
	"os"
)

type Sink func(string, string) error

// MakeSinkToWriter returns a Sink that writes to io.Writer out with strings
// separated by sepNames and pairs separated by sepPairs.
func MakeSinkToWriter(out io.Writer, sepNames string, sepPairs string, skipIdentical bool) Sink {
	return func(oldName string, newName string) error {
		if skipIdentical == false || oldName != newName {
			out.Write([]byte(oldName))
			out.Write([]byte(sepNames))
			out.Write([]byte(newName))
			out.Write([]byte(sepPairs))
		}
		return nil
	}
}

var (
	StdoutSink = MakeSinkToWriter(os.Stdout, " → ", "\n")
	NULSink    = MakeSinkToWriter(os.Stdout, "\x00", "\x00")
)

func MakeSinkFileRenamer() Sink {
	return func(oldName string, newName string, overwriteDestination bool) error {
		// skip equal names
		if oldName == newName {
			return nil
		}

		oldStat, err := os.Stat(oldName)
		if err != nil {
		} // TODO
		newStat, err := os.Stat(newName)
		// destination file exists
		if err == nil {
			if !overwriteDestination {
				return nil
			}
		}
		return os.Rename(oldName, newName)
	}
}
