package controller

import (
	"fmt"
	"net/http"
	"restaurant/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Arrival(ctx *gin.Context) {
	var arr []models.Arrival
	db := models.OpenDB()
	check:= true
	models.GetArrivals(db, &arr, check)
	ctx.HTML(http.StatusOK, "arrival.html", arr)
}
func ArrivalAction(ctx *gin.Context) {
	fmt.Print("DA VAO ARRIVAL")
	var aa models.Arrival
	err := ctx.ShouldBind(&aa)
	if err != nil {
		panic(err)
	}
	fmt.Print(aa.Uname)
	fmt.Print(aa.Umail)
	fmt.Print(aa.Uphone)
	fmt.Print(aa.Date)
	fmt.Print(aa.Time)
	fmt.Print(aa.Ofpeople)
	fmt.Print(aa.Message)
	db := models.OpenDB()
	err = models.CreateArrival(db, aa)
	if err != nil {
		panic(err)
	}
	ctx.Redirect(http.StatusSeeOther, "/")
}
func ArrivalCheck(ctx *gin.Context){
	fmt.Print("Checked")
	idStr := ctx.Param("id_arrival")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		panic(err)
	}
	db:= models.OpenDB()
	models.UpdaTeCheck(db, id, true)
	ctx.Redirect(http.StatusSeeOther, "/admin/arrival")
}
func ArrivalChecked(ctx *gin.Context){
	fmt.Print("Checked")
	idStr := ctx.Param("id_arrival")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		panic(err)
	}
	db:= models.OpenDB()
	models.UpdaTeCheck(db, id, false)
	ctx.Redirect(http.StatusSeeOther, "/admin/arrival")
}