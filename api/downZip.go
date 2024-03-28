package api

import (
	"ally/model"
	"ally/utils"
	"archive/zip"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
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

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
	path := exPath + "/storage/app/public"
	fileName := path + "/1.zip"
	newZipFile, err := os.Create(fileName)
	if err != nil {
		slog.Error("zip create fail", err)
		return
	}

	defer newZipFile.Close()
	zipWriter := zip.NewWriter(newZipFile)

	err = utils.AddFileToZip(zipWriter, path+"/1.xlsx", "img/")
	if err != nil {
		slog.Warn("压缩失败")
	}
	err = utils.AddFileToZip(zipWriter, path+"/1.png", "excel/")
	if err != nil {
		slog.Warn("压缩失败")
	}

	defer zipWriter.Close()
	c.String(200, "/public/storage/1.zip")

}
