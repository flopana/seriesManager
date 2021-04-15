package commands

import (
	"os/exec"
	"seriesManager/config"
)

func LaunchVlv(config config.Config, filePath string)  {
	exec.Command(config.VlcPath, filePath)
}
