package client

import (
	"net/http"

	"github.com/goodaye/fakeeyes/protos"
	"github.com/goodaye/fakeeyes/protos/request"
	"github.com/goodaye/fakeeyes/protos/response"
	"github.com/gorilla/websocket"
)

type User struct {
	client *Client
	token  string
}

//
func (c *Client) NewUser(token string) *User {

	user := User{
		client: c,
		token:  token,
	}
	return &user
}

// 用户登陆
func (c *Client) SignIn(req request.UserSignIn) (user *User, err error) {

	// var err error
	apiname := "/UserSignIn"

	type rs struct {
		Data response.UserLogin
	}
	resp := &rs{}

	err = c.httpproxy(apiname, req, &resp)
	if err != nil {
		return
	}
	user = c.NewUser(resp.Data.Token)
	return
}

// 用户注册
func (c *Client) SignUp(req request.UserSignUp) (user *User, err error) {

	apiname := "/UserSignUp"

	type rs struct {
		Data response.UserLogin
	}
	resp := &rs{}

	err = c.httpproxy(apiname, req, &resp)
	if err != nil {
		return
	}
	user = c.NewUser(resp.Data.Token)
	return
}

//
func (u *User) ConnectDevice(device_uuid string) (conn *websocket.Conn, err error) {
	api := "/User/ConnectDevice"

	req := request.ConnectDevice{
		DeviceUUID: device_uuid,
	}

	header := http.Header{}
	header.Add(protos.HeaderKey.UserToken, u.token)
	conn, err = u.client.WSConnect(api, req, header)
	return
}
