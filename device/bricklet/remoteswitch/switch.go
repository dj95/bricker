// Collection of subscriber for the Remote SwitchBricklet
package remoteswitch

import (
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
)

// SwitchSocketA creates a subscriber to set the state to the receiver with the
// house code on socket a remotes.
func SwitchSocketA(id string, uid uint32, d *SwitchA, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SwitchSocketA"),
		Fid:        function_switch_socket_a,
		Uid:        uid,
		Data:       d,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// SwitchSocketAFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SwitchSocketAFuture(brick *bricker.Bricker, connectorname string, houseCode, receiverCode, switchTo uint8, uid uint32) bool {
	future := make(chan bool)
	sub := SwitchSocketA("switchsocket_a"+device.GenId(), uid,
		&SwitchA{
			HouseCode:    houseCode,
			ReceiverCode: receiverCode,
			SwitchTo:     switchTo,
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

type SwitchA struct {
	HouseCode    uint8
	ReceiverCode uint8
	SwitchTo     uint8
}

// SwitchSocketB creates a subscriber to set the state to the unit with the
// address on socket b remotes.
func SwitchSocketB(id string, uid uint32, d *SwitchB, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SwitchSocketB"),
		Fid:        function_switch_socket_b,
		Uid:        uid,
		Data:       d,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// SwitchSocketBFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SwitchSocketBFuture(brick *bricker.Bricker, connectorname string, address, unit, switchTo uint8, uid uint32) bool {
	future := make(chan bool)
	sub := SwitchSocketB("switchsocket_b"+device.GenId(), uid,
		&SwitchB{
			Address:  address,
			Unit:     unit,
			SwitchTo: switchTo,
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

type SwitchB struct {
	Address  uint8
	Unit     uint8
	SwitchTo uint8
}

// SwitchSocketC creates a subscriber to set the state to the deivce with the
// system code on socket c remotes.
func SwitchSocketC(id string, uid uint32, d *SwitchC, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SwitchSocketC"),
		Fid:        function_switch_socket_c,
		Uid:        uid,
		Data:       d,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// SwitchSocketCFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SwitchSocketCFuture(brick *bricker.Bricker, connectorname string, systemCode, deviceCode, switchTo uint8, uid uint32) bool {
	future := make(chan bool)
	sub := SwitchSocketC("switchsocket_c"+device.GenId(), uid,
		&SwitchC{
			SystemCode: systemCode,
			DeviceCode: deviceCode,
			SwitchTo:   switchTo,
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

type SwitchC struct {
	SystemCode uint8
	DeviceCode uint8
	SwitchTo   uint8
}
