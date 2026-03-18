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

	for idx, sym := range reply.Keysyms {
		if sym == 0 {
			continue
		}
		kc := xproto.Keycode(int(minKc) + idx/perKeycode)
		if _, exists := keysymToKeycode[uint32(sym)]; !exists {
			keysymToKeycode[uint32(sym)] = kc
		}
	}

	return keysymToKeycode, nil
}
