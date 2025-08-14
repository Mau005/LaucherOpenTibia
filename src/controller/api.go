package controller

import (
	"os/exec"

	"github.com/Mau005/LaucherOpenTibia/src/configuration"
)

type ApiController struct{}

func (ap *ApiController) RunClient() {

	cmd := exec.Command(configuration.API.PathCLient)
	_, err := cmd.Output()

	if err != nil {
		return
	}

}
