# Autostart SteamVR

Just a little tool to detect on start whether the SteamVR headset (PCI ID
`28de:2102`) is connected. If detected, `steam://run/250820` will be called
(which will tell Steam to install and run SteamVR).

## Building

### Requirements

- libusb

### Steps

```sh
go get github.com/icedream/autostart-steamvr/cmd/autostart-steamvr
```

or check this repository out and run `go build ./cmd/autostart-steamvr`.

For Windows build on Linux make sure you have mingw-w64 and libusb for mingw-w64
installed, as well as `GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc
CGO_ENABLED=1` set and exported in environment. Copy `libusb-1.0.dll` and
`libssp-0.dll` along with the binary.
