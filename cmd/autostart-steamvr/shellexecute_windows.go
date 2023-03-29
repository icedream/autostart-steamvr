//go:build windows
// +build windows

package main

import (
	"golang.org/x/sys/windows"
)

func ShellExecute(program string) (err error) {
	program16, _ := windows.UTF16PtrFromString(program)

	err = windows.ShellExecute(0,
		nil,
		program16,
		nil,
		nil,
		windows.SW_SHOWNORMAL)
	return
}
