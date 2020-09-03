package controllers

import (
	"cblog/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

type Controller struct {
	C *gin.Context

	Offset int
	Limit  int

	Pagination
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Pagination struct {
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
	Total    int64 `json:"total"`
}

func (this *Controller) Success(data interface{}, message string) {
	if message == "" {
		message = "操作成功"
	}
	this.C.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	})
	return
}

func (this *Controller) Error(code int, message string, data interface{}) {
	if data == nil {
		data = ""
	}
	this.C.JSON(code, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
	return
}

func (this *Controller) Paginate(data interface{}, total int64) {
	this.Pagination.Total = total
	this.Pagination.Total = total
	this.Pagination.Total = total
	this.C.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "操作成功",
		Data: map[string]interface{}{
			"list":       data,
			"pagination": this.Pagination,
		},
	})
	return
}

func (this *Controller) SetPage() {
	page := com.StrTo(this.C.Query("page")).MustInt()
	pageSize := com.StrTo(this.C.Query("page_size")).MustInt()
	if page <= 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = setting.AppSetting.PageSize
	}
	this.Pagination.Page = page
	this.Pagination.PageSize = pageSize

	this.Limit = pageSize
	this.Offset = (page - 1) * pageSize
	return
}
