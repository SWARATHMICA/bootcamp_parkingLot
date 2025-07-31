package parkinglot

import "errors"

type Lot struct {
	Id          int
	NumberPlate string
	Occupied    bool
}

func NewParkingLot(capacity int) ([]Lot, error) {
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

	return lots, nil
}

func Park(s string) Lot {
	lots, _ := NewParkingLot(10)
	lots[2].NumberPlate = s
	lots[2].Occupied = true

	return lots[2]
}
