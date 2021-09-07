package main

import (
	"cinemas/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func cinemasGET(context *gin.Context) {
	cinemas := model.GetAllCinemas()
	context.JSON(http.StatusOK, gin.H{"cinemas": cinemas})
}

func cinemaGET(context *gin.Context) {
	idCinema, _ := strconv.Atoi(context.Param("idCinema"))
	cinema := model.FindCinema(idCinema)
	context.JSON(http.StatusOK, gin.H{"cinema": cinema})
}

func cinemaPOST(context *gin.Context) {
	var cinema model.Cinema

	if err := context.ShouldBindBodyWith(&cinema, binding.JSON); err != nil {
		log.Printf("error: %+v", err)
	}

	model.NewCinema(cinema)
	context.JSON(http.StatusCreated, gin.H{"message": "cinema created!"})
}

func cinemaPUT(context *gin.Context) {
	var cinema model.Cinema

	if err := context.ShouldBindBodyWith(&cinema, binding.JSON); err != nil {
		log.Printf("error: %+v", err)
	}

	idCinema, _ := strconv.Atoi(context.Param("idCinema"))

	model.UpdateCinema(idCinema, cinema)
	context.JSON(http.StatusOK, gin.H{"message": "cinema" + context.Param("idCinema") + " updated!"})
}

func cinemaDELETE(context *gin.Context) {
	context.JSON(http.StatusGone, gin.H{"message": "cinema" + context.Param("idCinema") + " deleted!"})
}
