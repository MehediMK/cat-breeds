package utils

import (
	"encoding/json"
	"fmt"

	"github.com/beego/beego/httplib"
)

// for get api request
func Get_api_request(url string) string {
	req := httplib.Get(url)
	res, err := req.String()
	if err != nil {
		fmt.Println(err)
	}
	return res
}

// here use unmarchal
func Unmarshaldata(byte_data string, api_data interface{}) error {
	err := json.Unmarshal([]byte(byte_data), &api_data)
	if err != nil {
		fmt.Println("Here some error get")
	}
	return err
}
