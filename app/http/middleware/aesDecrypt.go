package middleware

import (
	"ally/pkg/config"
	"ally/utils/crypto"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func AesDecrypt() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := c.GetRawData()
		var body map[string]string
		_ = json.Unmarshal(data, &body)
		key := []byte(config.GetString("app.aes-128-ecb"))
		str := string(body["encrypt"])

		decrypt := crypto.AesDecryptECB(str, key)

		if len(decrypt) == 0 {
			c.JSON(200, gin.H{
				"code": 500,
				"msg":  "参数解密失败",
			})
			c.Abort()
		}
		c.Set("decrypt", decrypt)
		c.Set("encrypt", str)
		c.Next()
	}
}
