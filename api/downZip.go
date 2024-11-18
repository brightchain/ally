package api

import (
	"archive/zip"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"h5/app/http/models"
	"h5/pkg/goredis"
	"h5/pkg/model"
	"h5/utils"
	"io"
	"log/slog"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
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
	n, err := goredis.Client.Get(c, md5str).Result()
	if err != redis.Nil {
		c.String(200, n)
		return
	}

	var param map[string]string
	str, _ := decrypt.([]byte)
	_ = json.Unmarshal(str, &param)
	slog.Info("优化", param)

	db := model.RDB[model.MASTER]
	tx := db.Db.Model(&models.CarOrderPhoto{})
	for k, v := range param {
		//tmp := strings.Split(k, " ")
		//where := fmt.Sprintf("%s %s ?", tmp[0], tmp[1])
		where := fmt.Sprintf("%s ?", k)
		vl := strings.Split(v, ",")
		tx.Where(where, vl)
	}
	var orders []models.PhotoCy

	tx.Order("id asc").Find(&orders)
	data := make([]models.PhotoOrder, len(orders))
	for k, v := range orders {
		data[k] = models.FormatDataCy(v)
	}

	path := "./storage/app/public"
	currentTime := time.Now()
	r := rand.New(rand.NewSource(currentTime.UnixNano()))
	name := fmt.Sprintf("%s%d", currentTime.Format("20060102150405"), r.Int63n(1000))
	fileName := path + "/" + name + ".xlsx"
	utils.SaveFile(data, fileName)

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
	
	for _, order := range data {
		name := strings.ToLower(order.OrderNo)
		fileName := models.FilePath + "/" + name + "/" + order.ProId + ".jpeg"
		zName := order.Uid + " " + order.Contact + " " + name + "+" + order.ProName + ".jpeg"
		if order.Remark != "" {
			zName = order.Remark + " " + name + "+" + order.ProName + ".jpeg"
		}

		err = utils.AddFileToZip(zipWriter, fileName, zName)
		if err != nil {
			slog.Warn("压缩失败", err)
			break
		}
	}
	if err != nil {
		newZipFile.Close()
		os.Remove(zipName)
		os.Remove(fileName)
		c.String(200, "zip 压缩失败")
		return
	}

	// upData := map[string]interface{}{
	// 	"status": 1,
	// 	"u_time": currentTime.Unix(),
	// }

	// result := tx.Updates(upData)
	// if result.Error != nil {
	// 	slog.Error("更新失败", result.Error)
	// 	c.String(200, "更新失败")
	// 	return
	// }

	goredis.Client.Set(c, md5str, zipName, 7*24*time.Hour)

	c.String(200, zipName)
}

func MouseOrderDown(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit := ""
	if page > 0 {
		limit = fmt.Sprintf("%d,100", (page-1)*100)
	}

	// 获取其他参数
	status, _ := strconv.Atoi(c.DefaultQuery("status", "1"))
	sdate := c.DefaultQuery("sdate", "")
	edate := c.DefaultQuery("edate", "")
	name := c.DefaultQuery("name", "")
	sign := c.DefaultQuery("sign", "")

	if len(name) <= 0 {
		c.String(200, "缺少参数！")
		return
	}
	h := md5.New()
	io.WriteString(h, name+"13687391951")
	if fmt.Sprintf("%x", h.Sum(nil))!= sign {
		c.String(200, "sign签名失败！")
		return
	}

	// 构建 WHERE 条件
	where := fmt.Sprintf("`status` = %d and company = ", status)

	if sdate != "" && edate != "" {
		where += fmt.Sprintf(" AND `c_time` >= '%s' AND `c_time` < '%s'", sdate, edate)
	}

	type Result struct {
		Order_no        string `json:"order_no" tag:"订单编号"`
		Ship_name       string `json:"ship_name" tag:"物流公司"`
		Ship_no         string `json:"ship_no" tag:"物流单号"`
		Contact        string `json:"contact" tag:"收货人"`
		Mobile         string `json:"mobile" tag:"收货手机"`
		Customer_info   string `json:"customer_info" tag:"客户姓名"`
		Province         string `json:"province" tag:"省"`
		City         string `json:"city" tag:"市"`
		Area         string `json:"area" tag:"区"`
		Address         string `json:"address" tag:"详细地址"`
		Remark string `json:"remark" tag:"备注"`
		C_time          string `json:"c_time" tag:"创建时间"`
	}

	var result []Result
	db := model.RDB[model.MASTER]
	sqlQuery := fmt.Sprintf("select order_no,contact,mobile,province,city,area,address,customer_info,c_time from car_order_shirt where %s limit %s",where,limit)
	db.Db.Raw(sqlQuery).Find(&result)
	utils.Down(result, "摆台订单", c)
	
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
	goredis.Client.Set(c, "key", "value", 60*time.Second)
	val2, err := goredis.Client.Get(c, "key").Result()
	if err == redis.Nil {
		fmt.Println("[ERROR] - Key [name] not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		panic(err)
	}
	c.String(200, val2)
}
