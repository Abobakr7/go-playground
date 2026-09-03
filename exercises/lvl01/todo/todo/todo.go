package todo

import "fmt"

type Status int

const (
	Open Status = iota
	Done
)

func (s Status) String() string {
	switch s {
	case Open:
		return " "
	case Done:
		return "x"
	}
	return fmt.Sprintf("Unknown Status: <%d>", s)
}

type Todo struct {
	ID     int
	Title  string
	Status Status
}
