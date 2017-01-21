package grename

import (
	"bufio"
	"errors"
	"io"
)

var (
	DONE = errors.New("Iterator finished")
)

// InputIterator returns its elements one by one
type InputIterator struct {
	// Next returns the next filename, nil and finally nil, DONE
	Next func() (string, error)
}

func NewLinewiseInputIterator(infile io.Reader) InputIterator {
	scanner := bufio.NewScanner(infile)
	return InputIterator{
		Next: func() (string, error) {
			if scanner.Scan() {
				return scanner.Text(), nil
			} else {
				if err := scanner.Err(); err == nil {
					return "", DONE
				} else {
					return "", err
				}
			}
		},
	}
}
