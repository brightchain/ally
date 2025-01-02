package controllers

import (
	"encoding/json"
	"fmt"
	"h5/pkg/model"
	"h5/utils"
	"log/slog"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ExportExcel struct{}

func (*ExportExcel) Xinhua(c *gin.Context) {
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

	db := model.RDB[model.MASTER]
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

func (*ExportExcel) Fjpa(c *gin.Context) {

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

	db := model.RDB[model.MASTER]
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

func (*ExportExcel) Ydln(c *gin.Context) {
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

	db := model.RDB[model.MASTER]
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

func (*ExportExcel) ShTp(c *gin.Context) {
	type Result struct {
		Code            string `json:"code" tag:"卡券编号"`
		Name            string `json:"name" tag:"卡券名称"`
		Sn              string `json:"sn" tag:"卡券编码"`
		Password        string `json:"password" tag:"兑换码"`
		Active_time     string `json:"active_time" tag:"激活时间"`
		Mobile          string `json:"mobile" tag:"手机号"`
		Work_num        string `json:"work_num" tag:"业务员工号"`
		Name1           string `json:"name1" tag:"业务员姓名"`
		Organ           string `json:"organ" tag:"机构名称"`
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

	sqlQuery := "select a.code,a.name,a.sn,a.`password`,if(a.active_time,FROM_UNIXTIME(a.active_time),'') as active_time,a.mobile,d.name as name1,d.work_num,d.organ,if(b.remark is NULL,'未分享',if(b.remark=1,'已分享','已领取')) as remark,case b.status when '0' then '未激活' when 1 then '已激活' when 2 then '已下单' end as status ,c.order_no,c.contact as contact1,c.mobile as mobile1,concat(c.province,c.city,c.area,c.address) as address,c.customer_info,c.ship_no,c.ship_name,if(c.c_time<>0,FROM_UNIXTIME(c.c_time),'') as c_time from car_coupon_pkg a LEFT JOIN car_coupon b on (a.id = b.pkg_id) LEFT JOIN car_order_photo_agent d on a.mobile = d.mobile and d.company = 24 and a.mobile <> 0 LEFT JOIN car_order_photo c on c.coupon_id = b.id and c.`status` <> -1 WHERE a.batch_num ='PP2404301621'"
	type Customer struct {
		Contact string `json:"contact"`
		Mobile  string `json:"mobile"`
	}

	db := model.RDB[model.MASTER]
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

	utils.Down(result, "上海太平个险", c)
}

func (*ExportExcel) FjTp(c *gin.Context) {
	type Result struct {
		OrderNo      string `json:"order_no" tag:"订单编号"`
		Name         string `json:"name" tag:"产品名称"`
		Num          string `json:"num" tag:"购买数量"`
		Order_amount string `json:"order_amount" tag:"订单金额"`
		PayNo        string `json:"pay_no" tag:"支付单号"`
		PayAt        string `json:"pay_at" tag:"支付时间"`
		Mobile       string `json:"mobile" tag:"手机号"`
		Work_num     string `json:"work_num" tag:"业务员工号"`
		Name1        string `json:"name1" tag:"业务员姓名"`
		Contact      string `json:"contact" tag:"中支"`
		Organ        string `json:"organ" tag:"营服"`
		Status       string `json:"status" tag:"订单状态"`
		C_time       string `json:"c_time" tag:"创建时间"`
	}

	var result []Result

	sqlQuery := "select a.order_no, '福建太平10寸照片摆台' as name,a.num,a.order_amount,a.pay_no,if(a.pay_at,FROM_UNIXTIME(a.pay_at),'') as 'pay_at',case a.status when 0 then '未付款' when 1 then '已付款' when 2 then '已完成' when -1 then '已取消' end as 'status',b.name as 'name1',b.mobile,b.contact,b.organ,b.work_num,FROM_UNIXTIME(a.c_time) as 'c_time'  from car_order_gdpa a LEFT JOIN car_order_photo_agent b on (a.uid = b.uid and b.company = 30) where a.pro_id = 'TP001' "

	db := model.RDB[model.MASTER]
	db.Db.Raw(sqlQuery).Find(&result)

	utils.Down(result, "福建太平摆台购买", c)
}

func (*ExportExcel) Hngx(c *gin.Context) {
	at := c.Query("at")
	if at != "sfdjwie2ji239324" {
		slog.Error("非法访问")
		c.String(200, "非法访问")
		return
	}

	type Result struct {
		Sn            string `json:"sn" tag:"卡券编号"`
		Password      string `json:"password" tag:"兑换码"`
		Status        string `json:"status" tag:"状态"`
		Active_time   string `json:"active_time" tag:"激活时间"`
		Order_no      string `json:"order_no" tag:"订单号"`
		Contact       string `json:"contact" tag:"联系人"`
		Mobile        string `json:"mobile" tag:"手机号"`
		Province      string `json:"province" tag:"省"`
		City          string `json:"city" tag:"市"`
		Area          string `json:"area" tag:"区"`
		Address       string `json:"address" tag:"地址"`
		Organ         string `json:"organ" tag:"机构"`
		Work_num      string `json:"work_num" tag:"工号"`
		Customer_info string `json:"customer_info" tag:"客户姓名"`
		Cus_mobile    string `json:"cus_mobile" tag:"客户手机"`
		Ship_name     string `json:"ship_name" tag:"快递公司"`
		Ship_no       string `json:"ship_no" tag:"快递单号"`
	}

	var result []Result

	sqlQuery := "select a.active_time,a.status,b.sn,b.password,c.order_no,c.contact,c.mobile,c.province,c.city,c.area,c.address,c.customer_info,c.ship_name,c.ship_no,c.organ,c.work_num from car_coupon a left join  car_coupon_pkg b on a.pkg_id = b.id left join car_order_photo c on a.id = c.coupon_id where a.tp_code = 'CT000564' and a.status in(1,2) and a.active_time >1704038400"
	db := model.RDB[model.MASTER]
	db.Db.Raw(sqlQuery).Find(&result)

	for k, v := range result {
		type Customer struct {
			Contact  string `json:"contact"`
			Work_num int    `json:"work_num"`
		}

		if v.Customer_info != "" {
			var tom Customer
			err := json.Unmarshal([]byte(v.Customer_info), &tom)
			if err == nil {
				result[k].Customer_info = tom.Contact
			}
		}

		status := "已激活"
		num, _ := strconv.Atoi(v.Status)
		if num == 2 {
			status = "已下单"
		}
		result[k].Status = status
		if v.Active_time != "0" {
			result[k].Active_time = utils.FormatDateByString(v.Active_time)
		}
	}

	utils.Down(result, "河南阳光个险", c)
}

func (*ExportExcel) Hnkj(c *gin.Context) {
	at := c.Query("at")
	if at != "sfdjwie2ji239324" {
		c.String(200, "非法访问")
		return
	}

	type Result struct {
		Sn            string `json:"sn" tag:"卡券编号"`
		Password      string `json:"password" tag:"兑换码"`
		Status        string `json:"status" tag:"状态"`
		Active_time   string `json:"active_time" tag:"激活时间"`
		Order_no      string `json:"order_no" tag:"订单号"`
		Contact       string `json:"contact" tag:"联系人"`
		Mobile        string `json:"mobile" tag:"手机号"`
		Province      string `json:"province" tag:"省"`
		City          string `json:"city" tag:"市"`
		Area          string `json:"area" tag:"区"`
		Address       string `json:"address" tag:"地址"`
		Organ         string `json:"organ" tag:"机构"`
		Work_num      string `json:"work_num" tag:"工号"`
		Customer_info string `json:"customer_info" tag:"客户姓名"`
		Cus_mobile    string `json:"cus_mobile" tag:"客户手机"`
		Ship_name     string `json:"ship_name" tag:"快递公司"`
		Ship_no       string `json:"ship_no" tag:"快递单号"`
	}

	var result []Result

	sqlQuery := "select a.active_time,a.status,b.sn,b.password,c.order_no,c.contact,c.mobile,c.province,c.city,c.area,c.address,c.customer_info,c.ship_name,c.ship_no,c.organ,c.work_num from car_coupon a left join  car_coupon_pkg b on a.pkg_id = b.id left join car_order_photo c on a.id = c.coupon_id where a.tp_code = 'CT001089' and a.status in(1,2) "
	db := model.RDB[model.MASTER]
	db.Db.Raw(sqlQuery).Find(&result)
	type Customer struct {
		Contact  string `json:"contact"`
		Work_num int    `json:"work_num"`
	}

	for k, v := range result {
		if v.Customer_info != "" {
			var tom Customer
			err := json.Unmarshal([]byte(v.Customer_info), &tom)
			if err == nil {
				result[k].Customer_info = tom.Contact
			}
		}

		status := "已激活"
		num, _ := strconv.Atoi(v.Status)
		if num == 2 {
			status = "已下单"
		}
		result[k].Status = status
		if v.Active_time != "0" {
			result[k].Active_time = utils.FormatDateByString(v.Active_time)
		}
	}

	utils.Down(result, "河南阳光客经", c)
}

func (*ExportExcel) Smwj(c *gin.Context) {
	at := c.Query("at")
	if at != "sfdjwie2ji239324" {
		c.String(200, "非法访问")
		return
	}

	type Result struct {
		Openid        string `json:"openid" tag:"openid"`
		Name          string `json:"name" tag:"名称"`
		Mobile        string `json:"mobile" tag:"手机号"`
		Sex           string `json:"sex" tag:"性别"`
		Question1     string `json:"question1" tag:"答题1"`
		Question2     string `json:"question2" tag:"答题2"`
		Question3     string `json:"question3" tag:"答题3"`
		Question_time string `json:"question_time" tag:"答题时间"`
		Agent_name    string `json:"agent_name" tag:"业务员姓名"`
		Agent_mobile  string `json:"agent_mobile" tag:"业务员手机"`
		Work_num      string `json:"work_num" tag:"工号"`
		Status        string `json:"status" tag:"状态"`
		C_time        string `json:"c_time" tag:"创建时间"`
	}
	sqlQuery := "select openid,name,mobile,sex,question1,question2,question3,question_time,agent_name,agent_mobile,work_num,organ,branch,agent,c_time from (select a.id, a.openid,a.work_num,a.name,a.mobile,a.sex,a.question1,a.question2,a.question3,a.question_time,a.c_time,b.mobile as agent_mobile,b.name as agent_name,b.code,c.agent,c.branch,c.organ from cs_sino_wj a ,cs_sino_cus b ,  car.car_order_photo_organ c where a.work_num = b.work_num and c.code = b.code and c.company = 21   ) as t where 1=1"

	organ, ok := c.GetQuery("organ")
	if ok {
		sqlQuery += fmt.Sprintf(" and code like '%s%%'", organ)
	}

	branch, ok := c.GetQuery("branch")
	if ok {
		sqlQuery += fmt.Sprintf(" and code like '%s%%'", branch)
	}

	agent, ok := c.GetQuery("agent")
	if ok {
		sqlQuery += fmt.Sprintf(" and code like '%s%%'", agent)
	}

	code, ok := c.GetQuery("code")
	if ok {
		sqlQuery += fmt.Sprintf(" and code like '%s%%'", code)
	}

	status, ok := c.GetQuery("status")
	if status == "1" {
		sqlQuery += " and `question_time` = 0"
	}

	if status == "2" {
		sqlQuery += " and `question_time` <> 0"
	}

	sqlQuery += " order by c_time"

	// c.String(200, sqlQuery)
	// return
	var result []Result
	db := model.RDB["db2"]
	db.Db.Raw(sqlQuery).Find(&result)

	for k, v := range result {
		t, _ := strconv.ParseInt(v.Question_time, 10, 64)
		status := "已邀约"
		if t > 0 {
			result[k].Question_time = utils.FormatDate(t)
			status = "已答题"
		}
		result[k].Status = status
		result[k].C_time = utils.FormatDateByString(v.C_time)
	}

	utils.Down(result, "问卷调查", c)
}

func (*ExportExcel) NyOrder(c *gin.Context) {
	type Result struct {
		Serial_no    string `json:"serial_no" tag:"流水号"`
		Pro_code     string `json:"pro_code" tag:"产品编码"`
		Name         string `json:"name" tag:"产品名称"`
		Thd_code     string `json:"thd_code" tag:"用户id"`
		Start_time   string `json:"start_time" tag:"权益开始时间"`
		End_time     string `json:"end_time" tag:"权益结束时间"`
		Org_code     string `json:"org_code" tag:"机构代码"`
		Org_name     string `json:"org_name" tag:"机构名称"`
		Status       string `json:"status" tag:"状态"`
		C_time      string `json:"c_time" tag:"创建时间"`
	}

	var result []Result
	sqlQuery := "select a.serial_no as 'serial_no',a.pro_code as 'pro_code',b.`name` as 'name',a.thd_code as 'thd_code',a.start_time as 'start_time',a.end_time as 'end_time',org_code as 'org_code',org_name as 'org_name',case a.status when 1 then '已激活' when 2 then '已使用'when 3 then '已激活' when -1 then '已撤销' end as 'status',FROM_UNIXTIME(c_time) as 'c_time'  from car_nongyin_coupon_list a LEFT JOIN car_api_product b on a.pro_code = b.code  "

	db := model.RDB[model.MASTER]
	db.Db.Raw(sqlQuery).Find(&result)

	utils.Down(result, "农业人寿客户生日礼", c)
}
