package router

import (
	accountController "go-user/controller/account"
	"go-user/service/account"

	"github.com/gin-gonic/gin"
)

func RouteAccount(v1 *gin.RouterGroup, as *account.Service) {
	handler := accountController.AccountControllerHandler(as)
	account := v1.Group("")
	{
		account.POST("/DisplayUser", handler.DisplayUser)
		account.GET("/DisplayAllUsers", handler.DisplayAllUsers)
	}
}
