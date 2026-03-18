package main

import (
	"errors"
	"io"
	"log"
	"net"
	"os"

	"github.com/jezek/xgb"
	"github.com/jezek/xgb/xproto"
	"github.com/jezek/xgb/xtest"
)

func main() {
	display := getenvDefault("DISPLAY", ":1")
	port := getenvDefault("PORT", "9300")

	xconn, err := xgb.NewConnDisplay(display)
	if err != nil {
		log.Fatalf("connect X display %q: %v", display, err)
	}
	defer xconn.Close()

	if err := xtest.Init(xconn); err != nil {
		log.Fatalf("init XTEST: %v", err)
	}

	root := xproto.Setup(xconn).DefaultScreen(xconn).Root

	keysymMap, err := buildKeysymMap(xconn)
	if err != nil {
		log.Fatalf("build keysym map: %v", err)
	}

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("listen on %s: %v", port, err)
	}
	defer ln.Close()

	log.Printf("input-forwarder listening on :%s display=%s", port, display)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}

		log.Printf("client connected from %s", conn.RemoteAddr())
		handleConn(xconn, root, keysymMap, conn)
		_ = conn.Close()
		log.Printf("client disconnected")
	}
}

func handleConn(
	xconn *xgb.Conn,
	root xproto.Window,
	keysymMap map[uint32]xproto.Keycode,
	conn net.Conn,
) {
	for {
		ev, err := ReadEvent(conn)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return
			}
			log.Printf("connection closed: %v", err)
			return
		}

		switch ev.Type {
		case EventMouseMove:
			MouseMove(xconn, root, ev.X, ev.Y)

		case EventMouseDown:
			MouseButton(xconn, root, ev.Button, true)

		case EventMouseUp:
			MouseButton(xconn, root, ev.Button, false)

		case EventKeyDown:
			kc, ok := keysymMap[ev.Keysym]
			if ok {
				Key(xconn, root, byte(kc), true)
			}

		case EventKeyUp:
			kc, ok := keysymMap[ev.Keysym]
			if ok {
				Key(xconn, root, byte(kc), false)
			}

		case EventScroll:
			Scroll(xconn, root, ev.DX, ev.DY)
		}

	}
}

func getenvDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
