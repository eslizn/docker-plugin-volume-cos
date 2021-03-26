package main

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/docker/go-plugins-helpers/volume"
	mountedvolume "github.com/marcelo-ochoa/docker-volume-plugins/mounted-volume"
)

type Option struct {
	Bucket    string
	AppId     string
	SecretId  string
	SecretKey string
}

func (o Option) String() string {
	return fmt.Sprintf("%s-%s:%s:%s", o.Bucket, o.AppId, o.SecretId, o.SecretKey)
}

type Volume struct {
	sync.RWMutex
	mountedvolume.Driver
	Options map[string]Option
}

func (v *Volume) Validate(req *volume.CreateRequest) error {
	var args = []string{"app_id", "secret_id", "secret_key"}
	for _, v := range args {
		if _, ok := req.Options[v]; !ok {
			return fmt.Errorf("argument: %s missing", v)
		}
	}
	return nil
}

func (v *Volume) MountOptions(req *volume.CreateRequest) []string {
	option := Option{
		Bucket:    req.Name,
		AppId:     req.Options["app_id"],
		SecretId:  req.Options["secret_id"],
		SecretKey: req.Options["secret_key"],
	}
	v.Lock()
	defer v.Unlock()
	v.Options[req.Name] = option
	//return v.Driver.MountOptions(req)
	return []string{fmt.Sprintf("%s-%s", req.Name, option.AppId)}
}

func (v *Volume) PreMount(*volume.MountRequest) error {
	return nil
}

func (v *Volume) PostMount(*volume.MountRequest) {}

func (v *Volume) Mount(req *volume.MountRequest) (*volume.MountResponse, error) {
	v.RLock()
	defer v.RUnlock()
	option, ok := v.Options[req.Name]
	if !ok {
		return nil, errors.New("option missing")
	}
	fp, err := os.OpenFile("/etc/passwd-cosfs", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0640)
	if err != nil {
		return nil, err
	}
	_, err = fp.WriteString(option.String())
	if err != nil {
		return nil, err
	}
	return v.Driver.Mount(req)
}

func main() {
	//log.SetFlags(0)
	driver := &Volume{
		Options: make(map[string]Option),
		Driver:  *mountedvolume.NewDriver("cosfs", false, "cosfs", "local"),
	}
	driver.Init(driver)
	defer driver.Close()
	driver.ServeUnix()
}
