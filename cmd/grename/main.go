package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/k0kubun/pp"
	//. "github.com/hreese/grename"
	"os"
)

var (
    opts map[string]interface{}
	usage string = `Rename files using regular expressions.

Usage:
  grename [-l|-z] [-iv] [-c|-f] [-n|-0] MATCHRE SUBST [- | FILENAMES ... ]

Arguments:
  MATCHRE    Regular expression (RE2) matching parts of FILENAMES that are replaced
             by SUBST.
  SUBST      String to replace matched filename parts with. Submatches from MATCHRE
             can be used via $1 or ${name}. Use $$ to insert a single dollar sign.
             See https://golang.org/pkg/regexp/#Regexp.ReplaceAllString for details.
  FILENAMES  One or more filenames to rename. Ignored when -l, -z or -Z is used.

Options:
  -l, --inlines, -        Read list of input files from stdin (separated by newlines)
  -z, --in0               Read list of input files from stdin (separated by \0)
  -i, --case-insensitive  Make regular expression case insensitive
                          (Alternative: use the »(?i:RE)« flag in MATCHRE)
  -n, --dry-run           Don't rename files, write changes to stdout
  -0, --out0              Don't rename files, write changes to stdout (separated by \0)
  -v, --verbose           Be verbose.
  -c, --cautious          Immediatly exit on error like nonexisting source files or
                          existing destination files.
  -f, --force             Overwrite existing files (default: skip).
`
)


func main() {
    // parse commandline options
	opts, err := docopt.Parse(usage, nil, true, "", true, false)

    // process parsing errors
	if err != nil {
		switch err.(type) {
		case *docopt.LanguageError:
            fmt.Println("This should not happen, please report this to the author (docopt.LanguageError): ", err.Error())
			os.Exit(64)
		case *docopt.UserError:
			fmt.Println("Invalid options.", err.Error())
			os.Exit(65)
		}
	}

	pp.Print(opts)
}
