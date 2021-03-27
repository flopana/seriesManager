package commands

import (
	"os/exec"
	"serienManager/config"
)

func LaunchVlv(config config.Config, filePath string)  {
	exec.Command(config.VlcPath, filePath)
}
