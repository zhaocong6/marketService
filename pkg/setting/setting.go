package setting

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type httpType struct {
	Port int `yaml:"port"`
}

type rpcType struct {
	Port int `yaml:"port"`
}

type app struct {
	Version string `yaml:"version"`
	App     string `yaml:"app"`
	RunMode string `yaml:"run_mode"`

	Http httpType `yaml:"http"`
	Rpc  rpcType  `yaml:"rpc"`
}

var conf app

func init() {
	b, err := ioutil.ReadFile("../../conf/app.yaml")
	if err != nil {
		log.Panicln(err)
	}

	err = yaml.Unmarshal(b, &conf)
	if err != nil {
		log.Panicln(err)
	}

	loadHTTP()
	loadRPC()
	loadVer()
	loadApp()
	loadRunMode()
}

var HTTP httpType

func loadHTTP() {
	HTTP = conf.Http
}

var RPC rpcType

func loadRPC() {
	RPC = conf.Rpc
}

var Ver string

func loadVer() {
	Ver = conf.Version
}

var App string

func loadApp() {
	App = conf.App
}

var RunMode string

func loadRunMode() {
	RunMode = conf.RunMode
}
