//go:build windows

package main

import (
	"os/exec"
	"syscall"
)

// esconderJanela evita que o console do PowerShell apareça no Windows.
func esconderJanela(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
}
