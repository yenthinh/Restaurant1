package controller

import (
	"github.com/gin-gonic/gin"
)

func InitController() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")
	router.Static("/assets", "./views/assets/")

	router.GET("/", Index)
	router.GET("/signin", Signin)
	router.GET("/login", Login)
	router.POST("/signin/action", SigninAction)
	router.POST("login/action", LoginAction)

	router.GET("/admin", ChefInfo)
	router.GET("/admin/menu", AdminMenu)
	router.GET("/admin/addmenu", AddMenu)
	router.GET("/admin/addchef", AddChef)
	router.POST("/admin/addmenu/action", AddMenuAction)
	router.POST("/admin/addchef/action", AddChefAction)
	router.POST("/admin/updatemenu/action/:id_product", UpdateMenu) //API
	router.POST("/admin/displaymenu/action/:id_product", DisplayMenu)
	
	router.GET("/admin/arrival",Arrival)
	router.POST("/admin/arrival/action",ArrivalAction)
	router.POST("/admin/arrivalcheck/action/:id_arrival",ArrivalCheck)
	router.POST("/admin/arrivalchecked/action/:id_arrival",ArrivalChecked)

	router.GET("/admin/contact",Contact)
	router.POST("admin/contact/action",ContactAction)
	router.Run()
}
