package controller

import (
	"net/http"
	"restaurant/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil) // nil==null
}

func LoginAction(ctx *gin.Context) {
	var dataFromFE models.Account
	var account models.Account
	err := ctx.ShouldBind(&dataFromFE)
	if err != nil {
		panic(err)
	}
	db := models.OpenDB()
	models.GetAccount(db, &account, dataFromFE.Email, dataFromFE.Password)
	if account.ID == 0 {
		ctx.Redirect(http.StatusSeeOther, "/login")
	}
	ctx.SetCookie("account_id", strconv.Itoa(int(account.ID)), int(time.Now().Unix()+time.Now().AddDate(0, 0, 1).Unix()), "/", "*", false, false)
	ctx.Redirect(http.StatusSeeOther, "/admin/menu")
}
