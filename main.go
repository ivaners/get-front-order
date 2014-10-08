package main

import (
	"fmt"
	. "get-front-order/common"
	"get-front-order/model"
	. "github.com/bitly/go-simplejson"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //使用多核
	data := make(chan map[string]string) // 数据交换队列
	exit := make(chan bool)              // 退出通知
	var global *Json = Global
	go func() {
		for v := range data { // 从队列迭代接收数据，直到 close 。
			siteId, _ := strconv.Atoi(v["site_id"])
			model.SiteId = int64(siteId)

			model.HostId, _ = global.Get(v["host_code"]).Get("id").Int64()
			host_code := strings.Split(v["host_code"], ",")

			//如果配置有多个源，从多个源抓取数据
			for _, val := range host_code {
				model.HostCode = val

				model.DbUser, _ = global.Get(val).Get("db_user").String()
				model.DbPass, _ = global.Get(val).Get("db_pass").String()
				model.DbHost, _ = global.Get(val).Get("db_host").String()
				model.DbPort, _ = global.Get(val).Get("db_port").String()

				model.DbName = v["db_name"]
				model.DbPrefix = v["db_prefix"]

				orderList, err := model.GetOrder()
				if err != nil {
					Log.Error("sql查询错误: %s", err)
				}
				num := model.FrontOrderDb(orderList)
				Log.Info("从 %s.%s 查询到 %d 条数据,成功转入订单 %d 条", val, v["db_name"], len(orderList), num)
			}

		}
		fmt.Println("退出队列")
		exit <- true
	}()

	for _, val := range Conf.GetSectionList() {
		if val == "gzbms" {
			continue
		}

		getConf, err := Conf.GetSection(val)
		if err != nil {
			continue
		}
		data <- getConf
	}

	close(data)
	fmt.Println("队列执行完毕")
	<-exit
}
