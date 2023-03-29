package main

import (
	"log"

	"github.com/google/gousb"
)

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
	// pci, err := ghw.PCI()
	// if err != nil {
	// 	return false, err
	// }

	// for _, device := range pci.Devices {
	// 	if device.Vendor.ID == pciVendorIDValve &&
	// 		device.Product.ID == pciDeviceIDValveIndex {
	// 		return true, nil
	// 	}
	// }

	// Only one context should be needed for an application.  It should always be closed.
	ctx := gousb.NewContext()
	defer ctx.Close()

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
	if deviceExists, err := deviceExists(usbVendorIDValve, usbProductIDValveIndex); err != nil {
		return err
	} else if !deviceExists {
		log.Println("Valve Index not found, exiting without doing anything")
		return nil
	}

	// start SteamVR
	log.Println("Valve Index found, starting:", steamVRURL)
	return ShellExecute(steamVRURL, "", "")
}
