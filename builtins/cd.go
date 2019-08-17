package builtins

import "os"

type Cd struct {
}

func (Cd) Name() string {
	return "cd"
}

func (Cd) Exec(args []string) error {
	return os.Chdir(args[0])
}
