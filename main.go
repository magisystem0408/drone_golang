package main

import (
	"gotello/app/models"
	"gotello/config"
	"gotello/utils"
	"time"
)

func main() {
	//setting log
	utils.LoggingSettings(config.Config.LogFile)

	droneManager := models.NewDroneManager()
	droneManager.TakeOff()
	time.Sleep(10 * time.Second)
	droneManager.Land()
}
