package client

import (
	"fmt"
	"testing"
	"time"

	"github.com/goodaye/fakeeyes/protos/request"
	"github.com/gorilla/websocket"
)

func TestConnectDevice(*testing.T) {

	name := "testuser"
	req := request.UserSignIn{
		Name: name,
	}
	user, err := client.SignIn(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := user.ConnectDevice("xxxxx")
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			_, data, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("receve err :", err.Error())
				break
			}
			fmt.Println("recieve message :", string(data))

		}
	}()
	for i := 0; i < 10; i++ {
		err = conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("message %d", i)))
		if err != nil {
			fmt.Println("send error :", err.Error())
		}
		fmt.Println("send message ", i)
	}
	time.Sleep(5 * time.Second)
	// err = conn.Close()
	if err != nil {
		panic(err)
	}
}

func TestListDevice(*testing.T) {

	name := "testuser"
	req := request.UserSignIn{
		Name: name,
	}
	user, err := client.SignIn(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	devs, err := user.ListDevices()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, dev := range devs {
		fmt.Println(dev)
	}
}
