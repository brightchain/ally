package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func PhotoDirClear(c *gin.Context) {
	dirPath := "/home/www/sharelive/src/static/upload/photo/order"
	duration := time.Hour * 24 * 30 * 6

	now := time.Now()
	files, err := os.ReadDir(dirPath)

	if err != nil {
		c.String(http.StatusOK, "获取目录列表失败")
		c.Abort()
	}

	for _, file := range files {
		folderPath := dirPath + "/" + file.Name()
		if file.IsDir() {
			info, err := os.Stat(folderPath)
			if err != nil {
				fmt.Printf("Failed to get info of directory: %s, error: %v\n", folderPath, err)
				continue
			}
			if now.Sub(info.ModTime()) > duration {
				os.RemoveAll(folderPath)
			}
		}
	}

	if err != nil {
		c.String(http.StatusOK, err.Error())
	}

	c.String(http.StatusOK, "OK")
}
