package main

import (
	"fmt"

	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
)

func buildKeysymMap(xconn *xgb.Conn) (map[uint32]xproto.Keycode, error) {
	setup := xproto.Setup(xconn)
	minKc := setup.MinKeycode
	maxKc := setup.MaxKeycode

	count := byte(maxKc - minKc + 1)
	reply, err := xproto.GetKeyboardMapping(xconn, minKc, count).Reply()
	if err != nil {
		return nil, fmt.Errorf("get keyboard mapping: %w", err)
	}

	keysymToKeycode := make(map[uint32]xproto.Keycode)
	perKeycode := int(reply.KeysymsPerKeycode)

	for kc := minKc; kc <= maxKc; kc++ {
		base := int(kc-minKc) * perKeycode
		for i := 0; i < perKeycode; i++ {
			sym := reply.Keysyms[base+i]
			if sym == 0 {
				continue
			}
			if _, exists := keysymToKeycode[uint32(sym)]; !exists {
				keysymToKeycode[uint32(sym)] = kc
			}
		}
	}

	return keysymToKeycode, nil
}
