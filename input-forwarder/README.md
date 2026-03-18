# input-forwarder

A small TCP server that receives a fixed-width binary input protocol and injects it into an X11 desktop using XTEST.

## Environment

- `DISPLAY` default: `:1`
- `PORT` default: `9300`

## Build

```bash
go build -o bin/input-forwarder .

