package main

import (
	"bufio"
	"fmt"
	"github.com/docopt/docopt-go"
	. "github.com/hreese/grename"
	_ "github.com/k0kubun/pp"
	"os"
)

var (
	opts   map[string]interface{}
	input  chan string
	filter Filter
	sink   Sink
	usage  string = `Rename files using regular expressions.

Usage:
  grename [-l|-z] [-nivs] [-c|-f] [-p|-0] MATCHRE SUBST [- | FILENAMES ... ]

Arguments:
  MATCHRE    Regular expression (RE2) matching parts of FILENAMES that are replaced
             by SUBST.
  SUBST      String to replace matched filename parts with. Submatches from MATCHRE
             can be used via $1 or ${name}. Use $$ to insert a single dollar sign.
             See https://golang.org/pkg/regexp/#Regexp.ReplaceAllString for details.
  FILENAMES  One or more filenames to rename. Ignored when -l or-z is used.

Options:
  -l, --inlines, -        Read list of input files from stdin (separated by newlines)
  -z, --in0               Read list of input files from stdin (separated by \0)
  -i, --case-insensitive  Make regular expression case insensitive
                          (Alternative: use the »(?i:RE)« flag in MATCHRE)
  -s, --skip-same         Don't output filenames that are unchanged (only affects -n and -0)
  -n, --dry-run           Run normally (including all tests) but do NOT rename anything.
  -p, --print             Don't rename files, write changes to stdout.
  -0, --out0              Don't rename files, write changes to stdout (separated by \0)
  -v, --verbose           Be verbose (currently unimplemented).
  -c, --cautious          Immediatly exit on error like nonexisting source files or
                          existing destination files.
  -f, --force             Overwrite existing files (default: skip).
`
)

func init() {
	// parse commandline options
	opts, err := docopt.Parse(usage, nil, true, "", true, false)

	// process parsing errors
	if err != nil {
		switch err.(type) {
		case *docopt.LanguageError:
			fmt.Println("This should not happen, please report this to the author (docopt.LanguageError): ", err.Error())
			os.Exit(63)
		case *docopt.UserError:
			fmt.Println("Invalid options. Did you supply MATCHRE and SUBST? Did you set contradicting options?", err.Error())
			os.Exit(64)
        default:
            fmt.Println("Unknown error, please report this to the author", err.Error())
            os.Exit(65)
		}
	}

    if len(opts) == 0 {
        os.Exit(1)
    }

    //pp.Print(opts)

	// build input source
	if opts["-l"] == true || opts["-"] == true {
		input = MakeSourceFromScanner(os.Stdin, bufio.ScanLines)
	} else if opts["--in0"] == true {
		input = MakeSourceFromScanner(os.Stdin, ScanNUL)
	} else {
		input = MakeSourceFromStrings(opts["FILENAMES"].([]string))
	}

	// build filter
	match := opts["MATCHRE"].(string)
	if opts["--case-insensitive"] == true {
		match = "(?i:" + match + ")"
	}
	// compile regular expression
	renamer, err := MakeRE2Renamer(match, opts["SUBST"].(string))
	if err != nil {
		fmt.Println("Error compiling regular expressions:", err.Error())
		os.Exit(66)
	}
	filter = MakeFilenameFilter(renamer)

	// build output sink
	if opts["--out0"] == true {
		sink = MakeSinkToWriter(os.Stdout, "\x00", "\x00", false)
	} else if opts["--dry-run"] == true {
		sink = MakeSinkToWriter(os.Stdout, " → ", "\n", false)
	} else {
		var mode int
		if opts["--cautious"] == true {
			mode = FileOpCautious
		} else if opts["--force"] == true {
			mode = FileOpAggressive
		} else {
			mode = FileOpSensible
		}
		sink = MakeSinkFileRenamer(mode, opts["--dry-run"].(bool))
	}
}

func main() {
    // iterate over input filenames
	for oldName := range input {
        // rename filename
		newName := filter(oldName)
        // write modification / rename file
		err := sink(oldName, newName)
        // handle errors
		switch err.(type) {
		case *FatalErrorInputfileAccess:
			fmt.Fprintf(os.Stderr, "Input file does not exists, existing: %s\n", oldName)
			os.Exit(128)
		case *FatalErrorOutputExists:
			fmt.Fprintf(os.Stderr, "Output file exists, existing: %s\n", newName)
			os.Exit(129)
		case *os.LinkError:
			fmt.Fprintf(os.Stderr, "Error while renaming: %s\n", err.Error())
			os.Exit(130)
		}
	}
    os.Exit(0)
}
