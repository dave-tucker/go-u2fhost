package hid

import (
	"fmt"
	"time"

	"github.com/karalabe/hid"
)

type RawHidDevice struct {
	Device hid.DeviceInfo
	Handle *hid.Device
}

func newRawHidDevice(dev hid.DeviceInfo) *RawHidDevice {
	return &RawHidDevice{
		Device: dev,
	}
}

func (dev *RawHidDevice) Open() error {
	handle, err := dev.Device.Open()
	if err != nil {
		return err
	}
	dev.Handle = handle
	return nil
}

func (dev *RawHidDevice) Close() {
	if dev.Handle != nil {
		(*dev.Handle).Close()
		dev.Handle = nil
	}
}

func (dev *RawHidDevice) Write(data []byte) (int, error) {
	return dev.Handle.Write(data)
}

func (dev *RawHidDevice) ReadTimeout(response []byte, timeout int) (int, error) {
	bytesRead := 0
	var err error
	done := make(chan bool, 1)
	defer close(done)

	go func() {
		bytesRead, err = dev.Handle.Read(response)
		done <- true
	}()

	timer := time.NewTimer(time.Duration(timeout) * time.Second)
	defer timer.Stop()

	select {
	case <-done:
		return bytesRead, err
	case <-timer.C:
		return bytesRead, fmt.Errorf("timed out")
	}

}
