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
	topic := "rubixio/inputs/all"
	if inst.LoopDelayMs == 0 {
		inst.LoopDelayMs = 1000
	}

	if inst.Enable {
		delay := inst.LoopDelayMs
		var connected bool
		client := &mqttclient.Client{}
		for {

			if !connected {
				log.Info("inputs-runner try and connect to mqtt broker")
				mqttBroker := "tcp://0.0.0.0:1883"
				_, err := mqttclient.InternalMQTT(mqttBroker)
				if err != nil {
					log.Errorf(fmt.Sprintf("mqttbase-subscribe-connect err:%s", err.Error()))
				}
				client, connected = mqttclient.GetMQTT()
			}
			if !connected {
				log.Errorf("inputs-runner failed to connect to broker")
				time.Sleep(1000 * time.Millisecond) // try and reconnect
			}

			data, err := inst.inputs.Read(inst.TestMode)
			if err != nil {
				log.Errorf("inputs-runner err:%s", err.Error())
			}
			payload, err := json.Marshal(data)
			if err != nil {
				log.Errorf("inputs-runner error on json marshal:%s", err.Error())
			}
			err = client.Publish(topic, mqttclient.AtMostOnce, false, payload)
			if err != nil {
				log.Errorf("inputs-runner mqtt publish:%s", err.Error())
				connected = false // assume broker is not connected
			}
			time.Sleep(delay * time.Millisecond)

		}
	}

}
