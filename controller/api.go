package controller

import (
	"os/exec"

	conf "github.com/Mau005/LaucherOpenTibia/configuration"
)

type ApiController struct{}

func (ap *ApiController) RunClient() {

	cmd := exec.Command(conf.API.PathCLient)
	_, err := cmd.Output()

	if err != nil {
		return
	}

}
