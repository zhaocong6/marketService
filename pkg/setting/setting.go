package setting

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type app struct {
	Version string `yaml:"version"`
	App     string `yaml:"app"`
	RunMode string `yaml:"run_mode"`

	Http httpType `yaml:"http"`
	Rpc  rpcType  `yaml:"rpc"`
	DB   dbType   `yaml:"db"`
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
	loadDB()
}

type httpType struct {
	Port int `yaml:"port"`
}

var HTTP *httpType

func loadHTTP() {
	HTTP = &conf.Http
}

type rpcType struct {
	Port int `yaml:"port"`
}

var RPC *rpcType

func loadRPC() {
	RPC = &conf.Rpc
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

type dbType struct {
	Type      string `yaml:"type"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	Database  string `yaml:"database"`
	Charset   string `yaml:"charset"`
	Collation string `yaml:"collation"`
}

var DB *dbType

func loadDB() {
	DB = &conf.DB
}
