package terminal

import (
	"os/exec"
	"strconv"
	"strings"
)

func Width() (result int) {
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