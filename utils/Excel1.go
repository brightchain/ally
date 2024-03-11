package utils

import (
	"fmt"
	"net/url"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func Save(params []string, data []interface{}, col []string, filename string, ctx *gin.Context) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 创建一个工作表
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	var word = 'A'
	for _, key := range params {

		line := fmt.Sprintf("%c%v", word, 1)
		f.SetCellValue("Sheet1", line, key)
		word++
	}
	rowNum := 1 //数据开始行数
	for _, v := range data {
		t := reflect.TypeOf(v)
		value := reflect.ValueOf(v)
		row := make([]interface{}, 0)
		for l := 0; l < t.NumField(); l++ {
			val := value.Field(l).Interface()
			row = append(row, val)
		}
		rowNum++
		f.SetSheetRow("Sheet1", fmt.Sprintf("A%d", rowNum), &row)

	}

	f.SetCellValue("Sheet1", "B2", 100)
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 根据指定路径保存文件
	if err := f.SaveAs(filename + ".xlsx"); err != nil {
		fmt.Println(err)
	}
	return
}

func Download(titleList []string, data []map[string]string, col []string, filename string, c *gin.Context) {
	sheetName := "Sheet1"
	f := excelize.NewFile()
	f.SetSheetName("Sheet1", sheetName)
	header := make([]string, 0)
	for _, v := range titleList {
		header = append(header, v)
	}

	_ = f.SetSheetRow(sheetName, "A1", &header)
	_ = f.SetRowHeight(sheetName, 1, 30)

	rowNum := 1 //数据开始行数
	for _, value := range data {
		row := make([]interface{}, 0)
		for _, v := range col {
			if val, ok := value[v]; ok {
				row = append(row, val)
			}
		}
		rowNum++
		f.SetSheetRow("Sheet1", fmt.Sprintf("A%d", rowNum), &row)
	}

	disposition := fmt.Sprintf("attachment; filename=%s-%s.xlsx", url.QueryEscape(filename), time.Now().Format("2006-01-02"))
	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	c.Writer.Header().Set("Content-Disposition", disposition)
	c.Writer.Header().Set("Content-Transfer-Encoding", "binary")
	c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")
	_ = f.Write(c.Writer)
}
