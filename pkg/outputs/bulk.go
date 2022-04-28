package outputs

import (
	"github.com/NubeIO/nubeio-rubix-app-pi-gpio-go/pkg/common"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/types"
	"github.com/gin-gonic/gin"
)

type BulkWrite struct {
	IONum string `json:"io_num"`
	Value int    `json:"value"`
}

type BulkResponse struct {
	Ok bool `json:"ok"`
}

func (inst *Outputs) BulkWrite(ctx *gin.Context) {
	body, err := getBodyBulk(ctx)
	if err != nil {
		common.ReposeHandler(nil, err, ctx)
		return
	}
	for _, io := range body {
		writeValue := types.ToFloat64(io.Value)
		inst.Value = setWriteScale(writeValue)
		inst.valueOriginal = writeValue
		inst.IONum = io.IONum
		_, err := inst.write()
		if err != nil {
			common.ReposeHandler(nil, err, ctx)
			return
		}
	}
	res := &BulkResponse{Ok: true}
	common.ReposeHandler(res, nil, ctx)
	return
}
