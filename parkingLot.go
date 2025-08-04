package parkinglot

import (
	"errors"
)

type observer interface {
	notify()
}

type slot struct {
	id          int
	numberPlate string
	occupied    bool
}

type ParkingLot struct {
	capacity int
	slots    []slot
	observer observer
	isFull   bool
}

func (p *ParkingLot) setObserver(observer observer) {
	p.observer = observer
}

type Car struct {
	numberPlate string
}

func NewParkingLot(capacity int) (*ParkingLot, error) {
	if capacity < 1 {
		return nil, errors.New("cannot create parking lot with capacity less than 1")
	}
	lots := make([]slot, 0, capacity)
	for i := 0; i < capacity; i++ {
		newLot := slot{
			id:          i + 1,
			numberPlate: "",
			occupied:    false,
		}
		lots = append(lots, newLot)
	}

	return &ParkingLot{slots: lots, capacity: capacity}, nil
}

func (p *ParkingLot) Park(c Car) (bool, error) {
	if p.IsParked(c) {
		return false, errors.New("Car already parked")
	}
	for i := 0; i < p.capacity; i++ {
		if !p.slots[i].occupied {
			p.slots[i].numberPlate = c.numberPlate
			p.slots[i].occupied = true
			if i == p.capacity-1 {
				p.isFull = true
				p.observer.notify()
			}
			return true, nil

		}
	}
	return false, errors.New("ParkingLot is full")

}

func (p *ParkingLot) Unpark(car Car) (bool, error) {
	for i := 0; i < p.capacity; i++ {
		if p.IsParked(car) {
			p.slots[i].numberPlate = ""
			p.slots[i].occupied = false
			if p.isFull {
				p.isFull = false
				p.observer.notify()
			}
			return true, nil
		}
	}
	return false, errors.New("Car is not found in the parking lot")
}

func (p *ParkingLot) IsParked(car Car) bool {
	for _, slot := range p.slots {
		if slot.occupied && slot.numberPlate == car.numberPlate {
			return true
		}
	}
	return false
}

func (p *ParkingLot) notify() string {
	if p.isFull {
		message := "The parking lot is full"
		return message
	}
	return "Parking lot has space again. Please remove the sign."

}
