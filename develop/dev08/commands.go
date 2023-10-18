package main

import (
	"fmt"
	"os"
)

func cd(argv []string) {
	if len(argv) != 2 {
		fmt.Fprintln(os.Stderr, "cd: too many arguments")
		return
	}
	err := os.Chdir(argv[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func kill(argv []string) {

}

func ps(argv []string) {

}
