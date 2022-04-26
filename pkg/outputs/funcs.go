package outputs

import (
	"github.com/gin-gonic/gin"
)

func resolveValue(ctx *gin.Context) string {
	return ctx.Param("value")
}

func getBody(ctx *gin.Context) (dto *Body, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func getBodyBulk(ctx *gin.Context) (dto []BulkWrite, err error) {
	err = ctx.ShouldBindJSON(&dto)
	return dto, err
}

func reposeHandler(body interface{}, err error, ctx *gin.Context) {
	if err != nil {
		if err == nil {
			ctx.JSON(404, Message{Message: "unknown error"})
		} else {
			if body != nil {
				ctx.JSON(404, Message{Message: err.Error()})
			} else {
				ctx.JSON(404, Message{Message: err.Error()})
			}
		}
	} else {
		ctx.JSON(200, body)
	}
}

// Percent - calculate what %[number1] of [number2] is.
// ex. 25% of 200 is 50
func Percent(percent int, all int) float64 {
	return (float64(all) * float64(percent)) / float64(100)
}
