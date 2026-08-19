//go:build !windows

package main

import "os/exec"

// esconderJanela não faz nada fora do Windows.
func esconderJanela(cmd *exec.Cmd) {}
