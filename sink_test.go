package grename

import (
	"bytes"
	"testing"
	//"fmt"
)

func TestMakeSinkToWriter(t *testing.T) {
	var (
		buf bytes.Buffer
	)

	s := MakeSinkToWriter(&buf, " -> ", "\n")

	tests := []string{"image.jpg", "öäüß µ”¹²³¬¼", "     ", "/foo/bar/baz/bum"}
	result := "image.jpg -> image.jpg.test\nöäüß µ”¹²³¬¼ -> öäüß µ”¹²³¬¼.test\n      ->      .test\n/foo/bar/baz/bum -> /foo/bar/baz/bum.test\n"

	for _, name := range tests {
		s.Input <- FileRenameOp{name, name + ".test"}
	}
	close(s.Input)
	<-s.Done

	if buf.String() != result {
		t.Error("Unexpected result")
	}
	//fmt.Printf("%#v\n", buf.String())
}
