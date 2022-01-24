package client

import (
	"github.com/goodaye/fakeeyes/protos/request"
	"github.com/goodaye/fakeeyes/protos/response"
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
