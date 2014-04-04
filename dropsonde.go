package dropsonde

import (
	"github.com/cloudfoundry-incubator/dropsonde/emitter"
	"github.com/cloudfoundry-incubator/dropsonde/events"
	"github.com/cloudfoundry-incubator/dropsonde/heartbeat"
	"sync"
)

var DefaultEmitterRemoteAddr = "localhost:42420"
var HeartbeatEmitterRemoteAddr = "localhost:42421"

var heartbeatState struct {
	sync.Mutex
	stopChannel chan<- interface{}
}

func Initialize(origin *events.Origin) error {
	if emitter.DefaultEmitter == nil {
		udpEmitter, err := emitter.NewUdpEmitter(DefaultEmitterRemoteAddr)
		if err != nil {
			return err
		}

		emitter.DefaultEmitter, err = emitter.NewInstrumentedEmitter(udpEmitter)
		if err != nil {
			return err
		}
	}

	emitter.DefaultEmitter.SetOrigin(origin)

	heartbeatState.Lock()
	defer heartbeatState.Unlock()

	if heartbeatState.stopChannel != nil {
		return nil
	}

	if heartbeatEventSource, ok := emitter.DefaultEmitter.(heartbeat.HeartbeatEventSource); ok {
		var err error
		if heartbeat.HeartbeatEmitter == nil {
			heartbeat.HeartbeatEmitter, err = emitter.NewTcpEmitter(HeartbeatEmitterRemoteAddr)
			if err != nil {
				return err
			}
		}

		heartbeatState.stopChannel, err = heartbeat.BeginGeneration(heartbeatEventSource, origin)
		if err != nil {
			return err
		}
	}

	return nil
}

func Cleanup() {
	heartbeatState.Lock()
	defer heartbeatState.Unlock()

	if heartbeatState.stopChannel != nil {
		close(heartbeatState.stopChannel)
		heartbeatState.stopChannel = nil
	}
}
