package client

import (
	"fmt"
	"testing"

	"github.com/goodaye/fakeeyes/protos/request"
)

func TestDeviceRegister(t *testing.T) {

	devinfo := request.DeviceInfo{
		SN:        "fakesn-xyz",
		Name:      "mockdevice",
		ModelName: "testmachine",
		ModelID:   "xyz--abc",
	}

	fmt.Println(devinfo)
	devcie, err := client.RegisterDevice(devinfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(devcie)

}
