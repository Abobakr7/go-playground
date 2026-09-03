package cli

import (
	"fmt"
	"strconv"
	"strings"
	"todo/todo"
)

type Processor struct {
	store *todo.Store
}

func NewProcessor(store *todo.Store) *Processor {
	return &Processor{
		store: store,
	}
}

func (p *Processor) Process(command *Command) error {
	switch strings.ToLower(command.Name) {
	case "add":
		return p.add(command)
	case "list":
		return p.list(command)
	case "complete":
		return p.complete(command)
	case "delete":
		return p.delete(command)
	default:
		return fmt.Errorf("Unknown command: %s", command.Name)
	}
}

func (p *Processor) add(command *Command) error {
	if len(command.Args) == 0 {
		return fmt.Errorf("Missing title for new todo")
	}

	todo := p.store.Add(command.Args[0])
	fmt.Printf("Added todo: %d - %s\n", todo.ID, todo.Title)

	return nil
}

func (p *Processor) list(command *Command) error {
	if len(command.Args) != 0 {
		return fmt.Errorf("list does not accept arguments")
	}

	todos := p.store.List()
	for _, t := range todos {
		fmt.Printf("%d - %s [%s]\n", t.ID, t.Title, t.Status)
	}

	return nil
}

func (p *Processor) complete(command *Command) error {
	if len(command.Args) == 0 {
		return fmt.Errorf("Missing ID for todo to complete")
	}

	id, err := strconv.Atoi(command.Args[0])
	if err != nil {
		return fmt.Errorf("Invalid ID: %s", command.Args[0])
	}

	if err := p.store.Complete(id); err != nil {
		return err
	}

	fmt.Printf("Marked todo #%d as done\n", id)

	return nil
}

func (p *Processor) delete(command *Command) error {
	if len(command.Args) == 0 {
		return fmt.Errorf("Missing ID for todo to delete")
	}

	id, err := strconv.Atoi(command.Args[0])
	if err != nil {
		return fmt.Errorf("Invalid ID: %s", command.Args[0])
	}

	if err := p.store.Delete(id); err != nil {
		return err
	}

	fmt.Printf("Deleted todo #%d\n", id)

	return nil
}
