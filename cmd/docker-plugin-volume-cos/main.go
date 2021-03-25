package main

import "log"

func main() {
	log.SetFlags(0)
	driver := NewVolumeDriver()
	defer driver.Close()
	driver.ServeUnix()
}
