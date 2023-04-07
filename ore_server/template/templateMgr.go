package template

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/recordfile"
	"ore_server/conf"
	"reflect"
)

var (
	oreTemplate OreInfoTemplate
)

func readRf(st interface{}) *recordfile.RecordFile {
	rf, err := recordfile.New(st)
	if err != nil {
		log.Fatal("readRf err:%v", err)
	}
	fn := reflect.TypeOf(st).Name() + ".csv"
	err = rf.Read(conf.Server.GameDataPath + fn)
	if err != nil {
		log.Fatal("%v: %v", fn, err)
	}
	return rf
}

func LoadTempalte() {
	oreTemplate.load()
}

func GetOreTempalte() *OreInfoTemplate {
	return &oreTemplate
}
