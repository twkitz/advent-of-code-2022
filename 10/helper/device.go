package helper

type Device struct {
	cycle                int
	x                    int
	interestedCycles     []int
	interestedCycleIndex int
	totalSignal          int
	image                string
}

func (device Device) GetSignal() int {
	return device.cycle * device.x
}

func (device *Device) Noop() {
	device.cycle++
	device.UpdateTotalSignal()
}

func (device *Device) AddX(value int) {
	for i := 0; i < 2; i++ {
		device.Noop()
	}
	device.x += value
}

func (device *Device) UpdateTotalSignal() {
	if device.interestedCycleIndex < len(device.interestedCycles) &&
		device.cycle == device.interestedCycles[device.interestedCycleIndex] {
		device.totalSignal += device.GetSignal()
		device.interestedCycleIndex++
	}
	if device.IsLitPixel() {
		device.image += "#"
	} else {
		device.image += "."
	}
	if device.cycle%40 == 0 {
		device.image += "\n"
	}
}

func (device Device) IsLitPixel() bool {
	pixelPosition := device.cycle % 40
	if pixelPosition == 0 {
		pixelPosition = 40
	}
	return pixelPosition >= device.x && pixelPosition <= device.x+2
}
