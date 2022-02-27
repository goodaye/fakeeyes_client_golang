package client

import (
	"fmt"
	"testing"

	"github.com/goodaye/fakeeyes/protos/request"
	"github.com/gorilla/websocket"
)

func TestDeviceRegister(t *testing.T) {

	devinfo := request.DeviceInfo{
		SN:        "fakesn_client-abc",
		Name:      "mockdevice",
		ModelName: "testmachine",
		ModelID:   "xyz--abc",
		CPUCores:  1,
		CPUVendor: "Intel",
		CPUArch:   "arm",
	}

	fmt.Println(devinfo)
	devcie, err := client.RegisterDevice(devinfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(devcie)

}

func TestDeviceConnect(t *testing.T) {
	devinfo := request.DeviceInfo{
		SN:        "fakesn_client-abc",
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
	conn, err := devcie.Connect()

	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		mt, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		if mt == websocket.TextMessage {
			fmt.Println(string(p))
		}
	}
}
