package grename

type Filter func(string) string

type FileRenameOp struct {
	OldName string
	NewName string
}

type Sink chan FileRenameOp
