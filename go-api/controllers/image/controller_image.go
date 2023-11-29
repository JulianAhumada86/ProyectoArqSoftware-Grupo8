package controllers

import (
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func InsertImage(ctx *gin.Context) {
	log.Print(ctx)
	ctx.Status(200)
}
