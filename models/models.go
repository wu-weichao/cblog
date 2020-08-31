package models

import (
	"fmt"
	"log"

	"cblog/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Model struct {
	ID        int            `gorm:"primary_key" json:"id"`
	CreatedAt int            `json:"created_at"`
	UpdatedAt int            `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"default:null;index" json:"deleted_at"`
}

var db *gorm.DB

func init() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name,
		setting.DatabaseSetting.Charset)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("models.Init err: %v", err)
	}
	// table migrate
	db.AutoMigrate(&User{})

	// modify Callback
	//db.Callback().Create().Replace("gorm:create", updateStampForCreateCallback)
	//db.Callback().Create().Replace("gorm:update", updateStampForUpdateCallback)
	//	//db.Callback().Create().Replace("gorm:delete", deleteCallback)

	//user := User{}
	//user.Name = "wwc"
	//db.Create(&user)
	//db.First(&user)
	//user.Name = "wwc改"
	//db.Save(&user)
	//db.Delete(&user)
	//fmt.Printf("Update: %v", user)
	CreateAdmin()
}
