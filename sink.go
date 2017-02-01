package grename

import (
	"io"
	"os"
)

const (
	FileOpCautious   = 1 << iota
	FileOpSensible   = 1 << iota
	FileOpAggressive = 1 << iota
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

type FatalErrorInputfileAccess struct {
	StatError error
}

func (e *FatalErrorInputfileAccess) Error() string {
	return e.StatError.Error()
}

type FatalErrorOutputExists struct {
	errormsg string
}

func (e *FatalErrorOutputExists) Error() string {
	return e.errormsg
}

// MakeSinkFileRenamer returns a Sink that renames files
func MakeSinkFileRenamer(fileOpMode int, dryrun bool) Sink {
	return func(oldName string, newName string) error {
		// skip equal names
		if oldName == newName {
			return nil
		}

		// check if input file exists and is a file
		_, oldStatErr := os.Lstat(oldName)

		if oldStatErr != nil {
			switch fileOpMode {
			case FileOpCautious:
				return &FatalErrorInputfileAccess{oldStatErr}
			case FileOpSensible, FileOpAggressive:
				// TODO: logging
				return nil
			}
		}

		// check if destination file exists
		_, err := os.Stat(newName)
		if err == nil {
			switch fileOpMode {
			case FileOpCautious:
				return &FatalErrorOutputExists{"File " + newName + " already exists"}
			case FileOpSensible:
				// TODO: logging
				return nil
			}
		}
		// TODO handle non-file dest

		if dryrun == true {
            // TODO: logging
			return nil
		} else {
            // TODO: logging
			return os.Rename(oldName, newName)
		}
	}
}
