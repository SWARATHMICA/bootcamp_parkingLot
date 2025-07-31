package parkinglot

import "testing"

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
