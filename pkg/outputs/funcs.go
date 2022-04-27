package outputs

import (
	"github.com/gin-gonic/gin"
)

func resolveValue(ctx *gin.Context) string {
	return ctx.Param("value")
}

func resolveIONum(ctx *gin.Context) string {
	return ctx.Param("io")
}

func getBody(ctx *gin.Context) (dto *Body, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func getBodyBulk(ctx *gin.Context) (dto []BulkWrite, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}
