package api

import (
	"ally/model"
	"ally/utils"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

	"github.com/gin-gonic/gin"
)

func PhotoOrder(c *gin.Context) {
	decrypt, _ := c.Get("decrypt")
	var param map[string]string
	str, _ := decrypt.([]byte)
	_ = json.Unmarshal(str, &param)
	slog.Info("优化", param)

	var values []interface{}
	where := " 1=1"

	for k, v := range param {
		tmp := strings.Split(v, " ")
		where += fmt.Sprintf(" and %s %s ?", k, tmp[0])
		vl := strings.Split(tmp[1], ",")
		values = append(values, vl)
	}
	slog.Info("where:", where)
	slog.Info("values:", values)
	db := model.RDBs[model.MASTER]
	var orders []model.PhotoCy
	db.Db.Model(&model.CarOrderPhoto{}).Where(where, values...).Find(&orders)

	c.JSON(200, orders)
}

func Zip(c *gin.Context) {
	// var abPath string
	// _, filename, _, ok := runtime.Caller(0)
	// if ok {
	// 	abPath = path.Dir(filename)
	// }
	fileName := "1.zip"

	zipWriter, err := utils.ZipFiles(fileName)
	if err != nil {
		slog.Warn("创建压缩文件失败", err)
	}
	defer zipWriter.Close()
	err = utils.AddFileToZip(zipWriter, "1.xlsx")
	if err != nil {
		slog.Warn("压缩失败")
	}

	c.String(200, "/public/storage/1.zip")

}
