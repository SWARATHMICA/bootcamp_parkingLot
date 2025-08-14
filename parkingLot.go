package parkinglot

import (
	"errors"
)

type ParkingFullReceiver interface {
	receiveFull()
}

type ParkingAvailableReceiver interface {
	receiveAvailable()
}

type slot struct {
	id       int
	car      *Car
	occupied bool
}

func (s slot) isEmpty() bool {
	return !s.occupied
}

func (s *slot) occupy(c *Car) {
	s.car = c
	s.occupied = true
}

func (s *slot) free() {
	s.car = nil
	s.occupied = false
}

type ParkingLot struct {
	capacity          int
	slots             []slot
	fullReceiver      []ParkingFullReceiver
	availableReceiver ParkingAvailableReceiver
	isFull            bool
}

func (p *ParkingLot) setParkingAvailableReceiver(parkingAvailableReceiver ParkingAvailableReceiver) {
	p.availableReceiver = parkingAvailableReceiver
}

func (p *ParkingLot) addParkingFullReceiver(r ParkingFullReceiver) {
	p.fullReceiver = append(p.fullReceiver, r)
}

type Car struct {
	numberPlate string
}

func (car1 Car) isEqual(car2 Car) bool {
	return car1.numberPlate == car2.numberPlate
}

func NewParkingLot(capacity int) (*ParkingLot, error) {
	if capacity < 1 {
		return nil, errors.New("cannot create parking lot with capacity less than 1")
	}
	lots := make([]slot, 0, capacity)
	for i := 0; i < capacity; i++ {
		newLot := slot{
			id:       i + 1,
			car:      nil,
			occupied: false,
		}
		lots = append(lots, newLot)
	}

	return &ParkingLot{slots: lots, capacity: capacity}, nil
}

func (p *ParkingLot) park(c *Car) (bool, error) {
	if c == nil {
		return false, errors.New("park: car cannot be nil")
	}
	if p.isParked(*c) {
		return false, errors.New("Car already parked")
	}
	for i := 0; i < p.capacity; i++ {
		if !p.slots[i].isEmpty() {
			continue
		}
		p.slots[i].occupy(c)

		if i == p.capacity-1 {
			p.isFull = true
			if p.fullReceiver != nil {
				p.notifyReceiver()

			}

		}

		return true, nil

	}
	return false, errors.New("ParkingLot is full")

}

func (p *ParkingLot) unPark(car *Car) (bool, error) {
	if car == nil {
		return false, errors.New("unpark: car cannot be nil")
	}
	for i := 0; i < p.capacity; i++ {
		if p.isNotParked(*car) {
			continue
		}

		p.slots[i].free()

		if p.isNotFull() {
			continue
		}
		p.isFull = false

		p.notifyAvailableReciever()

		return true, nil
	}
	return false, errors.New("Car is not found in the parking lot")
}

func (p *ParkingLot) isNotFull() bool {
	return !p.isFull
}

func (p *ParkingLot) notifyAvailableReciever() {
	if p.availableReceiver != nil {
		p.availableReceiver.receiveAvailable()
	}
}

func (p *ParkingLot) isParked(car Car) bool {
	for _, slot := range p.slots {
		if slot.isEmpty() {
			continue
		}
		if car.isEqual(*slot.car) {
			return true
		}
	}
	return false
}

func (p *ParkingLot) notifyReceiver() {
	for _, r := range p.fullReceiver {
		r.receiveFull()
	}

}

func (p *ParkingLot) isNotParked(car Car) bool {
	return !p.isParked(car)
}
