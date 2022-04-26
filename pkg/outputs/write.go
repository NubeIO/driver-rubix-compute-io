package outputs

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nils"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/numbers"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/types"
	"github.com/gin-gonic/gin"
	"time"
)

func setWriteScale(in float64) (out float64) {
	out = numbers.Scale(in, 0, 100, 0, 1)
	return
}

func (inst *Outputs) Write(ctx *gin.Context) {
	body, err := getBody(ctx)
	if err != nil {
		reposeHandler(nil, err, ctx)
		return
	}
	inst.IONum = body.IONum
	inst.valueOriginal = body.Value
	inst.Value = setWriteScale(body.Value)
	if nils.BoolIsNil(body.Debug) {
		inst.TestMode = true
	}
	time.Sleep(50 * time.Millisecond)
	ok, err := inst.write()
	reposeHandler(ok, err, ctx)
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
			reposeHandler(write, err, ctx)
			return
		}
	}
	reposeHandler(true, nil, ctx)
}
