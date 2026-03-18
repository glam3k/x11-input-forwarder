package main

import "io"

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
	// TODO: implement protocol decoder
	return Event{}, nil
}
