package common

import (
	"github.com/Unknwon/goconfig"
	"github.com/astaxie/beego/logs"
	json "github.com/bitly/go-simplejson"
	"os"
)

var Conf *goconfig.ConfigFile
var err error
var Global *json.Json

var Log *logs.BeeLogger

func init() {
	Log = logs.NewLogger(10000)
	Log.SetLogger("file", `{"filename":"log/cron.log"}`)
	Log.EnableFuncCallDepth(true)

	Conf, err = goconfig.LoadConfigFile("conf/config.ini")
	if err != nil {
		Log.Error("配置文件读取错误：%s", err)
		os.Exit(-1)
	}

	r, err := os.Open("conf/server.json")
	if err != nil {
		Log.Error("配置文件读取错误：%s", err)
		os.Exit(-1)
	}

	Global, _ = json.NewFromReader(r)

}
