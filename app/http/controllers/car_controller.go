package controllers

import (
	"encoding/json"
	"fmt"
	"h5/app/http/models"
	"h5/pkg/config"
	"h5/pkg/model"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Car struct {
}

// 定义结构体来匹配 JSON 数据
type Response struct {
	Status int           `json:"status"`
	Msg    string        `json:"msg"`
	Result []interface{} `json:"result"`
}

type CarBrand struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Initial   string `json:"initial"`
	ParentID  int    `json:"parentid"`
	Logo      string `json:"logo"`
	Depth     int    `json:"depth"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// 指定表名
func (CarBrand) TableName() string {
	return "car_brand"
}
func js(url string) ([]byte, error) {
	// 读取响应体
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()
	// 解析 JSON 数据

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%v", err)
		return nil, err
	}

	return body, nil
}

func (*Car) Index(c *gin.Context) {
	url := fmt.Sprintf("%s/car/brand?appkey=%s", config.GetString("jisu.url"), config.GetString("jisu.key"))
	result, err := js(url)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	// 解析 JSON 数据
	var response Response
	json.Unmarshal(result, &response)

	fmt.Printf("%v", response.Result)
	db := model.RDB["db2"]

	res := db.Db.Create(response.Result)
	if res.Error != nil {
		log.Printf("Error inserting car brand : %v", res.Error)
	} else {
		fmt.Printf("Inserted car brand: %v\n", res.RowsAffected)
	}
}

func (*Car) CarModel(c *gin.Context) {
	db := model.RDB["db2"]
	var brands []models.CarBrand
	dateTime := "2024-10-17 14:00:00"
	err := db.Db.Where("updated_at < ?", dateTime).Limit(50).Find(&brands)
	if err != nil {
		log.Printf("Error inserting car type : %v", err)
	}
	fmt.Printf("%v", len(brands))
	var ids []uint
	url := fmt.Sprintf("%s/car/type?appkey=%s", config.GetString("jisu.url"), config.GetString("jisu.key"))

	type Res struct {
		ID       string            `json:"id" `
		Name     string            `json:"name"`
		Fullname string            `json:"fullname"`
		Initial  string            `json:"initial"`
		List     []models.CarModel `json:"list"`
	}

	type Response struct {
		Status int    `json:"status"`
		Msg    string `json:"msg"`
		Result []Res  `json:"result"`
	}

	for _, val := range brands {
		url = fmt.Sprintf("%s&parentid=%d", url, val.Id)
		res, err := js(url)
		if err != nil {
			fmt.Printf("%v", err)
			break
		}

		var result Response
		var list []models.CarModel
		json.Unmarshal(res, &result)
		for v := range result.Result {
			for v1 := range result.Result[v].List {
				result.Result[v].List[v1].Brand = val.Name
				list = append(list, result.Result[v].List[v1])
			}
		}
		fmt.Printf("%v", list)
		db.Db.Create(&list)
		ids = append(ids, uint(val.Id))
	}
	if ids != nil {
		now := time.Now().Format("2006-01-02 15:04:05")
		res := db.Db.Model(&models.CarBrand{}).Where("id in ?", ids).Update("updated_at", now)
		if res.Error != nil {
			log.Printf("Error inserting car type : %v", res.Error)
		} else {
			fmt.Printf("Inserted car type: %v\n", res.RowsAffected)
		}

	}
}

func (*Car) CarDetail(c *gin.Context) {
	db := model.RDB["db2"]
	var carModels []models.CarModel
	dateTime := "2024-10-17 14:00:00"
	err := db.Db.Where("updated_at < ?", dateTime).Limit(500).Find(&carModels)
	if err != nil {
		log.Printf("Error inserting car type : %v", err)
	}
	fmt.Printf("%v", len(carModels))
	var ids []uint
	url := fmt.Sprintf("%s/car/car?appkey=%s", config.GetString("jisu.url"), config.GetString("jisu.key"))
	type JSONData struct {
		Status int    `json:"status"`
		Msg    string `json:"msg"`
		Result struct {
			ID        int32  `json:"id"`
			Name      string `json:"name"`
			Initial   string `json:"initial"`
			Fullname  string `json:"fullname"`
			Logo      string `json:"logo"`
			Salestate string `json:"salestate"`
			Depth     string `json:"depth"`
			List      []struct {
				ID              int32  `json:"id"`
				Name            string `json:"name"`
				Logo            string `json:"logo"`
				Price           string `json:"price"`
				Yeartype        string `json:"yeartype"`
				Productionstate string `json:"productionstate"`
				Salestate       string `json:"salestate"`
				Sizetype        string `json:"sizetype"`
			} `json:"list"`
		} `json:"result"`
	}
	for _, val := range carModels {
		url = fmt.Sprintf("%s&parentid=%d", url, val.Id)
		res, err := js(url)
		if err != nil {
			fmt.Printf("%v", err)
			break
		}
		var result JSONData
		json.Unmarshal(res, &result)
		var list []models.CarDetail
		fmt.Printf("%v", result)
		for v := range result.Result.List {
			data := models.CarDetail{
				Id:              result.Result.List[v].ID,
				Pid:             val.Id,
				Brand:           val.Brand,
				Name:            result.Result.List[v].Name,
				Fullname:        result.Result.Fullname,
				Logo:            result.Result.List[v].Logo,
				Price:           result.Result.List[v].Price,
				Yeartype:        result.Result.List[v].Yeartype,
				Productionstate: result.Result.List[v].Productionstate,
				Salestate:       result.Result.List[v].Salestate,
				Sizetype:        result.Result.List[v].Sizetype,
			}
			list = append(list, data)
		}
		fmt.Printf("%v", list)
		db.Db.Create(&list)
		ids = append(ids, uint(val.Id))
	}
	if ids != nil {
		now := time.Now().Format("2006-01-02 15:04:05")
		res := db.Db.Model(&models.CarModel{}).Where("id in ?", ids).Update("updated_at", now)
		if res.Error != nil {
			log.Printf("Error inserting car type : %v", res.Error)
		} else {
			fmt.Printf("Inserted car type: %v\n", res.RowsAffected)
		}

	}

}
