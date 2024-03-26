package api

import (
	"ally/model"
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
