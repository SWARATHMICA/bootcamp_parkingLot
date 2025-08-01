package parkinglot

import (
	"errors"
)

type slot struct {
	id          int
	numberPlate string
	occupied    bool
}

type ParkingLot struct {
	capacity int
	slots    []slot
	observer observer
}

func (p *ParkingLot) setObserver(observer observer) {
	p.observer = observer
}

type observer interface {
	notify()
}

type Car struct {
	numberPlate string
}

type Owner struct {
	Name string
}

type Other struct {
	Name string
	Role string
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
				p.notifyObserver()
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

func (o *Owner) parkingLotFull() string {
	// o.notified = true
	return "we are closed"
}
func (p *ParkingLot) notifyObserver() {
	p.observer.notify()

}
