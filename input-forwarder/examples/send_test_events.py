#!/usr/bin/env python3

import socket
import struct
import time

HOST = "127.0.0.1"
PORT = 9300

MOUSEMOVE = 0x01
MOUSEDOWN = 0x02
MOUSEUP = 0x03
KEYDOWN = 0x04
KEYUP = 0x05
SCROLL = 0x06

KEY_A = 0x0061
KEY_ENTER = 0xFF0D

def mousemove(x: int, y: int) -> bytes:
    return struct.pack(">BHH", MOUSEMOVE, x, y)

def mousedown(button: int) -> bytes:
    return struct.pack(">BB", MOUSEDOWN, button)

def mouseup(button: int) -> bytes:
    return struct.pack(">BB", MOUSEUP, button)

def keydown(keysym: int) -> bytes:
    return struct.pack(">BI", KEYDOWN, keysym)

def keyup(keysym: int) -> bytes:
    return struct.pack(">BI", KEYUP, keysym)

def scroll(dx: int, dy: int) -> bytes:
    return struct.pack(">Bbb", SCROLL, dx, dy)

with socket.create_connection((HOST, PORT)) as s:
    # Move near center
    s.sendall(mousemove(960, 540))
    time.sleep(0.2)

    # Left click
    s.sendall(mousedown(1))
    s.sendall(mouseup(1))
    time.sleep(0.2)

    # Type "a"
    s.sendall(keydown(KEY_A))
    s.sendall(keyup(KEY_A))
    time.sleep(0.2)

    # Press Enter
    s.sendall(keydown(KEY_ENTER))
    s.sendall(keyup(KEY_ENTER))
    time.sleep(0.2)

    # Scroll down one tick, then up one tick
    s.sendall(scroll(0, 1))
    time.sleep(0.2)
    s.sendall(scroll(0, -1))

print("sent test events")
