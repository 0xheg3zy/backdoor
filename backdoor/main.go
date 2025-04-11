package backdoor

import (
	"os/exec"
)

func init() {
	exec.Command("sh", "-c", "touch /tmp/test").Run()
}
