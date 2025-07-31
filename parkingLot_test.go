package parkinglot

import (
	"testing"
)

func TestCreateLot(t *testing.T) {
	firstLot := &Lot{
		Id:          1,
		NumberPlate: "KJ-09-AK-123",
		Occupied:    false,
	}

	if firstLot.Id == 0 {
		t.Error("Struct Lot cannot be created with Id 0")
	}
}

func TestCreateParkingLot(t *testing.T) {
	_, err := NewParkingLot(10)
	if err != nil {
		t.Error("Parking lot is not created")
	}
}

func TestCreateEmptyLotsInParkingLot(t *testing.T) {
	_, err := NewParkingLot(10)
	if err != nil {
		t.Error("Empty lots are not created in ParkingLot")
	}
}

func TestCannotCreateParkingLotWithCapacityLessThanOne(t *testing.T) {
	_, err := NewParkingLot(-1)
	if err == nil {
		t.Error("Parking lot cannot be created with capacity less than 1")
	}
}

func TestParkCar(t *testing.T) {
	lot := Park("KK-10-AA-1234")
	if lot.NumberPlate != "KK-10-AA-1234" {
		t.Errorf("Vehicle not parked")
	}
}

func TestCreateParkingLotStruct(t *testing.T) {
	_, err := NewParkingLot((10))
	if err != nil {
		t.Errorf("Parking Lot is not created")
	}

}
