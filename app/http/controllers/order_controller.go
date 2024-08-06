package controllers

import (
	"ally/pkg/model"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type PayOrder struct{}

func (*PayOrder) GetOrderProduct(c *gin.Context) {
	f, err := excelize.OpenFile("order.xlsx")
	if err != nil {
		slog.Warn("读取失败！", err)
		c.String(http.StatusOK, "读取失败")
		return
	}

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		slog.Warn("读取失败sheet1！", err)
		c.String(http.StatusOK, "读取失败")
		return
	}

	db := model.RDB[model.MASTER]
	for i, row := range rows {
		if i == 0 { // 跳过标题行
			continue
		}

		// 查询数据库
		var result map[string]interface{}
		cellValue := ""
		cellValue = row[10]

		fmt.Printf("行数%v值%v\n", i, cellValue)
		if cellValue == "" {
			continue
		}

		orderNo := strings.ReplaceAll(cellValue, "`", "")
		orderType := orderNo[:2]
		
		if orderType == "YZ" {
			err := db.Db.Raw("SELECT * FROM car_shop_yz_order_v WHERE order_no = ?", orderNo).Scan(&result).Error
			if err != nil {
				log.Printf("Error querying database for ID %s: %v", orderNo, err)
				continue
			}
			f.SetCellValue("Sheet1", fmt.Sprintf("M%d", i+1), result["pro_name"])
			f.SetCellValue("Sheet1", fmt.Sprintf("N%d", i+1), 1)
			f.SetCellValue("Sheet1", fmt.Sprintf("O%d", i+1), result["total_amount"])
			f.SetCellValue("Sheet1", fmt.Sprintf("P%d", i+1), result["pay_amount"])
		} else if orderType == "DD" || orderType == "DA" {
			err := db.Db.Raw("SELECT * FROM car_shop_dadi_order WHERE order_no = ?", orderNo).Scan(&result).Error
			if err != nil {
				log.Printf("Error querying database for ID %s: %v", orderNo, err)
				continue
			}
			f.SetCellValue("Sheet1", fmt.Sprintf("M%d", i+1), result["product_name"])
			f.SetCellValue("Sheet1", fmt.Sprintf("N%d", i+1), result["amount"])
			f.SetCellValue("Sheet1", fmt.Sprintf("O%d", i+1), result["order_amount"])
			f.SetCellValue("Sheet1", fmt.Sprintf("P%d", i+1), result["pay_amount"])
		} else if orderType == "VC" {
			err := db.Db.Raw("SELECT * FROM car_vcard_order WHERE order_no = ?", orderNo).Scan(&result).Error
			if err != nil {
				log.Printf("Error querying database for ID %s: %v", orderNo, err)
				continue
			}
			f.SetCellValue("Sheet1", fmt.Sprintf("M%d", i+1), result["product_name"])
			f.SetCellValue("Sheet1", fmt.Sprintf("N%d", i+1), 1)
			f.SetCellValue("Sheet1", fmt.Sprintf("O%d", i+1), result["amount"])
			f.SetCellValue("Sheet1", fmt.Sprintf("P%d", i+1), result["pay_amount"])
		} else if orderType == "GP" {
			err := db.Db.Raw("SELECT *,case pro_id when 'PA001' then '平安专版水晶相框' when 'PA002' then '平安专版私定相册' when 'GS001' then '中国人寿专版水晶相框' when 'GS002' then '中国人寿专版时光相册' end as 'product_name' FROM car_order_gdpa WHERE order_no = ?", orderNo).Scan(&result).Error
			if err != nil {
				log.Printf("Error querying database for ID %s: %v", orderNo, err)
				continue
			}
			f.SetCellValue("Sheet1", fmt.Sprintf("M%d", i+1), result["product_name"])
			f.SetCellValue("Sheet1", fmt.Sprintf("N%d", i+1), result["num"])
			f.SetCellValue("Sheet1", fmt.Sprintf("O%d", i+1), result["order_amount"])
			f.SetCellValue("Sheet1", fmt.Sprintf("P%d", i+1), result["pay_amount"])
		}
		// 4. 更新Excel文档
		// 假设我们要更新第三列

	}

	// 5. 保存更新后的Excel文档
	if err := f.SaveAs("output.xlsx"); err != nil {
		log.Fatal(err)
	}
	c.String(http.StatusOK, "OK")
}
