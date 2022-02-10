package client

import (
	"encoding/json"
	"reflect"

	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/goodaye/fakeeyes/protos/response"
	"github.com/gorilla/websocket"
)

// Fakeeyes Client
type Client struct {
	Server string
	Token  string
	url    *url.URL
}

var APIPrefix = "/api/v1"

// NewClient
func NewClient(server string) (*Client, error) {
	client := Client{
		Server: server,
	}
	url, err := url.Parse(server)
	if err != nil {
		return nil, err
	}
	client.url = url

	return &client, nil
}

//
func (c *Client) httpproxy(api string, req interface{}, resp interface{}, header http.Header) error {
	var err error

	relurl := path.Join(APIPrefix, api)
	u, err := url.Parse(relurl)
	if err != nil {
		return err
	}
	queryURL := c.url.ResolveReference(u).String()

	var reqstr string
	if req == nil {
		reqstr = ""
	} else {

		reqbody, err := json.Marshal(req)
		if err != nil {
			return err
		}
		reqstr = string(reqbody)
	}
	httpreq, err := http.NewRequest(http.MethodPost, queryURL, strings.NewReader(reqstr))
	if err != nil {
		return err
	}
	// sb := c.sign(reqstr)
	// httpreq.Header.Add("Timestamp", fmt.Sprintf("%d", sb.Timestamp))
	// httpreq.Header.Add("Signature", sb.Sign)
	// httpreq.Header.Add("Accesskey", sb.Accesskey)
	if header != nil {
		httpreq.Header = header
	}
	httpclt := http.Client{}
	httpresp, err := httpclt.Do(httpreq)
	if err != nil {
		return err
	}
	defer httpresp.Body.Close()
	body, err := ioutil.ReadAll(httpresp.Body)
	if err != nil {
		return err
	}
	// fmt.Println(string(body))
	if httpresp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP Code: %d , HTTP Response: %s ", httpresp.StatusCode, string(body))
	}
	rm := response.ReturnMessage{}

	err = json.Unmarshal(body, &rm)
	if err != nil {
		return err
	}
	if !rm.Success {
		err = fmt.Errorf("ErrorCode: %s  ErrorMessage: %s", rm.ErrorCode, rm.ErrorMessage)
		return err
	}
	err = json.Unmarshal(body, resp)
	return err
}

func (c *Client) WSConnect(api string, req interface{}, header http.Header) (conn *websocket.Conn, err error) {

	v := ToQueryValue(req)
	u := url.URL{
		Scheme:   "ws",
		Host:     c.url.Host,
		Path:     path.Join(APIPrefix, api),
		RawQuery: v.Encode(),
	}

	conn, _, err = websocket.DefaultDialer.Dial(u.String(), header)
	if err != nil {
		return
	}
	return
}

func ToQueryValue(req interface{}) (val url.Values) {

	val = url.Values{}
	v := reflect.ValueOf(req)
	if req == nil {
		return
	}
	var key string
	for i := 0; i < v.NumField(); i++ {
		tagvalue := v.Type().Field(i).Tag.Get("form")
		if tagvalue != "" {
			key = strings.TrimSpace(tagvalue)
		} else {
			key = v.Type().Field(i).Name
		}
		// 如果是有个key value是数组，则需要特殊处理
		// if v.Field(i).Type().Kind() == reflect.Array || v.Field(i).Type().Kind() == reflect.Slice {
		// 	v.Field(i).Type().Kind()
		// } else {
		// }
		val.Set(key, v.Field(i).String())

	}
	return
}
