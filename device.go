package golosetup

import (
	"encoding/json"
)

type Device struct {
	Name      string `json:"name"`
	BackFile  string `json:"back-file"`
	SizeLimit string `json:"sizelimit"`
	Offset    string `json:"offset"`
	AutoClear string `json:"autoclear"`
	ReadOnly  string `json:"ro"`
	Dio       string `json:"dio"`
	LogSec    string `json:"log-sec"`
}

type Devices struct {
	LoopDevs []Device `json:"loopdevices"`
}

func parseDevices(output string) (*Devices, error) {
	devs := &Devices{}
	if len(output) == 0 {
		return devs, nil
	}

	err := json.Unmarshal([]byte(output), devs)
	return devs, err
}

func (devs Devices) GetDeviceByName(name string) *Device {
	for _, dev := range devs.LoopDevs {
		if dev.Name == name {
			return &dev
		}
	}
	return nil
}

func (devs Devices) GetDeviceByFile(filePath string) *Device {
	for _, dev := range devs.LoopDevs {
		if dev.BackFile == filePath {
			return &dev
		}
	}
	return nil
}
