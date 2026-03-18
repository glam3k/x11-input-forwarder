package main

import (
	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
	"github.com/jezek/xgb/xtest"
)

func MouseMove(xconn *xgb.Conn, root xproto.Window, x, y uint16) {
	_ = xtest.FakeInput(xconn, xproto.MotionNotify, 0, 0, root, int16(x), int16(y), 0)
}

func MouseButton(xconn *xgb.Conn, root xproto.Window, button byte, press bool) {
	eventType := byte(xproto.ButtonRelease)
	if press {
		eventType = byte(xproto.ButtonPress)
	}
	_ = xtest.FakeInput(xconn, eventType, button, 0, root, 0, 0, 0)
}

func Key(xconn *xgb.Conn, root xproto.Window, keycode byte, press bool) {
	eventType := byte(xproto.KeyRelease)
	if press {
		eventType = byte(xproto.KeyPress)
	}
	_ = xtest.FakeInput(xconn, eventType, keycode, 0, root, 0, 0, 0)
}

func Scroll(xconn *xgb.Conn, root xproto.Window, dx, dy int8) {
	if dy < 0 {
		repeatButton(xconn, root, 4, -int(dy))
	} else if dy > 0 {
		repeatButton(xconn, root, 5, int(dy))
	}

	if dx < 0 {
		repeatButton(xconn, root, 6, -int(dx))
	} else if dx > 0 {
		repeatButton(xconn, root, 7, int(dx))
	}
}

func repeatButton(xconn *xgb.Conn, root xproto.Window, button byte, n int) {
	for i := 0; i < n; i++ {
		MouseButton(xconn, root, button, true)
		MouseButton(xconn, root, button, false)
	}
}
