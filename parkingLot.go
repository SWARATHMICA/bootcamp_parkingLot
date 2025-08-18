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

func (s *slot) isEmpty() bool {
	return !s.occupied
}
func (s *slot) isNotEmpty() bool {
	return s.occupied
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
	capacity int
	slots    []slot
	//TODO reveal intention
	fullReceiver      []ParkingFullReceiver
	availableReceiver ParkingAvailableReceiver
	//TODO fewest elements
	isFull bool
}

// TODO idiomatic go
func (p *ParkingLot) setParkingAvailableReceiver(parkingAvailableReceiver ParkingAvailableReceiver) {
	p.availableReceiver = parkingAvailableReceiver
}

func (p *ParkingLot) addParkingFullReceiver(r ParkingFullReceiver) {
	p.fullReceiver = append(p.fullReceiver, r)
}

type Car struct {
	numberPlate string
}

func (car1 *Car) isEqual(car2 *Car) bool {
	return car1.numberPlate == car2.numberPlate
}

func NewParkingLot(capacity int) (*ParkingLot, error) {
	if capacity < 1 {
		return nil, errors.New("cannot create parking lot with capacity less than 1")
	}
	lots := make([]slot, 0, capacity)
	for i := 0; i < capacity; i++ {
		newLot := slot{
			id: i + 1,
		}
		lots = append(lots, newLot)
	}

	return &ParkingLot{slots: lots, capacity: capacity}, nil
}

// TODO fewest elements
func (p *ParkingLot) park(c *Car) (bool, error) {
	if c == nil {
		return false, errors.New("park: car cannot be nil")
	}
	if p.isParked(c) {
		return false, errors.New("Car already parked")
	}
	for i := 0; i < p.capacity; i++ {
		if p.slots[i].isNotEmpty() {
			continue
		}
		p.slots[i].occupy(c)
		//TODO: Refactor : Correct indentation
		if p.isFullyFilled() {
			p.isFull = true
			if p.fullReceiver != nil {
				p.notifyReceiver()

			}

		}

		return true, nil

	}
	return false, errors.New("ParkingLot is full")

}

func (p *ParkingLot) isFullyFilled() bool {
	for _, slot := range p.slots {
		if !slot.occupied {
			return false
		}
	}
	return true
}

func (p *ParkingLot) unPark(car *Car) (bool, error) {
	if car == nil {
		return false, errors.New("unpark: car cannot be nil")
	}
	for i := 0; i < p.capacity; i++ {
		if p.slots[i].isEmpty() {
			continue
		}

		if p.slots[i].car.isEqual(car) {
			p.slots[i].free()

			p.isFull = false
			//TODO: check coupling and correct indentation
			if !p.isFullyFilled() {
				p.notifyAvailableReciever()
			}
			return true, nil
		}

	}
	return false, errors.New("Car is not found in the parking lot")
}

func (p *ParkingLot) notifyAvailableReciever() {
	if p.availableReceiver != nil {
		p.availableReceiver.receiveAvailable()
	}
}

func (p *ParkingLot) isParked(car *Car) bool {
	for _, slot := range p.slots {
		if slot.isEmpty() {
			continue
		}
		if car.isEqual(slot.car) {
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
