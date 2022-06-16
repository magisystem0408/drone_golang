package models

import (
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
	"golang.org/x/sync/semaphore"
	"time"
)

const (
	DefaultSpeed      = 10
	WaitDroneStartSec = 10
)

type DroneManager struct {
	*tello.Driver
	Speed        int
	patrollSem   *semaphore.Weighted
	patrolQuit   chan bool
	isPatrolling bool
}

func NewDroneManager() *DroneManager {
	drone := tello.NewDriver("8889")
	droneManager := &DroneManager{
		Driver:       drone,
		Speed:        DefaultSpeed,
		patrollSem:   semaphore.NewWeighted(1),
		patrolQuit:   make(chan bool),
		isPatrolling: false,
	}
	work := func() {

	}
	robot := gobot.NewRobot("tello", []gobot.Connection{}, []gobot.Device{drone}, work)
	go robot.Start()
	time.Sleep(WaitDroneStartSec * time.Second)
	return droneManager
}
