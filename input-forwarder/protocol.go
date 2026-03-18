package main

import (
	"encoding/binary"
	"fmt"
	"io"
)

type EventType uint8

const (
	EventMouseMove EventType = 0x01
	EventMouseDown EventType = 0x02
	EventMouseUp   EventType = 0x03
	EventKeyDown   EventType = 0x04
	EventKeyUp     EventType = 0x05
	EventScroll    EventType = 0x06
)

type Event struct {
	Type   EventType
	X, Y   uint16
	Button uint8
	Keysym uint32
	DX, DY int8
}

func ReadEvent(r io.Reader) (Event, error) {
	var typeBuf [1]byte
	if _, err := io.ReadFull(r, typeBuf[:]); err != nil {
		return Event{}, err
	}

	ev := Event{Type: EventType(typeBuf[0])}

	switch ev.Type {
	case EventMouseMove:
		var buf [4]byte
		if _, err := io.ReadFull(r, buf[:]); err != nil {
			return Event{}, err
		}
		ev.X = binary.BigEndian.Uint16(buf[0:2])
		ev.Y = binary.BigEndian.Uint16(buf[2:4])
		return ev, nil

	case EventMouseDown, EventMouseUp:
		var buf [1]byte
		if _, err := io.ReadFull(r, buf[:]); err != nil {
			return Event{}, err
		}
		ev.Button = buf[0]
		return ev, nil

	case EventKeyDown, EventKeyUp:
		var buf [4]byte
		if _, err := io.ReadFull(r, buf[:]); err != nil {
			return Event{}, err
		}
		ev.Keysym = binary.BigEndian.Uint32(buf[:])
		return ev, nil

	case EventScroll:
		var buf [2]byte
		if _, err := io.ReadFull(r, buf[:]); err != nil {
			return Event{}, err
		}
		ev.DX = int8(buf[0])
		ev.DY = int8(buf[1])
		return ev, nil

	default:
		return Event{}, fmt.Errorf("unknown event type 0x%02x", typeBuf[0])
	}
}
