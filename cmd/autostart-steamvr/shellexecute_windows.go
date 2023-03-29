//go:build windows
// +build windows

package main

import (
	"golang.org/x/sys/windows"
)

const (
	releaseOffset      = 2
	shellExecuteOffset = 9

	cSEE_MASK_DEFAULT = 0
)

func ShellExecute(program, arguments, directory string) (err error) {
	var (
		arguments16 *uint16
		directory16 *uint16
	)

	program16, _ := windows.UTF16PtrFromString(program)

	if len(arguments) > 0 {
		arguments16, _ = windows.UTF16PtrFromString(arguments)
	}

	if len(directory) > 0 {
		directory16, _ = windows.UTF16PtrFromString(directory)
	}

	err = windows.ShellExecute(0,
		nil,
		program16,
		arguments16,
		directory16,
		windows.SW_SHOWNORMAL)
	return
}
