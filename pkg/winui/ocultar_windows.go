//go:build windows

package winui

import (
	"os/exec"
	"syscall"
)

func ocultarJanela(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
}
