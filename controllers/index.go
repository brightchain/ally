package controllers

import (
	"ally/model"
	"ally/utils"
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/plugin/dbresolver"
)

func Index(c *gin.Context) {
	c.String(200, "测试页面")
}

func Hngx(c *gin.Context) {
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

	model.DB.Raw(sqlQuery).Find(&result)

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

func Hnkj(c *gin.Context) {
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

	model.DB.Raw(sqlQuery).Find(&result)
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

func Smwj(c *gin.Context) {
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
	sqlQuery := "select openid,name,mobile,sex,question1,question2,question3,question_time,agent_name,agent_mobile,work_num,organ,branch,agent,c_time from (select a.id, a.openid,a.work_num,a.name,a.mobile,a.sex,a.question1,a.question2,a.question3,a.question_time,a.c_time,b.mobile as agent_mobile,b.name as agent_name,b.code,c.agent,c.branch,c.organ from cs_sino_wj a ,cs_sino_cus b ,  car.car_order_photo_organ c where a.work_num = b.work_num and c.code = b.code and c.company = 21 ) as t where 1=1"

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

	model.DB.Clauses(dbresolver.Use("custom")).Raw(sqlQuery).Find(&result)

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
