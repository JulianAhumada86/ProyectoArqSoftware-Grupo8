package controllers

import (
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func InsertImage(ctx *gin.Context) {
	im, err := ctx.FormFile("imagen")
	log.Print(ctx.FormFile("hotel"))
	if err != nil {
		print(err)
		ctx.Status(200)
		return
	}

	im2, err := im.Open()
	if err != nil {
		log.Print(err)
		ctx.Status(200)
		return
	}

	im3, err := ioutil.ReadAll(im2)

	if err != nil {
		log.Print(err)
		ctx.Status(200)
		return
	}

	ctx.Data(http.StatusOK, "image/jpeg", im3)
}
