package controllers

import (
	"ally/config"
	"ally/model"
	"ally/utils"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Xinhua(c *gin.Context) {

	type Result struct {
		Code      string `json:"code" tag:"编码"`
		Status    string `json:"status" tag:"状态"`
		Remark    string `json:"remark" tag:"标记"`
		Mobile    string `json:"mobile" tag:"手机号"`
		Contact   string `json:"contact" tag:"联系人"`
		Organ     string `json:"organ" tag:"机构"`
		Work_num  string `json:"work_num" tag:"工号"`
		Order_no  string `json:"order_no" tag:"订单号"`
		Contact1  string `json:"contact1" tag:"收货人"`
		Mobile1   string `json:"mobile1" tag:"收货手机"`
		Province  string `json:"province" tag:"省"`
		City      string `json:"city" tag:"市"`
		Area      string `json:"area" tag:"区"`
		Address   string `json:"address" tag:"地址"`
		Ship_name string `json:"ship_name" tag:"快递公司"`
		Ship_no   string `json:"ship_no" tag:"快递单号"`
		C_time    string `json:"c_time" tag:"创建时间"`
	}

	var result []Result

	sqlQuery := "select a.code,a.status,a.remark,c.work_num,c.mobile,c.contact,c.organ,d.order_no,d.contact as contact1,d.mobile as mobile1,d.province,d.city,d.area,d.address,d.ship_name,d.ship_no,d.c_time from car_coupon a left join car_member b on a.user_id = b.id LEFT JOIN car_order_photo_agent c  on b.mobile = c.mobile and c.company = 19 LEFT JOIN car_order_photo d on a.id = d.coupon_id and d.status != -1 where a.batch_num = 'P2401291323' and a.status !=0"

	db := model.RDBs[model.MASTER]
	db.Db.Raw(sqlQuery).Find(&result)
	for k, v := range result {
		status := "已激活"
		remark := "未分享"
		num, _ := strconv.Atoi(v.Status)
		if v.Remark == "1" {
			remark = "已分享"
		} else if v.Remark != "" && v.Remark != "1" {
			remark = "已领取"
		}
		if num == 2 {
			status = "已下单"
			remark = "已下单"
		}
		result[k].Status = status
		result[k].Remark = remark
		if v.C_time != "" {
			result[k].C_time = utils.FormatDateByString(v.C_time)
		}
	}

	utils.Down(result, "新华保险摆台", c)

}

func Fjpa(c *gin.Context) {

	type Result struct {
		Name          string `json:"nume" tag:"业务员姓名"`
		Mobile        string `json:"mobile" tag:"手机号"`
		Work_num      string `json:"work_num" tag:"工号"`
		Contact       string `json:"contact" tag:"营服"`
		Organ         string `json:"organ" tag:"机构"`
		Num           int    `json:"num" tag:"权益数量"`
		Code          string `json:"code" tag:"卡券编码"`
		Active_time   string `json:"active_time" tag:"激活时间"`
		Remark        string `json:"remark" tag:"分享状态"`
		Status        string `json:"status" tag:"卡券状态"`
		Order_no      string `json:"order_no" tag:"订单号"`
		Contact1      string `json:"contact1" tag:"收货人"`
		Mobile1       string `json:"mobile1" tag:"收货手机"`
		Customer_info string `json:"customer_info" tag:"客户备注"`
		Address       string `json:"address" tag:"地址"`
		Ship_name     string `json:"ship_name" tag:"快递公司"`
		Ship_no       string `json:"ship_no" tag:"快递单号"`
		C_time        string `json:"c_time" tag:"创建时间"`
	}

	var result []Result

	sqlQuery := "select a.name,a.mobile,a.work_num,a.contact ,a.organ ,b.num ,c.code ,if(c.active_time <>0,FROM_UNIXTIME(c.active_time),'') as active_time,if(c.remark is NULL,'未分享',if(c.remark=1,'已分享','已领取')) as 'remark',case c.status when '0' then '' when 1 then '已激活' when 2 then '已下单' end as status ,d.order_no,d.contact as cus_contact,d.mobile as cus_mobile,concat(d.province,d.city,d.area,d.address) as address,d.customer_info,d.ship_no,d.ship_name,if(d.c_time<>0,FROM_UNIXTIME(d.c_time),'') as c_time from car_order_photo_agent a left join ( select mobile,sum(num) as num from car_member_bind_logs where coupon_batch = 'P2403121036' and is_del = 0 GROUP BY mobile) b on a.mobile = b.mobile LEFT JOIN car_coupon c on c.batch_num = 'P2403121036' and c.mobile = a.mobile LEFT JOIN car_order_photo d on c.id = d.coupon_id and d.`status` != -1 where a.company = 22"
	type Customer struct {
		Contact  string `json:"contact"`
		Work_num int    `json:"work_num"`
	}

	db := model.RDBs[model.MASTER]
	db.Db.Raw(sqlQuery).Find(&result)
	for k, v := range result {
		if v.Customer_info != "" {
			var tom Customer
			err := json.Unmarshal([]byte(v.Customer_info), &tom)
			if err == nil {
				result[k].Customer_info = tom.Contact
			}
		}
	}

	utils.Down(result, "福建平安摆台", c)

}

func Ydln(c *gin.Context) {
	type Result struct {
		Code            string `json:"code" tag:"卡券编号"`
		Name            string `json:"name" tag:"卡券名称"`
		Sn              string `json:"sn" tag:"卡券编码"`
		Password        string `json:"password" tag:"兑换码"`
		Active_time     string `json:"active_time" tag:"激活时间"`
		Remark          string `json:"remark" tag:"分享状态"`
		Status          string `json:"status" tag:"卡券状态"`
		Order_no        string `json:"order_no" tag:"订单号"`
		Contact1        string `json:"contact1" tag:"收货人"`
		Mobile1         string `json:"mobile1" tag:"收货手机"`
		Customer_info   string `json:"customer_info" tag:"客户姓名"`
		Customer_mobile string `json:"customer_mobile" tag:"客户手机"`
		Address         string `json:"address" tag:"地址"`
		Ship_name       string `json:"ship_name" tag:"快递公司"`
		Ship_no         string `json:"ship_no" tag:"快递单号"`
		C_time          string `json:"c_time" tag:"创建时间"`
	}

	var result []Result

	sqlQuery := "select a.code,a.name,a.sn,a.`password`,if(a.active_time,FROM_UNIXTIME(a.active_time),'') as active_time,a.mobile,if(b.remark is NULL,'未分享',if(b.remark=1,'已分享','已领取')) as remark,case b.status when '0' then '' when 1 then '已激活' when 2 then '已下单' end as status ,c.order_no,c.contact as contact1,c.mobile as mobile1,concat(c.province,c.city,c.area,c.address) as address,c.customer_info,c.ship_no,c.ship_name,if(c.c_time<>0,FROM_UNIXTIME(c.c_time),'') as c_time from car_coupon_pkg a LEFT JOIN car_coupon b on (a.id = b.pkg_id) LEFT JOIN car_order_photo c on c.coupon_id = b.id and c.`status` <> -1 WHERE a.batch_num in ('PP2403041702','PP2403061702')"
	type Customer struct {
		Contact string `json:"contact"`
		Mobile  string `json:"mobile"`
	}

	db := model.RDBs[model.MASTER]
	db.Db.Raw(sqlQuery).Find(&result)
	for k, v := range result {
		if v.Customer_info != "" {
			var tom Customer
			err := json.Unmarshal([]byte(v.Customer_info), &tom)
			if err == nil {
				result[k].Customer_info = tom.Contact
				result[k].Customer_mobile = tom.Mobile
			}
		}
	}

	utils.Down(result, "英大辽宁摆台", c)
}

func ExcelDown(c *gin.Context) {
	dataBase := c.Query("database")
	db, err := config.GetDb()
	if dataBase != "" {
		db, err = config.GetDbDatabase(dataBase)
	}
	if err != nil {
		slog.Error("数据库连接失败", err)
		c.String(http.StatusOK, err.Error())
	}
	slog.Info("数据库信息", db)

}
