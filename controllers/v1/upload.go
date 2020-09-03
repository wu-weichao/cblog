package v1

import (
	"cblog/controllers"
	"cblog/pkg/upload"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"time"
)

func UploadFile(c *gin.Context) {
	controller := controllers.Controller{C: c}

	form, _ := c.MultipartForm()
	basePath := upload.GetCustomPath(time.Now().Format("20060102"))
	result := make(map[string]interface{})
	var filePath string
	fielPaths := make([]string, 10)

	var fileCount int = 0
	for key, files := range form.File {
		fielPaths = fielPaths[:0]
		for _, file := range files {
			filePath = path.Join(basePath, upload.CreateFileName(file.Filename))
			err := c.SaveUploadedFile(file, filePath)
			if err != nil {
				controller.Error(http.StatusBadRequest, err.Error(), nil)
			}
			fielPaths = append(fielPaths, filePath)
			fileCount++
		}
		if len(fielPaths) > 1 {
			result[key] = fielPaths
		} else {
			result[key] = fielPaths[0]
		}
	}
	if fileCount == 1 {
		controller.Success(filePath, "")
		return
	}

	controller.Success(result, "")
}
