// Collection of subscriber for the Remote SwitchBricklet
package remoteswitch

import (
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
)

// GetSwitchingState creates a subscriber to get the state of the remote switch
// bricklet.
func GetSwitchingState(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetSwitchingState"),
		Fid:        function_get_switching_state,
		Uid:        uid,
		Result:     &State{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetSwitchingStateFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func GetSwitchingStateFuture(brick *bricker.Bricker, connectorname string, uid uint32) bool {
	future := make(chan bool)
	sub := GetSwitchingState("getswitchingstate"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v bool = false
			if err == nil {
				if value, ok := r.(*State); ok {
					if value.State == uint8(1) {
						v = true
					}
				}
			}
			future <- v
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// State is the state of the remote switch bricklet. If State.State is 1, the
// remote switch bricklet is busy. If its 0, the bricklet is not busy.
type State struct {
	State uint8
}

// FromPacket creates from a packet a State.
func (s *State) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(s, p); err != nil {
		return err
	}
	return p.Payload.Decode(s)
}

// String fullfill the stringer interface.
func (s *State) String() string {
	if s.State == uint8(0) {
		return "off"
	}

	return "on"
}

// Copy creates a copy of the content.
func (s *State) Copy() device.Resulter {
	if s == nil {
		return nil
	}
	return &State{
		State: s.State,
	}
}
