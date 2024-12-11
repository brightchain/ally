package controllers

import (
	"fmt"
	"h5/pkg/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type Activity struct {
}

// 定义数据模型
type CsYdUser struct {
	Activity    string `gorm:"column:activity"`
	Contact     string `gorm:"column:contact"`
	Mobile      string `gorm:"column:mobile"`
	Address     string `gorm:"column:address"`
	Prize       int    `gorm:"column:prize"`
	CouponID    int    `gorm:"column:coupon_id"`
	LotteryTime int64  `gorm:"column:lottery_time"`
	OrgCode     int    `gorm:"column:org_code"`
	OrderNo     string `gorm:"column:order_no"`
	AgtWorkNum  string `gorm:"column:agt_work_num"`
}
func (a *CsYdUser) TableName() string {
	return "cs_yd_user"
}
type TmpYd struct {
	Mobile string `gorm:"column:mobile"`
	OrderNo string `gorm:"column:order_no"`
}

func (a *TmpYd) TableName() string {
	return "tmp_yd"
}


type CarCoupon struct {
	UserID int    `gorm:"column:user_id"`
	Mobile string `gorm:"column:mobile"`
	Status int    `gorm:"column:status"`
}

func (a *CarCoupon) TableName() string {
	return "car.car_coupon"
}

// 订单导出结构体
type ExportOrder struct {
	AgentName       string `gorm:"column:name"`
	AgentMobile     string `gorm:"column:mobile"`
	AgentCode       string `gorm:"column:code"`
	AgentOrgCode    string `gorm:"column:branch"`
	BranchName      string `gorm:"column:branch_name"`
	CustomerName    string `gorm:"column:customer_name"`
	CustomerMobile  string `gorm:"column:customer_mobile"`
	CustomerAddress string `gorm:"column:customer_address"`
	PrizeName       string `gorm:"column:prize_name"`
	PrizeID         int    `gorm:"column:prize_id"`
	CouponID        int    `gorm:"column:coupon_id"`
	LotteryTime     string `gorm:"column:lottery_time"`
	OrderNo         string `gorm:"column:order_no"`
	OrderTime       string `gorm:"column:order_time"`
}

func (a *Activity) UserReset(c *gin.Context) {
	// 初始化数据库连接
	db := model.RDB["db1"].Db
	// 开启事务
	tx := db.Begin()
	if tx.Error != nil {
		c.String(http.StatusOK, "开启事务失败")
		return 
	}

	// 查询订单数据
	orders, err := a.fetchOrders(tx)
	if err != nil {
		c.String(http.StatusOK, fmt.Sprintf("获取订单失败: %v", err))
		return
	}

	// 如果没有订单，直接返回
	if len(orders) == 0 {
		c.String(http.StatusOK, "没有找到符合条件的订单")	
		return
	}


	// 导出Excel
	err = exportToExcel(orders)
	if err != nil {
		tx.Rollback()
		c.String(http.StatusOK, fmt.Sprintf("导出Excel失败: %v", err))
		return
	}

	// 更新订单状态
	err = updateOrderStatus(tx, orders)
	if err != nil {
		tx.Rollback()
		c.String(200, fmt.Sprintf("更新订单状态失败: %v", err))
		return
	}

	// 更新优惠券状态
	err = updateCouponStatus(tx, orders)
	if err != nil {
		tx.Rollback()
		c.String(200, fmt.Sprintf("更新优惠券状态失败: %v", err))
		return
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		c.String(200, fmt.Sprintf("事务提交失败: %v", err))
		return
	}

	c.String(http.StatusOK, "订单导出和状态更新完成")	
}

