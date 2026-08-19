//go:build !windows

package winui

import "os/exec"

func ocultarJanela(*exec.Cmd) {}
