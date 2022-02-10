package client

import (
	"net/http"

	"github.com/goodaye/fakeeyes/protos"
	"github.com/goodaye/fakeeyes/protos/request"
	"github.com/goodaye/fakeeyes/protos/response"
	"github.com/gorilla/websocket"
)

type Device struct {
	client *Client
	SN     string
	token  string
}

func (c *Client) NewDevice(token string) *Device {

	d := Device{
		client: c,
		token:  token,
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

	err = c.httpproxy(apiname, req, &resp, nil)
	if err != nil {
		return
	}
	dev = c.NewDevice(resp.Data.Token)
	return
}

func (d *Device) SendHeartBeat(req request.DeviceInfo) (err error) {
	// var err error
	apiname := "SendHeartBeat"
	type rs struct {
	}
	resp := &rs{}

	err = d.client.httpproxy(apiname, req, &resp, nil)
	if err != nil {
		return err
	}
	return nil
}

func (d *Device) Connect() (conn *websocket.Conn, err error) {
	api := "/Device/Connect"
	header := http.Header{}
	header.Add(protos.HeaderKey.DeviceToken, d.token)
	conn, err = d.client.WSConnect(api, nil, header)
	return
}
