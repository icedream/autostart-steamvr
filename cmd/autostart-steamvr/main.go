package main

import (
	"log"

	"github.com/google/gousb"
)

var ctx *gousb.Context

const (
	usbVendorIDValve       gousb.ID = 0x28de
	usbProductIDValveIndex gousb.ID = 0x2102
	steamVRURL                      = `steam://run/250820`
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
		return
	}
}

func deviceExists(vendorID, productID gousb.ID) (bool, error) {
	dev, err := ctx.OpenDeviceWithVIDPID(vendorID, productID)
	if err != nil {
		return false, err
	}
	if dev == nil {
		return false, nil
	}
	defer dev.Close()

	return true, nil
}

func run() error {
	ctx = gousb.NewContext()
	defer ctx.Close()

	if deviceExists, err := deviceExists(usbVendorIDValve, usbProductIDValveIndex); err != nil {
		return err
	} else if !deviceExists {
		log.Println("Valve Index not found, exiting without doing anything")
		return nil
	}

	// start SteamVR
	log.Println("Valve Index found, starting:", steamVRURL)
	return ShellExecute(steamVRURL)
}
