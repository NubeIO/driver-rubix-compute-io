package outputs

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/types"
	"github.com/gin-gonic/gin"
)

type BulkWrite struct {
	IONum string  `json:"IONum"`
	Value float64 `json:"value"`
}

func (inst *Outputs) BulkWrite(ctx *gin.Context) {
	body, err := getBodyBulk(ctx)
	if err != nil {
		reposeHandler(false, err, ctx)
		return
	}
	for _, io := range body {
		writeValue := types.ToFloat64(io.Value)
		inst.Value = setWriteScale(writeValue)
		inst.valueOriginal = writeValue
		inst.IONum = io.IONum
		write, err := inst.write()
		if err != nil {
			reposeHandler(write, err, ctx)
			return
		}
	}
	reposeHandler(true, nil, ctx)
	return
}
