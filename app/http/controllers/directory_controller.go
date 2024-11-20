package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type DirectoryClear struct{}

func (*DirectoryClear) PhotoDirClear(c *gin.Context) {
	dirPath := "/home/www/car/static/upload/photo/order"
	duration := time.Hour * 24 * 30 * 6

	err := deleteOldDir(dirPath, duration)

	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	//清空临时文件
	dirPath = "/home/www/car/static/upload/photo/temp"
	duration = time.Hour * 24 * 30
	err = deleteOldDir(dirPath, duration)

	if err != nil {
		c.String(http.StatusOK, err.Error())
	}

	c.String(http.StatusOK, "OK")
}

func (*DirectoryClear) PhotoDirMonth(c *gin.Context) {
	dirPath := "/home/www/car/static/upload/photo/order"
	duration := time.Hour * 24 * 30 * 3

	err := deleteOldDir(dirPath, duration)

	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	//清空临时文件
	dirPath = "/home/www/car/static/upload/photo/temp"
	duration = time.Hour * 24 * 30
	err = deleteOldDir(dirPath, duration)

	if err != nil {
		c.String(http.StatusOK, err.Error())
	}

	c.String(http.StatusOK, "OK")
}

func (*DirectoryClear) AlbumDirClear(c *gin.Context) {
	dirPath := "/home/www/car/static/upload/album/order"
	duration := time.Hour * 24 * 30 * 6

	err := deleteOldDir(dirPath, duration)

	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	//清空临时文件
	dirPath = "/home/www/car/static/upload/album/temp"
	duration = time.Hour * 24 * 30
	err = deleteOldDir(dirPath, duration)

	if err != nil {
		c.String(http.StatusOK, err.Error())
	}

	c.String(http.StatusOK, "OK")
}

func (*DirectoryClear) CalendarDirClear(c *gin.Context) {
	dirPath := "/home/www/car/static/upload/calendar/order"
	duration := time.Hour * 24 * 30 * 6

	err := deleteOldDir(dirPath, duration)

	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	//清空临时文件
	dirPath = "/home/www/car/static/upload/calendar/temp"
	duration = time.Hour * 24 * 30
	err = deleteOldDir(dirPath, duration)

	if err != nil {
		c.String(http.StatusOK, err.Error())
	}

	c.String(http.StatusOK, "OK")
}

func (*DirectoryClear) TshirtDirClear(c *gin.Context) {
	dirPath := "/home/www/car/static/upload/tshirt/order"
	duration := time.Hour * 24 * 30 * 6

	err := deleteOldDir(dirPath, duration)

	if err != nil {
		c.String(http.StatusOK, err.Error())
	}
	//清空临时文件
	dirPath = "/home/www/car/static/upload/tshirt/temp"
	duration = time.Hour * 24 * 30
	err = deleteOldDir(dirPath, duration)

	if err != nil {
		c.String(http.StatusOK, err.Error())
	}

	c.String(http.StatusOK, "OK")
}

func deleteOldDir(dirPath string, duration time.Duration) error {
	now := time.Now()
	files, err := os.ReadDir(dirPath)

	if err != nil {
		return err
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


	return err

}
