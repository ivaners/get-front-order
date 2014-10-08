package model

import (
	"encoding/json"
	"get-front-order/common"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func FrontOrderDb(frontOrderList map[string]orm.Params) int {
	frontOrderInfo := new(FrontOrderInfo)
	gzbms := orm.NewOrm()
	gzbms.Using("default")
	frontOrderOrm := gzbms.QueryTable(frontOrderInfo)

	i := 0
	for _, v := range frontOrderList {
		count, _ := frontOrderOrm.Filter("SiteId", SiteId).Filter("FrontOrderId", v["order_id"]).Filter("FrontOrderSn", v["order_sn"]).Count()

		if count > 0 {
			continue
		}
		frontOrderInfo.HostId = HostId
		frontOrderInfo.HostCode = HostCode
		frontOrderInfo.SiteId = SiteId
		frontOrderInfo.FrontOrderSn = v["order_sn"].(string)
		frontOrderInfo.FrontOrderId, _ = strconv.ParseInt(v["order_id"].(string), 10, 64)
		frontOrderInfo.UserId, _ = strconv.ParseInt(v["user_id"].(string), 10, 64)
		frontOrderInfo.OrderStatus, _ = strconv.ParseInt(v["order_status"].(string), 10, 64)
		frontOrderInfo.GoodsAmount, _ = strconv.ParseFloat(v["goods_amount"].(string), 64)
		frontOrderInfo.OrderAmount, _ = strconv.ParseFloat(v["order_amount"].(string), 64)
		frontOrderInfo.IpAddress = v["ip_address"].(string)
		frontOrderInfo.IpInfoText = v["ip_info_text"].(string)
		frontOrderInfo.IntegralMoney, _ = strconv.ParseFloat(v["integral_money"].(string), 64)
		frontOrderInfo.PayFee, _ = strconv.ParseFloat(v["pay_fee"].(string), 64)
		frontOrderInfo.PayName = v["pay_name"].(string)
		frontOrderInfo.BestTime = v["best_time"].(string)
		frontOrderInfo.Surplus, _ = strconv.ParseFloat(v["surplus"].(string), 64)
		frontOrderInfo.MoneyPaid, _ = strconv.ParseFloat(v["money_paid"].(string), 64)
		frontOrderInfo.Tel = v["tel"].(string)
		frontOrderInfo.Mobile = v["mobile"].(string)
		frontOrderInfo.Qq = v["qq"].(string)
		frontOrderInfo.Email = v["email"].(string)
		frontOrderInfo.Zipcode = v["zipcode"].(string)
		frontOrderInfo.IntegralMoney, _ = strconv.ParseFloat(v["integral_money"].(string), 64)
		frontOrderInfo.City, _ = strconv.ParseInt(v["city"].(string), 10, 64)
		frontOrderInfo.Province, _ = strconv.ParseInt(v["province"].(string), 10, 64)
		frontOrderInfo.Consignee = v["consignee"].(string)
		frontOrderInfo.Country, _ = strconv.ParseInt(v["country"].(string), 10, 64)
		frontOrderInfo.District, _ = strconv.ParseInt(v["district"].(string), 10, 64)
		frontOrderInfo.Address = v["address"].(string)
		frontOrderInfo.Bonus, _ = strconv.ParseFloat(v["bonus"].(string), 64)
		frontOrderInfo.BonusId, _ = strconv.ParseInt(v["bonus_id"].(string), 10, 64)
		frontOrderInfo.PayStatus, _ = strconv.ParseInt(v["pay_status"].(string), 10, 64)
		frontOrderInfo.Postscript = v["postscript"].(string)
		frontOrderInfo.ShippingFee, _ = strconv.ParseFloat(v["shipping_fee"].(string), 64)
		frontOrderInfo.ShippingStatus, _ = strconv.ParseInt(v["shipping_status"].(string), 10, 64)
		frontOrderInfo.AddTime, _ = strconv.ParseInt(v["add_time"].(string), 10, 64)
		strJson, _ := json.Marshal(v["order_goods"])
		frontOrderInfo.OrderGoods = string(strJson)

		_, err := gzbms.Insert(frontOrderInfo)
		if err != nil {
			common.Log.Error("订单转入sql语句发生错误: %s", err)
		} else {
			i++
		}
	}
	return i
}
