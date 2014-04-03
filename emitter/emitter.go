package emitter

import (
	"errors"
	"github.com/cloudfoundry-incubator/dropsonde/events"
	"log"
)

type Emitter interface {
	Emit(events.Event) error
	SetOrigin(*events.Origin)
	Close()
}

var DefaultEmitter Emitter

func init() {
	udpEmitter, err := NewUdpEmitter()
	if err != nil {
		log.Printf("WARNING: failed to create udpEmitter: %v\n", err)
	}

	DefaultEmitter, err = NewInstrumentedEmitter(udpEmitter)
	if err != nil {
		log.Printf("WARNING: failed to create instrumentedEmitter: %v\n", err)
	}
}

func Emit(e events.Event) error {
	if DefaultEmitter != nil {
		return DefaultEmitter.Emit(e)
	}

	return errors.New("Default emitter not set")
}
