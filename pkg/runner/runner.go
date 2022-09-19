package runner

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/nubeio-rubix-app-pi-gpio-go/pkg/inputs"
	"github.com/NubeIO/nubeio-rubix-app-pi-gpio-go/pkg/mqttclient"
	log "github.com/sirupsen/logrus"
	"time"
)

type InputRunner struct {
	inputs      *inputs.Inputs
	Enable      bool
	TestMode    bool
	LoopDelayMs time.Duration // in ms
}

func NewRunner(in *InputRunner) *InputRunner {
	return in
}

func (inst *InputRunner) Runner() {
	inst.runner()
}

func (inst *InputRunner) runner() {
	if inst.LoopDelayMs == 0 {
		inst.LoopDelayMs = 1000
	}
	mqttBroker := "tcp://0.0.0.0:1883"
	_, err := mqttclient.InternalMQTT(mqttBroker)
	if err != nil {
		log.Errorf(fmt.Sprintf("mqttbase-subscribe-connect err:%s", err.Error()))
	}
	if inst.Enable {
		client, connected := mqttclient.GetMQTT()
		if !connected {
			log.Errorf("inputs-runner failed to connect to broker")
		}
		delay := inst.LoopDelayMs
		for {
			data, err := inst.inputs.Read(inst.TestMode)
			if err != nil {
				log.Errorf("inputs-runner err:%s", err.Error())
			}
			payload, err := json.Marshal(data)
			if err != nil {
				log.Errorf("inputs-runner error on json marshal:%s", err.Error())
			}
			err = client.Publish("test", mqttclient.AtMostOnce, false, payload)
			if err != nil {
				log.Errorf("inputs-runner mqtt publish:%s", err.Error())
			}
			time.Sleep(delay * time.Millisecond)
		}
	}

}
