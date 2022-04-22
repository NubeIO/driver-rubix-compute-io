package outputs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	gpio "github.com/stianeikeland/go-rpio/v4"
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

func (inst *Outputs) pinSelect() gpio.Pin {
	fmt.Println(88888)
	io := inst.IONum
	if io == OutputMaps.UO1.IONum {
		return UO3
	} else if io == OutputMaps.UO3.IONum {
		fmt.Println(88888)
		fmt.Println(UO3)
		return UO3
	} else if io == OutputMaps.UO5.IONum {
		return UO5
	}

	//if io == OutputMaps.UO1.IONum {
	//	return UO1
	//} else if io == OutputMaps.UO2.IONum {
	//	return UO2
	//} else if io == OutputMaps.UO3.IONum {
	//	return UO3
	//} else if io == OutputMaps.UO4.IONum {
	//	return UO4
	//} else if io == OutputMaps.UO5.IONum {
	//	return UO5
	//} else if io == OutputMaps.UO6.IONum {
	//	return UO6
	//} else if io == OutputMaps.DO1.IONum {
	//	return DO1
	//} else if io == OutputMaps.DO2.IONum {
	//	return DO2
	//}
	return UO3
}

func reposeHandler(body interface{}, err error, ctx *gin.Context) {
	if err != nil {
		if err == nil {
			ctx.JSON(404, Message{Message: "unknown error"})
		} else {
			if body != nil {
				ctx.JSON(404, body)
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
