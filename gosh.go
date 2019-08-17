package main

import(
"bufio"
"fmt"
	"gosh/builtins"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		cwd, err := os.Getwd(); if err != nil {
			cwd = ""
		}
		fmt.Print(fmt.Sprintf("%s > ", cwd))

		input, err := reader.ReadString('\n'); if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err.Error())
		}

		split := strings.Split(strings.TrimSuffix(input, "\n"), " ")
		root := split[0]
		args := split[1:]

		isBuiltIn := false
		for _, builtin := range builtins.BuiltIns {
			if builtin.Name() == root {
				isBuiltIn = true
				if err = builtin.Exec(args); err != nil {
					_, _ = fmt.Fprintln(os.Stderr, err.Error())
				}
			}
		}

		if !isBuiltIn {
			cmd := exec.Command(root, args...)
			cmd.Stdout = os.Stdout
			cmd.Stdin = os.Stdin
			cmd.Stderr = os.Stderr
			if err = cmd.Run(); err != nil {
				_, _ = fmt.Fprintln(os.Stderr, err.Error())
			}
		}
	}
}