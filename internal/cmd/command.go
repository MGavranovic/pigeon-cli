package cmd

type Command interface {
	Name() string
	Execute(args []string) error
}
