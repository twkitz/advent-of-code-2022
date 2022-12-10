package helper

import (
	"fmt"
	"testing"
)

func TestGetSignal(t *testing.T) {
	type test struct {
		cycle int
		x     int
		want  int
	}

	tests := []test{
		{1, 1, 1},
		{2, 3, 6},
	}

	for _, tc := range tests {
		device := Device{cycle: tc.cycle, x: tc.x}
		got := device.GetSignal()
		if got != tc.want {
			t.Error(fmt.Sprintf("cycle: %d, x: %d, want: %d, got: %d", device.cycle, device.x, tc.want, got))
		}
	}
}

func TestNoop(t *testing.T) {
	device := Device{cycle: 0}
	device.Noop()
	if device.cycle != 1 {
		t.Error(fmt.Sprintf("want: %d, got: %d", 1, device.cycle))
	}
}

func TestAddX(t *testing.T) {
	device := Device{cycle: 1, x: 2}
	device.AddX(2)
	if device.cycle != 3 || device.x != 4 {
		t.Error(fmt.Sprintf("wantedCycle: 3, cycle: %d, wantedX: 4, got: %d", device.cycle, device.x))
	}
}

func TestUpdateTotalSignal(t *testing.T) {
	type test struct {
		cycle                int
		interestedCycleIndex int
		totalSignal          int
		want                 int
	}

	tests := []test{
		{1, 0, 0, 0},
		{20, 0, 0, 20},
		{40, 1, 20, 60},
	}

	for _, tc := range tests {
		device := Device{cycle: tc.cycle, x: 1, interestedCycleIndex: tc.interestedCycleIndex, interestedCycles: []int{20, 40}, totalSignal: tc.totalSignal}
		device.UpdateTotalSignal()
		if device.totalSignal != tc.want {
			t.Error(fmt.Sprintf("want: %d, got: %d", tc.want, device.totalSignal))
		}
	}
}

func TestIsLitPixel(t *testing.T) {
	type test struct {
		cycle int
		x     int
		want  bool
	}

	tests := []test{
		{1, 1, true},
		{2, 1, true},
		{3, 1, true},
		{4, 1, false},
		{3, 16, false},
		{4, 16, false},
		{5, 5, true},
		{6, 5, true},
		{41, 1, true},
		{42, 1, true},
		{43, 1, true},
	}

	for _, tc := range tests {
		device := Device{cycle: tc.cycle, x: tc.x}
		got := device.IsLitPixel()
		if got != tc.want {
			t.Error(fmt.Sprintf("cycle: %d, x: %d, want: %t, got: %t", tc.cycle, tc.x, tc.want, got))
		}
	}
}