// 获取订单数据
func (a *Activity) fetchOrders(tx *gorm.DB) ([]ExportOrder, error) {
	var exportOrders []ExportOrder
	// 首先获取tmp_yd表中的手机号
	var mobileList []TmpYd
	err := tx.Model(&TmpYd{}).Scan(&mobileList).Error
	if err != nil || len(mobileList) == 0 {
		return nil, fmt.Errorf("未找到有效的手机号列表: %v", err)
	}
	var mobile []string
	for _, v := range mobileList {
		mobile = append(mobile, v.Mobile)
	}
	fmt.Println(mobileList)
    sql := `SELECT 
				b.name, b.mobile, b.code, b.agent, b.branch,b.branch_name, 
				a.contact customer_name, a.mobile customer_mobile, a.address customer_address, 
				CASE a.prize 
					WHEN 1 THEN '摆台' 
					WHEN 2 THEN '挂轴' 
					WHEN 3 THEN '桌垫' 
					WHEN 4 THEN '台历' 
				END prize_name, 
				a.prize prize_id, a.coupon_id, 
				FROM_UNIXTIME(a.lottery_time) lottery_time
				FROM cs_yd_user a 
				LEFT JOIN cs_yd_custom_v b ON a.activity = b.activity AND a.agt_work_num = b.code 
				WHERE a.org_code <> 8613 AND a.order_no = '' AND a.prize <> 0 
				AND a.mobile IN (?)
				`
	// 复杂的连表查询
	err1 := tx.Raw(sql, mobile).Scan(&exportOrders).Error

	return exportOrders, err1
}

// 更新订单状态
func updateOrderStatus(tx *gorm.DB, orders []ExportOrder) error {
	// 提取手机号
	var mobiles []string
	for _, order := range orders {
		mobiles = append(mobiles, order.CustomerMobile)
	}

    now := time.Now()

	// 批量更新
	return tx.Model(&CsYdUser{}).
		Where("mobile IN ?", mobiles).
		Where("order_no = ''").
		Where("org_code <> 8613").
		Where("prize <> 0").
		Updates(map[string]interface{}{
			"coupon_id":   0,
			"prize" :  0,
			"u_time": now.Unix(),
		}).Error
}

// 更新优惠券状态
func updateCouponStatus(tx *gorm.DB, orders []ExportOrder) error {
	// 提取优惠券ID
	var couponIDs []int
	for _, order := range orders {
		couponIDs = append(couponIDs, order.CouponID)
	}

	// 批量更新
	return tx.Model(&CarCoupon{}).
		Where("id IN ?", couponIDs).
		Updates(map[string]interface{}{
			"user_id":   0,
			"mobile":    "",
			"status":    0,
		}).Error
}



// 导出到Excel（保持不变）
func exportToExcel(orders []ExportOrder) error {
	// 之前的Excel导出代码保持不变
	f := excelize.NewFile()
	sheetName := "数据明细"
	index, _ := f.NewSheet(sheetName)
	f.SetActiveSheet(index)

	// 设置表头
	headers := []string{
		"代理人姓名", "代理人手机号", "工号",  "机构名称",
		"客户姓名", "手机号", "客户地址", "奖品名称",
		"奖品ID", "权益ID", "抽奖时间",
	}
	for col, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(col+1, 1)
		f.SetCellValue(sheetName, cell, header)
	}

	// 写入数据
	for row, order := range orders {
		data := []interface{}{
			order.AgentName, order.AgentMobile, order.AgentCode,
			 order.BranchName,
			order.CustomerName, order.CustomerMobile, order.CustomerAddress,
			order.OrderNo, order.OrderTime,
		}
		for col, value := range data {
			cell, _ := excelize.CoordinatesToCellName(col+1, row+2)
			f.SetCellValue(sheetName, cell, value)
		}
	}

	// 保存文件
	filename := fmt.Sprintf("英大订单取消%s.xlsx", time.Now().Format("20060102"))
	return f.SaveAs(filename)
}

