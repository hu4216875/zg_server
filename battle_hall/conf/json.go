package conf

import (
	"encoding/json"
	"github.com/name5566/leaf/log"
	"io/ioutil"
)

var Server struct {
	ServerId       uint32
	LogLevel       string
	LogPath        string
	TCPAddr        string
	MaxConnNum     int
	ConsolePort    int
	ProfilePath    string
	MaxMongoDbConn uint32
	MongoDbUrl     string
	EtcdServerAddr []string
	HttpServer     string
}

func init() {
	data, err := ioutil.ReadFile("conf/server.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}
