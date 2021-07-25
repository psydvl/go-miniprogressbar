package goltools

import (
	"os/exec"
	"strconv"
	"strings"
)

//TODO: check in Powershell
func TerminalWidth() (result int) {
	result = 0
	out, err := exec.Command("tput", "cols").Output()
	if err == nil {
		result, _ = strconv.Atoi(
			strings.TrimSpace(
				string(out),
			),
		)
	}
	return result
}