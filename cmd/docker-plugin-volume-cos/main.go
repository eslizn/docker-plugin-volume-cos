package main

import (
	"fmt"
	"github.com/docker/go-plugins-helpers/volume"
	"github.com/marcelo-ochoa/docker-volume-plugins/mounted-volume"
	"log"
)

type Volume struct {
	*mountedvolume.Driver
}

func (v *Volume) Mount(req *volume.MountRequest) (*volume.MountResponse, error) {
	fmt.Printf("%+v\n", req)
	return v.Driver.Mount(req)
}

func main() {
	log.SetFlags(0)
	driver := &Volume{
		Driver: mountedvolume.NewDriver("cosfs", false, "cosfs", "local"),
	}
	defer driver.Close()
	driver.ServeUnix()
}
