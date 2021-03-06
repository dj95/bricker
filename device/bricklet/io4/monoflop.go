// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io4

import (
	"fmt"
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/net/packet"
)

// SetMonoflop creates the subscriber to set the monoflop timer value for specifed output pins (per bitmap).
func SetMonoflop(id string, uid uint32, m *Monoflops, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "SetMonoflop"),
		Fid:        function_set_monoflop,
		Uid:        uid,
		Data:       m,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// SetMonoflopFuture is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is false.
func SetMonoflopFuture(brick *bricker.Bricker, connectorname string, uid uint32, m *Monoflops) bool {
	future := make(chan bool)
	defer close(future)
	sub := SetMonoflop("setmonoflopfuture"+device.GenId(), uid, m,
		func(r device.Resulter, err error) {
			future <- device.IsEmptyResultOk(r, err)
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return false
	}
	return <-future
}

// GetMonoflop creates a subscriber for getting the actual monoflop value.
func GetMonoflop(id string, uid uint32, pin *Pin, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetMonoflop"),
		Fid:        function_get_monoflop,
		Uid:        uid,
		Result:     &Monoflop{},
		Data:       pin,
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// GetMonoflopFuture is a future pattern version for a synchronized all of the subscriber.
// If an error occur, the result is nil.
func GetMonoflopFuture(brick *bricker.Bricker, connectorname string, uid uint32, pin *Pin) *Monoflop {
	future := make(chan *Monoflop)
	defer close(future)
	sub := GetMonoflop("getmonoflopfuture"+device.GenId(), uid, pin,
		func(r device.Resulter, err error) {
			var v *Monoflop = nil
			if err == nil {
				if value, ok := r.(*Monoflop); ok {
					v = value
				}
			}
			future <- v
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return nil
	}
	return <-future
}

/*
MonoflopDone creates a subscriber for the monoflop done callback.
This callback is triggered whenever a monoflop timer reaches 0.
The response values contain the involved pins and the current value of the pins
(the value after the monoflop).
*/
func MonoflopDone(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "MonoflopDone"),
		Fid:        callback_monoflop_done,
		Uid:        uid,
		Result:     &Values{},
		Handler:    handler,
		IsCallback: true,
		WithPacket: false}.CreateDevice()
}

// Monoflops is a type to set bitmask(4bit) based the time to hold the value.
// The monoflop mechanismus works only with output pins.
// Non output pins will be ignored.
// The time is given in ms.
type Monoflops struct {
	SelectionMask uint8  // Bitmask (4bit)
	ValueMask     uint8  // Bitmask (4bit)
	Time          uint32 // ms
}

// Monoflop is the monflop timer value of a specified pin.
type Monoflop struct {
	Value         uint8
	Time          uint32 // in ms
	TimeRemaining uint32 // in ms
}

// FromPacket creates a Monoflop from a packet.
func (m *Monoflop) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(m, p); err != nil {
		return err
	}
	return p.Payload.Decode(m)
}

// String fullfill the stringer interface.
func (m *Monoflop) String() string {
	txt := "Monoflop "
	if m == nil {
		txt += "[nil]"
	} else {
		txt += fmt.Sprintf("[Value: %d, Time: %d ms, Time Remaining: %d ms]",
			m.Value, m.Time, m.TimeRemaining)
	}
	return txt
}

// Copy creates a copy of the content.
func (m *Monoflop) Copy() device.Resulter {
	if m == nil {
		return nil
	}
	return &Monoflop{
		Value:         m.Value,
		Time:          m.Time,
		TimeRemaining: m.TimeRemaining}
}
