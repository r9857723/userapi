package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"userapi/services"
)

type ResponseWithJson struct {
	Code    int         `json="code"`
	Message string      `json="message"`
	Result  interface{} `json="result"`
}

func CreateUser(ctx *gin.Context) {
	account := ctx.PostForm("account")
	password := ctx.PostForm("password")

	ctx.JSON(http.StatusOK, gin.H{})
	r := &ResponseWithJson{
		Code:    0,
		Message: "",
		Result:  services.IsOk(false),
	}
	if err := services.CreateUser(account, password); err == nil {
		r.Result = services.IsOk(true)
	} else {
		r.Message = err.Error()
	}
	ctx.JSON(http.StatusOK, r)
}

func DeleteUser(ctx *gin.Context) {
	account := ctx.PostForm("account")
	r := &ResponseWithJson{
		Code:    0,
		Message: "",
		Result:  services.IsOk(false),
	}
	if rtn := services.DeleteUser(account); rtn {
		r.Result = services.IsOk(true)
	}

	ctx.JSON(http.StatusOK, r)
}

func ChangePwdByUser(ctx *gin.Context) {
	account := ctx.PostForm("account")
	password := ctx.PostForm("password")

	r := &ResponseWithJson{
		Code:    0,
		Message: "",
		Result:  services.IsOk(false),
	}
	if err := services.ChangePwdByUser(account,password); err == nil {
		r.Result = services.IsOk(true)
	} else {
		r.Message = err.Error()
	}
	ctx.JSON(http.StatusOK, r)
}

func Login(ctx *gin.Context) {
	account := ctx.Query("account")
	password := ctx.Query("password")
	r := &ResponseWithJson{
		Code:    0,
		Message: "",
		Result:  nil,
	}
	if rtn, err := services.Login(account, password); !rtn {
		r.Message = err.Error()
	}
	ctx.JSON(http.StatusOK, r)
}
