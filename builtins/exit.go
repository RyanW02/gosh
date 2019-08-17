package builtins

import "os"

type Exit struct {
}

func (Exit) Name() string {
	return "exit"
}

func (Exit) Exec(args []string) error {
	os.Exit(0)
	return nil
}

