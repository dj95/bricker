package remoteswitch

import (
	"fmt"
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
)

// SetRepeats sets the repeats(how often one signal should be send) with a
// subscriber
func SetRepeats(id string, uid uint32, d *Repeats, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SetRepeats"),
		Fid:        function_set_repeats,
		Uid:        uid,
		Data:       d,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// SetRepeatsFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetRepeatsFuture(brick *bricker.Bricker, connectorname string, repeats uint8, uid uint32) bool {
	future := make(chan bool)
	sub := SetRepeats("setrepeats"+device.GenId(), uid,
		&Repeats{
			Repeats: repeats,
		},
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	b := <-future
	close(future)
	return b
}

// GetRepeats creates a subscriber to get the repeats for one signal that is send
func GetRepeats(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetRepeats"),
		Fid:        function_get_repeats,
		Uid:        uid,
		Result:     &Repeats{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetRepeatsFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is nil.
func GetRepeatsFuture(brick *bricker.Bricker, connectorname string, uid uint32) uint8 {
	future := make(chan uint8)
	defer close(future)
	sub := GetRepeats("getrepeats"+device.GenId(), uid,
		func(r device.Resulter, err error) {
			var v uint8 = 0
			if err == nil {
				if value, ok := r.(*Repeats); ok {
					v = value.Repeats
				}
			}
			future <- v
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return 0
	}
	return <-future
}

// Repeats represents the repeats for the remote switch bricklets. The signal
// is send as many times, as Repeats.Repeats holds.
type Repeats struct {
	Repeats uint8
}

// FromPacket creates from a packet a Repeats.
func (r *Repeats) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(r, p); err != nil {
		return err
	}
	return p.Payload.Decode(r)
}

// String fullfill the stringer interface.
func (r *Repeats) String() string {
	return fmt.Sprintf("Repeats: %d", r.Repeats)
}

// Copy creates a copy of the content.
func (r *Repeats) Copy() device.Resulter {
	if r == nil {
		return nil
	}
	return &Repeats{
		Repeats: r.Repeats,
	}
}
