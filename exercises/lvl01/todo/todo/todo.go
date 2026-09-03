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
		return "Open"
	case Done:
		return "Done"
	}
	return fmt.Sprintf("Unknown Status: <%d>", s)
}

type Todo struct {
	ID     int
	Title  string
	Status Status
}
