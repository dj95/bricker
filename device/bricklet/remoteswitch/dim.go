package remoteswitch

import (
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
)

// DimSocketB creates a subscriber to set the dim value to the receiver with the
// house code on socket a remotes.
func DimSocketB(id string, uid uint32, d *DimB, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "DimSocketB"),
		Fid:        function_dim_socket_b,
		Uid:        uid,
		Data:       d,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// DimSocketBFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func DimSocketBFuture(brick *bricker.Bricker, connectorname string, address, unit, dimValue uint8, uid uint32) bool {
	future := make(chan bool)
	sub := DimSocketB("dimsocket_b"+device.GenId(), uid,
		&DimB{
			Address:  address,
			Unit:     unit,
			DimValue: dimValue,
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

type DimB struct {
	Address  uint8
	Unit     uint8
	DimValue uint8
}
