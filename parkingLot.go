package parkinglot

type Lot struct {
	Id          int
	NumberPlate string
	Occupied    bool
}

func NewParkingLot(capacity int) ([]Lot, error) {

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
