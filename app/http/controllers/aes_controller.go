package controllers

import (
	"ally/pkg/config"
	"ally/pkg/model"
	"ally/utils/crypto"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AesEcb struct {
}

func (a *AesEcb) Aes(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	log.Printf("%v", &json)
	c.JSON(200, gin.H{
		"name":     json["name"],
		"password": json["password"],
	})
}

func (a *AesEcb) Encrypt(c *gin.Context) {
	bindParams := map[string]interface{}{}
	c.ShouldBindBodyWith(&bindParams, binding.JSON)
	str, err := json.Marshal(bindParams)
	if err != nil {
		log.Println(err)
	}

	slog.Info("str2", str)

	key := []byte(config.Data.GetString("crypto.aes-128-ecb"))

	encrypt := crypto.AesEncryptECB(str, key)
	if err != nil {
		slog.Warn("加密失败", err)
	}

	slog.Info("加密内容", encrypt)
	c.JSON(200, gin.H{"encrypt": encrypt})

}

func (a *AesEcb) Down(c *gin.Context) {
	data, _ := c.GetRawData()
	var body map[string]string
	_ = json.Unmarshal(data, &body)
	key := []byte(config.Data.GetString("crypto.aes-128-ecb"))
	str := string(body["encrypt"])

	decrypt := crypto.AesDecryptECB(str, key)
	var par map[string]string
	_ = json.Unmarshal(decrypt, &par)
	slog.Info("解密内容", par)
	db := model.RDB[model.MASTER]
	where := " 1=1 "
	var values []interface{}
	for k, v := range par {
		tm := strings.Split(v, " ")
		where += fmt.Sprintf(" and %s %s ?", k, tm[0])
		values = append(values, tm[1])
	}

	type Result struct {
		Mobile    string `json:"mobile" tag:"手机号"`
		Contact   string `json:"contact" tag:"联系人"`
		Work_num  string `json:"work_num" tag:"工号"`
		Order_no  string `json:"order_no" tag:"订单号"`
		Province  string `json:"province" tag:"省"`
		City      string `json:"city" tag:"市"`
		Area      string `json:"area" tag:"区"`
		Address   string `json:"address" tag:"地址"`
		Ship_name string `json:"ship_name" tag:"快递公司"`
		Ship_no   string `json:"ship_no" tag:"快递单号"`
		C_time    string `json:"c_time" tag:"创建时间"`
	}
	var order []Result
	db.Db.Table("car_order_photo").Where(where, values...).Find(&order)

	c.JSON(200, order)

}
