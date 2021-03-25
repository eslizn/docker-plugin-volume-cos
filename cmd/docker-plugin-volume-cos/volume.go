package main

import (
	"github.com/marcelo-ochoa/docker-volume-plugins/mounted-volume"
)

type Volume struct {
	*mountedvolume.Driver
}

func NewVolumeDriver() *Volume {
	return &Volume{
		Driver: mountedvolume.NewDriver("cosfs", false, "cosfs", "local"),
	}
}
