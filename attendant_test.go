package parkinglot

import "testing"

func TestCreateNewAttendant(t *testing.T) {
	parkinglog, _ := NewParkingLot(1)
	_, err := NewAttendant(parkinglog)

	if err != nil {
		t.Error("new attendant should be created")
	}
}

func TestAttedantCannotBeCreatedWithNilParkingLot(t *testing.T) {
	_, err := NewAttendant(nil)

	if err == nil {
		t.Error("attendant should have be created with nil parking lot")
	}
}
