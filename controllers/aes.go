package controllers

import (
	"ally/utils/crypto"
	"encoding/json"
	"log"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Aes(c *gin.Context) {
	json := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&json)
	log.Printf("%v", &json)
	c.JSON(200, gin.H{
		"name":     json["name"],
		"password": json["password"],
	})
}

func Encrypt(c *gin.Context) {
	bindParams := map[string]interface{}{}
	c.ShouldBindBodyWith(&bindParams, binding.JSON)
	str, err := json.Marshal(bindParams)
	if err != nil {
		log.Println(err)
	}

	slog.Info("str2", str)

	key := []byte("ABCDABCDABCDABCD")

	encrypt := crypto.AesEncryptECB(str, key)
	if err != nil {
		slog.Warn("加密失败", err)
	}

	slog.Info("加密内容", encrypt)
	c.JSON(200, gin.H{"name": encrypt})

}
