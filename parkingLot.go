package parkinglot

import "errors"

type Lot struct {
	Id          int
	NumberPlate string
	Occupied    bool
}

type ParkingLot struct {
	slots []Lot
}

func NewParkingLot(capacity int) (*ParkingLot, error) {
	if capacity < 1 {
		return nil, errors.New("cannot create parking lot with capacity less than 1")
	}
	lots := make([]Lot, 0, capacity)
	for i := 0; i < capacity; i++ {
		newLot := Lot{
			Id:          i + 1,
			NumberPlate: "",
			Occupied:    false,
		}
		lots = append(lots, newLot)
	}

	return &ParkingLot{lots}, nil
}

func Park(s string) Lot {
	parkinglot, _ := NewParkingLot(10)
	parkinglot.slots[2].NumberPlate = s
	parkinglot.slots[2].Occupied = true

	return parkinglot.slots[2]
}
