package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() //new gin router initialization
	router.GET("/cinemas", cinemasGET)
	router.GET("/cinemas/:idCinema", cinemaGET)
	router.POST("/cinemas", cinemaPOST)
	router.PUT("/cinemas/:idCinema", cinemaPUT)
	router.DELETE("/cinemas/:idCinema", cinemaDELETE)
	router.Run(":8000") //running application, Default port is 8080
}
