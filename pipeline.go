package grename

type Source func(chan<- string)
type Filter func(<-chan string, chan<- string)
type Sink func(<-chan string)

