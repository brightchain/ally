package api

import (
	"ally/config"
	"ally/model"
	"ally/utils"
	"archive/zip"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func PhotoOrderCy(c *gin.Context) {
	decrypt, _ := c.Get("decrypt")
	encrypt, _ := c.Get("encrypt")
	has := md5.Sum([]byte(encrypt.(string)))
	md5str := fmt.Sprintf("%x", has)
	n, err := config.RedisClient.Get(c, md5str).Result()
	if err != redis.Nil {
		c.String(200, n)
		return
	}

	var param map[string]string
	str, _ := decrypt.([]byte)
	_ = json.Unmarshal(str, &param)
	slog.Info("优化", param)

	db := model.RDBs[model.MASTER]
	tx := db.Db.Model(&model.CarOrderPhoto{})
	for k, v := range param {
		//tmp := strings.Split(k, " ")
		//where := fmt.Sprintf("%s %s ?", tmp[0], tmp[1])
		where := fmt.Sprintf("%s ?", k)
		vl := strings.Split(v, ",")
		tx.Where(where, vl)
	}
	var orders []model.PhotoCy
	tx.Find(&orders)
	if len(orders) == 0 {
		c.String(200, "数据不存在！")
		return
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
		c.String(200, "zip 创建失败")
		return
	}
	defer newZipFile.Close()
	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()
	err = utils.AddFileToZip(zipWriter, fileName, "")
	if err != nil {
		slog.Warn("压缩失败")
		c.String(200, "zip 压缩失败")
		return
	}
	for _, order := range orders {
		name := strings.ToLower(order.OrderNo)
		fileName := model.FilePath + "/" + name + "/" + order.ProId + ".jpeg"
		zName := order.Uid + " " + order.Contact + " " + name + "+" + order.ProName + ".jpeg"
		err = utils.AddFileToZip(zipWriter, fileName, zName)
		if err != nil {
			slog.Warn("压缩失败", err)
			break
		}
	}
	if err != nil {
		newZipFile.Close()
		os.Remove(zipName)
		c.String(200, "zip 压缩失败")
		return
	}

	upData := map[string]interface{}{
		"status": 1,
		"u_time": currentTime.Unix(),
	}

	result := tx.Updates(upData)
	if result.Error != nil {
		slog.Error("更新失败", result.Error)
		c.String(200, "更新失败")
		return
	}

	config.RedisClient.Set(c, md5str, zipName, 7*24*time.Hour)

	c.String(200, zipName)
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

func Redis(c *gin.Context) {
	config.RedisClient.Set(c, "key", "value", 60*time.Second)
	val2, err := config.RedisClient.Get(c, "key").Result()
	if err == redis.Nil {
		fmt.Println("[ERROR] - Key [name] not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		panic(err)
	}
	c.String(200, val2)
}
