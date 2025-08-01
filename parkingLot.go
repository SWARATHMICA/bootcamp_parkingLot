package parkinglot

import "errors"

type Slot struct {
	Id          int
	NumberPlate string
	Occupied    bool
}

type ParkingLot struct {
	capacity int
	slots    []Slot
}

type Car struct {
	numberPlate string
}

func NewParkingLot(capacity int) (*ParkingLot, error) {
	if capacity < 1 {
		return nil, errors.New("cannot create parking lot with capacity less than 1")
	}
	lots := make([]Slot, 0, capacity)
	for i := 0; i < capacity; i++ {
		newLot := Slot{
			Id:          i + 1,
			NumberPlate: "",
			Occupied:    false,
		}
		lots = append(lots, newLot)
	}

	return &ParkingLot{slots: lots, capacity: capacity}, nil
}

func (p *ParkingLot) Park(c Car) (bool, error) {
	for i := 0; i < p.capacity; i++ {
		if !p.slots[i].Occupied {
			p.slots[i].NumberPlate = c.numberPlate
			p.slots[i].Occupied = true
			return true, nil
		}
	}
	return false, errors.New("ParkingLot is full")

}

func (p *ParkingLot) Unpark(car Car) (bool, error) {
	for i := 0; i < p.capacity; i++ {
		if p.IsParked(car) {
			p.slots[i].NumberPlate = ""
			p.slots[i].Occupied = false
			return true, nil
		}
	}
	return false, errors.New("Car is not found in the parking lot")
}

func (p *ParkingLot) IsParked(car Car) bool {
	for _, slot := range p.slots {
		if slot.Occupied && slot.NumberPlate == car.numberPlate {
			return true
		}
	}
	return false
}