func (a *Activity) CancelOrder(c *gin.Context){
	// 初始化数据库连接
	db := model.RDB["db1"].Db
	// 开启事务
	tx := db.Begin()
	if tx.Error != nil {
		c.String(http.StatusOK, "开启事务失败")
		return 
	}

	// 查询订单数据
	orders, err := a.getCanelOrders(tx)
	if err != nil {
		c.String(http.StatusOK, fmt.Sprintf("获取订单失败: %v", err))
		return
	}

	// 如果没有订单，直接返回
	if len(orders) == 0 {
		c.String(http.StatusOK, "没有找到符合条件的订单")	
		return
	}

	// 导出Excel
	err = exportOrderToExcel(orders)
	if err != nil {
		tx.Rollback()
		c.String(http.StatusOK, fmt.Sprintf("导出Excel失败: %v", err))
		return
	}

	var couponIDs []int
	for _, order := range orders {
		couponIDs = append(couponIDs, order.CouponID)
	}

	now := time.Now()

	// 批量更新
	err = tx.Model(&CsYdUser{}).
		Where("coupon_id IN ?", couponIDs).
		Updates(map[string]interface{}{
			"order_no":   "",
			"order_time" :  0,
			"u_time": now.Unix(),
		}).Error

	if err != nil {
		tx.Rollback()
		c.String(http.StatusOK, fmt.Sprintf("更新订单失败: %v", err))
		return
	}

	// 更新优惠券状态
	err = tx.Model(&CarCoupon{}).
		Where("id IN ?", couponIDs).
		Updates(map[string]interface{}{
			"status":    1,
		}).Error

	if err != nil {
		tx.Rollback()
		c.String(http.StatusOK, fmt.Sprintf("更新优惠券状态失败: %v", err))
		return
	}

	// 提交事务
	tx.Commit()
	c.String(http.StatusOK, "订单取消成功")
	
}

// 获取订单数据
func (a *Activity) getCanelOrders(tx *gorm.DB) ([]ExportOrder, error) {
	var exportOrders []ExportOrder
	// 首先获取tmp_yd表中的手机号
	var mobileList []TmpYd
	err := tx.Model(&TmpYd{}).Scan(&mobileList).Error
	if err != nil || len(mobileList) == 0 {
		return nil, fmt.Errorf("未找到有效的手机号列表: %v", err)
	}
	var orders []string
	for _, v := range mobileList {
		orders = append(orders, v.OrderNo)
	}
	fmt.Println(orders)
    sql := `select 
				b.name,
				b.mobile,
				b.code,
				b.agent_name branch_name,
				a.contact customer_name,
				a.mobile customer_mobile,
				a.address customer_address,
				a.order_no,
				FROM_UNIXTIME(order_time) order_time
				from cs_yd_user a LEFT JOIN cs_yd_custom_v b on a.activity=25 and a.agt_work_num = b.code and b.activity =25 
				where a.order_no in (?) 
				`
	// 复杂的连表查询
	err1 := tx.Raw(sql, orders).Scan(&exportOrders).Error

	return exportOrders, err1
}

func exportOrderToExcel(orders []ExportOrder) error {
	// 之前的Excel导出代码保持不变
	f := excelize.NewFile()
	sheetName := "数据明细"
	index, _ := f.NewSheet(sheetName)
	f.SetActiveSheet(index)

	// 设置表头
	headers := []string{
		"代理人", "代理人手机号", "工号", "机构名称",
		"客户姓名", "客户手机号", "客户地址", "订单编号",
		"下单时间",
	}
	for col, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(col+1, 1)
		f.SetCellValue(sheetName, cell, header)
	}

	// 写入数据
	for row, order := range orders {
		data := []interface{}{
			order.AgentName, order.AgentMobile,
			order.AgentOrgCode, order.BranchName,
			order.CustomerName, order.CustomerMobile, order.CustomerAddress,
			order.PrizeName, order.PrizeID, order.CouponID,
			order.LotteryTime,
		}
		for col, value := range data {
			cell, _ := excelize.CoordinatesToCellName(col+1, row+2)
			f.SetCellValue(sheetName, cell, value)
		}
	}

	// 保存文件
	filename := fmt.Sprintf("order_export_%s.xlsx", time.Now().Format("20060102_150405"))
	return f.SaveAs(filename)
}

