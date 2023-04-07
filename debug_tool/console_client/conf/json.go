package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var Server struct {
	UserId     string
	ServerAddr string
}

func init() {
	data, err := ioutil.ReadFile("conf/server.json")
	if err != nil {
		fmt.Errorf("err:%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		fmt.Errorf("err:%v", err)
	}
}
