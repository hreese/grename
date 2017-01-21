package grename

import (
	"testing"
    "strings"
)

func TestNewLinewiseInputIterator(t *testing.T) {
    tests := [][]string{
        []string{"foo", "bar", "baz", "123"},
    }

    for _, teststrings := range tests {
        inReader := strings.NewReader(strings.Join(teststrings, "\n"))
        iter := NewLinewiseInputIterator(inReader)
        resultstrings := make([]string, len(teststrings))
        for s, err := iter.Next() && err == nil { // XXX
            switch err {
            case nil:
                resultstrings = append(resultstrings, s)
            case DONE:
            default:
            }
        }
    }
}

