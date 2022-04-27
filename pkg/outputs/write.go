package outputs

import (
	"github.com/NubeIO/nubeio-rubix-app-pi-gpio-go/pkg/common"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/numbers"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/types"
	"github.com/gin-gonic/gin"
)

func setWriteScale(in float64) (out float64) {
	out = numbers.Scale(in, 0, 100, 0, 1)
	return
}

func (inst *Outputs) Write(ctx *gin.Context) {
	body, err := getBody(ctx)
	if err != nil {
		common.ReposeHandler(nil, err, ctx)
		return
	}
	inst.IONum = body.IONum
	inst.valueOriginal = body.Value
	inst.Value = setWriteScale(body.Value)
	ok, err := inst.write()
	common.ReposeHandler(ok, err, ctx)
}

func (inst *Outputs) WriteOne(ctx *gin.Context) {
	val := resolveValue(ctx)
	writeValue := types.ToFloat64(val)
	inst.Value = setWriteScale(writeValue)
	inst.valueOriginal = writeValue
	inst.IONum = resolveIONum(ctx)
	ok, err := inst.write()
	common.ReposeHandler(ok, err, ctx)
}

func (inst *Outputs) WriteAll(ctx *gin.Context) {
	val := resolveValue(ctx)
	writeValue := types.ToFloat64(val)
	inst.Value = setWriteScale(writeValue)
	inst.valueOriginal = writeValue
	arr := []string{"UO1", "UO2", "UO3", "UO4", "UO5", "UO6", "DO1", "DO1"}
	for _, io := range arr {
		inst.IONum = io
		write, err := inst.write()
		if err != nil {
			common.ReposeHandler(write, err, ctx)
			return
		}
	}
	common.ReposeHandler(true, nil, ctx)
}
