package main

import (
	"fmt"

	"github.com/goodaye/fakeeyes/protos/request"
	client "github.com/goodaye/fakeeyes_client_golang"
)

func main() {

	var feclient *client.Client

	var server = "http://127.0.0.1:8080/"

	var err error
	feclient, err = client.NewClient(server)
	if err != nil {
		panic(err)
	}

	name := "testuser"
	req := request.UserSignUp{
		Name: name,
	}
	resp, err := feclient.SignUp(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)

}
