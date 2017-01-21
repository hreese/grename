package grename

type Renamed struct {
    Original string
    Renamed string
}

type Source func() <-chan Renamed
type Filter func(<-chan Renamed) <-chan Renamed
type Sink   func(<-chan Renamed)

