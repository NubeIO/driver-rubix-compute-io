package ping

import (
	"github.com/NubeIO/nubeio-rubix-app-pi-gpio-go/pkg/common"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/nube/nube"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/systemd/systemctl"

	"github.com/gin-gonic/gin"
)

type Ping struct {
	Ok     bool `json:"ok"`
	PigPIO bool `json:"pigio_is_running"`
}

func (inst *Ping) Ping(ctx *gin.Context) {
	active, err := systemctl.IsActive(nube.Services.PigPIO.ServiceName, systemctl.Options{})
	if err != nil {
		common.ReposeHandler(inst, err, ctx)
		return
	}
	inst.Ok = true
	inst.PigPIO = active
	common.ReposeHandler(inst, nil, ctx)

}
