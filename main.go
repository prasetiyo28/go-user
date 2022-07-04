package main

import (
	"fmt"
	"net/http"
	"os"

	"go-user/infrastructure"
	"go-user/pkg"
	"go-user/presenter"
	repositoryAccount "go-user/repository/account"
	"go-user/router"
	serviceAccount "go-user/service/account"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	serverHost := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("PORT")
	serverString := fmt.Sprintf("%s:%s", serverHost, serverPort)
	r := gin.Default()

	dbUrl := os.Getenv("DB_URL")

	gormConfig := infrastructure.InitDB(dbUrl)

	repoAccount := repositoryAccount.NewRepository(gormConfig)
	serviceAccount := serviceAccount.NewService(repoAccount)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hellow",
		})
	})

	v1 := r.Group("/api/v1")
	{
		router.RouteAccount(v1, serviceAccount)
	}

	errorHandler(r)
	r.Run(serverString)

}

//handle error method and not found endpoint
func errorHandler(r *gin.Engine) {
	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, presenter.Response{
			Status:  pkg.HTTP_STATUS_ERROR,
			Message: pkg.ErrMethodNotAllow.Error(),
		})
		c.Abort()
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, presenter.Response{
			Code:    http.StatusNotFound,
			Message: pkg.ErrInvalidURL.Error(),
		})
		c.Abort()
	})
}
