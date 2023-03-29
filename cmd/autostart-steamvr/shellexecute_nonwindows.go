//go:build !windows
// +build !windows

package main

import (
	"github.com/skratchdot/open-golang/open"
)

func ShellExecute(program, arguments, directory string) (err error) {
	return open.Start(program)
}
