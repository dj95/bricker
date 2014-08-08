// Copyright 2014 Dirk Jablonowski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Type for a device identity result. Every device should have a subscriber getidentity.
package identity

import (
	"fmt"
	"github.com/dirkjabl/bricker"
	"github.com/dirkjabl/bricker/device"
	"github.com/dirkjabl/bricker/device/name"
	"github.com/dirkjabl/bricker/net/base58"
	"github.com/dirkjabl/bricker/net/packet"
)

const (
	function_get_identity = uint8(255)
)

// GetIdentity creates the subscriber to get the identity of a device.
func GetIdentity(id string, uid uint32, handler func(device.Resulter, error)) *device.Device {
	return device.Generator{
		Id:         device.FallbackId(id, "GetIdentity"),
		Fid:        function_get_identity,
		Uid:        uid,
		Result:     &Identity{},
		Handler:    handler,
		WithPacket: true}.CreateDevice()
}

// Future is a future pattern version for a synchronized call of the subscriber.
// If an error occur, the result is nil.
func GetIdentityFuture(brick bricker.Bricker, connectorname string, uid uint32) *Identity {
	future := make(chan *Identity)
	defer close(future)
	sub := GetIdentity("getidentityfuture", uid,
		func(r device.Resulter, err error) {
			var v *Identity = nil
			if err == nil {
				if value, ok := r.(*Identity); ok {
					v = value
				}
			}
			future <- v
		})
	err := brick.Subscribe(sub, connectorname)
	if err != nil {
		return nil
	}
	v := <-future
	return v
}

// Result type for a getidentity subscriber.
type Identity struct {
	Uid             [8]byte
	ConnectedUid    [8]byte
	Position        byte
	HardwareVersion [3]uint8
	FirmwareVersion [3]uint8
	DeviceIdentifer uint16
}

// FromPacket fill up the values of the identity object from a net packet.
// Fullfill the resulter interface.
func (i *Identity) FromPacket(p *packet.Packet) error {
	if err := device.CheckForFromPacket(i, p); err != nil {
		return err
	}
	return p.Payload.Decode(i)
}

// Stringer interface fulfill.
func (e *Identity) String() string {
	uid := base58.Convert32(base58.Decode(e.Uid))
	cuid := base58.Convert32(base58.Decode(e.ConnectedUid))
	txt := "Identity ["
	txt += fmt.Sprintf("UID: %s (%d), ", e.Uid, uid)
	txt += fmt.Sprintf("Connected UID: %s (%d), ", e.ConnectedUid, cuid)
	txt += fmt.Sprintf("Position: %c, ", e.Position)
	txt += fmt.Sprintf("Hardware Version: %d.%d.%d, ", e.HardwareVersion[0],
		e.HardwareVersion[1], e.HardwareVersion[2])
	txt += fmt.Sprintf("Firmware Version: %d.%d.%d, ", e.FirmwareVersion[0],
		e.FirmwareVersion[1], e.FirmwareVersion[2])
	txt += "Name: " + name.Name(e.DeviceIdentifer) + "]"
	return txt
}
