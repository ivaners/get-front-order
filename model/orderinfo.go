package model

import (
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
	"time"
)

/**
 * 获取前台订单
 */
func GetOrder() (map[string]orm.Params, error) {
	orm.RegisterDataBase(HostCode+DbName, "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", DbUser, DbPass, DbHost, DbPort, DbName))
	orderDb := orm.NewOrm()
	orderDb.Using(HostCode + DbName)
	getTime := time.Now().Unix() - 86400*10

	where := ""
	switch SiteId {
	case 14:
		where = " AND divide_region = '广州地区手机商城下单' "
	}

	var maps []orm.Params
	sql := "SELECT * FROM " + DbPrefix + "order_info WHERE  1 AND (order_status = 0 OR pay_status = 2) AND add_time > '" + strconv.FormatInt(getTime, 10) + "' AND referer <> '200后台推送' " + where + " ORDER BY order_id desc LIMIT 300"
	_, err := orderDb.Raw(sql).Values(&maps)

	if len(maps) <= 0 {
		return nil, nil
	}

	orderIdArr := make([]string, 0)
	data := make(map[string]orm.Params, 0)
	for _, v := range maps {
		orderId := v["order_id"].(string)
		orderIdArr = append(orderIdArr, orderId)
		data[orderId] = v
	}

	sql = "SELECT og.goods_id,og.extension_code,og.order_id,og.goods_name,og.goods_sn,og.goods_number,og.goods_price,og.is_gift,f.unique_id AS gift_unique_id,g.unique_id AS package_unique_id FROM " + DbPrefix + "order_goods AS og LEFT JOIN " + DbPrefix + "favourable_activity AS f ON og.is_gift = f.act_id LEFT JOIN " + DbPrefix + "goods_activity AS g ON g.act_id = og.goods_id WHERE og.order_id IN (" + strings.Join(orderIdArr, ",") + ") AND og.extension_code <> 'package_goods' ORDER BY og.order_id ASC "
	_, err = orderDb.Raw(sql).Values(&maps)

	orderGoods := make(map[string][]orm.Params, 0)
	for _, v := range maps {
		key, _ := v["order_id"].(string)
		goodsName, _ := v["goods_name"].(string)
		v["goods_name"] = base64.StdEncoding.EncodeToString([]byte(goodsName))
		orderGoods[key] = append(orderGoods[key], v)
		data[key]["order_goods"] = orderGoods[key]

	}

	return data, err
}
