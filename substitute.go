package grename

import (
	"path/filepath"
	"regexp"
)

// StringRenamer is a function that changes a string
type StringRenamer func(string) string

// MakeRE2Renamer creates StringRenamer using Go's RE2 regular expression
// engine from the regexp package. The StringRenamer funtion returned
// matches its input string using match and substitutes all matches with
// subst. It match is not a valid regular expression, an error is returned
// and the StringRenamer function is nil.
func MakeRE2Renamer(match, subst string) (StringRenamer, error) {
	var (
		matchRE *regexp.Regexp
		err     error
	)
	matchRE, err = regexp.Compile(match)
	if err != nil {
		return nil, err
	}

	return func(input string) string {
		return matchRE.ReplaceAllString(input, subst)
	}, nil
}

func MakeFilenameFilter(sr StringRenamer) Filter {
	return func(inC <-chan Renamed) chan<- Renamed {
		outC := make(chan Renamed)
		go func() {
			for filename := range inC {
				dir, file := filepath.Split(filename.renamed)
				newFilename := sr(file)
				outC <- Renamed{filename.original, filepath.Join(dir, newFilename)}
			}
			close(outC)
		}()
		return outC
	}
}
