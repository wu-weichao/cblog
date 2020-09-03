package upload

import (
	"cblog/pkg/setting"
	"github.com/satori/go.uuid"
	"log"
	"os"
	"path"
	"strings"
)

func GetPath() string {
	basePath := setting.UploadSetting.Path

	_, err := os.Stat(basePath)
	if err != nil {
		err = os.MkdirAll(basePath, os.ModePerm)
		if err != nil {
			log.Printf("[fail] cannot make dir %s", basePath)
		}
	}

	return basePath
}

func GetCustomPath(custom string) string {
	customPath := path.Join(setting.UploadSetting.Path, custom)

	_, err := os.Stat(customPath)
	if err != nil {
		err = os.MkdirAll(customPath, os.ModePerm)
		if err != nil {
			log.Printf("[fail] cannot make dir %s", customPath)
		}
	}

	return customPath
}

func CreateFileName(name string) string {
	ext := path.Ext(name)
	return strings.Replace(uuid.NewV4().String(), "-", "", -1) + ext
}
