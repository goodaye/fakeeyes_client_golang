package client

import (
	"github.com/goodaye/fakeeyes/protos/request"
	"github.com/goodaye/fakeeyes/protos/response"
)

type Device struct {
	client *Client
	SN     string
	Token  string
}

func (c *Client) NewDevice(token string) *Device {

	d := Device{
		client: c,
		Token:  token,
	}
	return &d
}

func (c *Client) RegisterDevice(req request.DeviceInfo) (dev *Device, err error) {
	// var err error
	apiname := "RegisterDevice"

	type rs struct {
		Data response.UserLogin
	}
	resp := &rs{}

	err = c.httpproxy(apiname, req, &resp)
	if err != nil {
		return
	}
	dev = c.NewDevice(resp.Data.Token)
	return
}

func (d *Device) SendHeartBeat() error {

	return nil
}
