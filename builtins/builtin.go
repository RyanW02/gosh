package builtins

type BuiltIn interface {
	Name() string
	Exec(args []string) error
}