#!/usr/bin/env python3

import socket

HOST = "127.0.0.1"
PORT = 9300

with socket.create_connection((HOST, PORT)) as s:
    # TODO: send test protocol bytes
    pass
