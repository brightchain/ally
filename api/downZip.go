package api

import (
	"ally/model"
	"ally/utils"
	"archive/zip"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

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
	db := model.RDBs[model.MASTER]
	var orders []model.PhotoCy
	db.Db.Model(&model.CarOrderPhoto{}).Where(where, values...).Find(&orders)
	if len(orders) == 0 {
		c.String(200, "数据不存在！")
	}
	path := "./storage/app/public"
	currentTime := time.Now()
	r := rand.New(rand.NewSource(currentTime.UnixNano()))
	name := fmt.Sprintf("%s%d", currentTime.Format("20060102150405"), r.Int63n(1000))
	fileName := path + "/" + name + ".xlsx"
	utils.SaveFile(orders, fileName)
	zipName := path + "/" + name + ".zip"
	newZipFile, err := os.Create(zipName)
	if err != nil {
		slog.Error("zip create fail", err)
		return
	}
	defer newZipFile.Close()
	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()
	err = utils.AddFileToZip(zipWriter, fileName, "")
	if err != nil {
		slog.Warn("压缩失败")
	}
	for _, order := range orders {
		name := strings.ToLower(order.OrderNo)
		fileName := model.FilePath + "/" + name + "/" + order.ProId + ".jpeg"
		zName := order.Uid + " " + order.Contact + " " + name + " " + order.ProName + ".jpeg"
		err = utils.AddFileToZip(zipWriter, fileName, zName)
		if err != nil {
			slog.Warn("压缩失败", err)
		}
	}
	c.String(200, "ok")
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
