package controllers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"h5/pkg/config"
	"h5/pkg/model"
	"h5/utils/crypto"
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

	key := []byte(config.GetString("app.aes-128-ecb"))

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
	key := []byte(config.GetString("app.aes-128-ecb"))
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

func (a *AesEcb) RsaDecrypt(c *gin.Context) {
	privateKeyPEM := `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCTNJ+n5YqgOyMS
ofkJjCVeuqKfF3k7PHHFhKPTOMkqYhjj4G5lCEt9o++pnnbO49cY9nmKc9Ybq+Oa
7k5NgCbSXunqPA+6myhFNbW/eeGGU04dSCUQvjd/sXS12MG5mqae6OygIiZqsBAn
T5bnNfGTa00e/4e+p35b92V8+JOUS3kFJehbGYybAOoXCdoWiMXP5u7Q27zObzTD
3dTTcptMTTO+22b3pR/+8WTHWx+dN2ME49RNoCqopRlEpdsDWds11Gxnq7KHnE0n
idAf2MYaREk8HtPpl1rqJgSyNVQHWXoQnQQsIO2mSyxrSFLNeP7yAuke6DEApc6/
7Vvg21MtAgMBAAECggEAKN7vhRpCRwKkVkQKdRAoQAjppepKipvZqtGM+tRFZjMe
XgauH/cRnqypmhqZOhAgZJnqXPqUc9Jhu1529yWhob8gixxp8ZGquKyU7bjVWQpA
Ifkp8WAe4KFQmjy4bOP3Zx+cs0lICU8g7Qk4CLH9hMTCAN1JvzGZ78bcsroBn6Z1
37ZBtF4X9MUtvKNDC+UV3XUt+S/wJW1EoibmIxvgepf4k9S/vWPidFK9EIpycJYb
AsPk1MTQ52rJgxKavPUmc4XyPAyzM2FPeTv486HjIjWETAs4Uc5LBc+nncEocQjm
PHu0YxlTIa1s6+rKg6Zu+JONZRvnDxnQuHKBj7u5gQKBgQD290BbbGj3AoCbCZb1
8VOq7Z6oCR4OX3TDiELv/PrU+fpWUCffzfx+1TtPztvaoIbzbCZf4KV6QAw1+ELC
1Wnxa6FJNY8MuMs6oX9KifZA/XBOLPSpfNuEY6CL8kE1+LoL/QHc17XKQcbJWAJI
IIQTyCFYBZX/iRS3MdWDgEbmmQKBgQCYlyb6zX1eWTscNawSePHGJxCvBsywNxOD
MWz38YIiosrNDxHL5uuoFPwgStYBW1aDKyEuhX4R/PIEIzD7tbukHE0OWHNvnOSA
9TG3K5h6XM4JVNdg4rU99aEG4iiXlSj6ooUjLqIWsDNef38gs5V+ZVRXobdRq8dY
PaETCq8xtQKBgQCNAPpDuI/tiAY5fq4sWc71ZpoRQrNgbWvTP/dH4l6sPtWJdlZt
18Pz/nbQyECKgODMFtGKKE8Sj4LpRYeIiiTk7Bi4HdZA4zlVjmOJAWASFyGJ4O0H
3/vNiLxfar+EuzAuLBRmLNrkUc3Xo2IaIEaIrNks0nOj5HZpahCH9jOE8QKBgFaR
CfAnPASWos4yNNiV/LPp3bEuLlmaJVu8YpGXVbjImj0TW4lODEti/FZlnltOshng
EgcOfKM/2R03ycZDJ5zG4YBN9c9QNuJiOD4uYWap18m7dCTm+OOZwizhiR3V5VWr
ddSr1BEDDWGC+2BWAW2fluXQPOv8hC8vZ34iBZoxAoGBAJOHWoCFmxoPbHSfFKzJ
c8xv+KsaPf1XPfTgLkJUL7k/zxHa4wcu3qNfFSjof3TH81wt2wKYyxBjQ+PD+FF4
JjxnHOynwq3uZIvqpkROJ7g3vvrAWEWgbgiKdo7B4cVviNlBEl4WI4B0WWNn1XP1
1z8vrxRRLs6u/Kv25MNpAByr
-----END PRIVATE KEY-----`
    prviaKey := []byte(privateKeyPEM)

	// Example encrypted data (replace with your actual encrypted data)
	encryptedBase64 := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAssm0BP07iStvAi2xnnlMp7w6yNSpWuiyDyF43TQiOAE3qeEJPKe89uI2U9Ci8rmFFhO0lu9iTmbNvUTRnM0PbK7qTTUVtelESGvWGhfksicxmBuktWWjrgl4lSFVv7StMqaygFmcFVnzgHNY5Hs1peYmvAr3Ikipeu5WFH5PCkRGatHRv/pWOJwOeiDmm8dv/WiYBXzJijMIvDAbLOQZv/oa13Zk6oBNnXjgjwKP7sAJTBaOwCCbZHHrP7Dm3+XsbRPaogOnC2yfm9G5k3FC/OPpHMMFZX/FCyqcVxSPc9RfoSpHodfPtchFiq1oyfM1aP1JeQPXqQx9ALsFVi0ofwIDAQAB"

	privateKeyBlock, _ := pem.Decode(prviaKey)
	privateKey, err := x509.ParsePKCS8PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		panic(err)
	}

	// Decrypt the ciphertext
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		panic(err)
	}

	decrypted, err := rsaDecrypt(ciphertext, privateKey.(*rsa.PrivateKey))
	if err != nil {
		panic(err)
	}

	// Print the decrypted plaintext
	println(string(decrypted))
	c.String(200, string(decrypted))
}

func rsaDecrypt(ciphertext []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	// Parse the ciphertext
	block, _ := pem.Decode(ciphertext)
	if block == nil {
		return nil, errors.New("invalid ciphertext")
	}

	// Decrypt the ciphertext
	decrypted, err := rsa.DecryptOAEP(
		sha256.New(),
		rand.Reader,
		privateKey,
		block.Bytes,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}

func decryptOAEP(privateKey, cipherdata []byte) ([]byte, error) {
    block, _ := pem.Decode(privateKey)
    if block == nil {
        err := fmt.Errorf("failed to parse certificate PEM")
        return nil, err
    }

    priv, err := x509.ParsePKCS1PrivateKey(block.Bytes) // ASN.1 PKCS#1 DER encoded form.
    if err != nil {
        return nil, err
    }

    h := sha256.New()
    return rsa.DecryptOAEP(h, rand.Reader, priv, cipherdata, nil)
}
