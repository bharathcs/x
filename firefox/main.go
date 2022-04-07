package firefox

import (
	"os"
	"os/exec"
)

var firefoxExecutable string

func setup() error {
	path, err := exec.LookPath("firefox")
	if err != nil {
		return FirefoxExecutableNotFoundErr
	}

	info, err := os.Stat(path)
	firefoxExecutable = path
}

func main() {
	exec.Command("firefox", "")
}
