package model

import (
	"fmt"
	"get-front-order/common"
	"github.com/astaxie/beego/orm"
)

type FrontOrderInfo struct {
	Id             int64
	HostId         int64
	HostCode       string
	SiteId         int64
	FrontOrderSn   string
	FrontOrderId   int64
	OrderId        int64
	OrderSn        string
	OrderGoods     string
	UserId         int64
	OrderStatus    int64
	ShippingStatus int64
	PayStatus      int64
	Consignee      string
	Country        int64
	Province       int64
	City           int64
	District       int64
	Address        string
	Zipcode        string
	Tel            string
	Mobile         string
	Qq             string
	Email          string
	BestTime       string
	Postscript     string
	PayName        string
	GoodsAmount    float64
	PayFee         float64
	ShippingFee    float64
	MoneyPaid      float64
	Surplus        float64
	IntegralMoney  float64
	Bonus          float64
	OrderAmount    float64
	AddTime        int64
	BonusId        int64
	Discount       float64
	IpAddress      string
	AcceptKefu     int64
	AssignTime     int64
	Status         int64
	TreatKefu      int64
	TreatTime      int64
	IpInfoText     string
	IsRelate       int64
	NedTel         string
	NedMobile      string
	InvalidReason  int64
}

var DbUser, DbPass, DbHost, DbPort, DbName, DbPrefix string
var SiteId int64
var HostId int64
var HostCode string

func init() {
	gzdb, _ := common.Conf.GetSection("gzbms")
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", gzdb["db_user"], gzdb["db_pass"], gzdb["db_host"], gzdb["db_port"], gzdb["db_name"]))
	orm.RegisterModel(new(FrontOrderInfo))
}
