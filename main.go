package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserCredential struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func login(c *gin.Context) {
	// username := c.PostForm("username")
	// c.PostForm("password")
	// c.String(http.StatusOK, "login berhasil, halo %s", username)

	var uc UserCredential
	if err := c.ShouldBind(&uc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": uc.Username,
		})
	}
}

func greeting(c *gin.Context) {
	name := c.Param("name")
	kec := c.Query("kecamatan")
	kel := c.Query("kelurahan")
	c.String(http.StatusOK, "Hello %s saat ini kamu berada di kec %s kel %s", name, kec, kel)
}

func main() {
	routerEngine := gin.Default()

	rgApiV1 := routerEngine.Group("/api/v1/")

	routerEngine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Healthy Check")
	})

	rgAuth := rgApiV1.Group("/auth")
	rgAuth.POST("/login", login)

	rgMaster := rgApiV1.Group("/master")
	rgMaster.GET("/greeting/:name", greeting)

	err := routerEngine.Run("localhost:8080")
	if err != nil {
		panic(err)
	} // secara default menggunakan port :8080
}
