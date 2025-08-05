package parkinglot

import (
	"errors"
)

type ParkingStatusReceiver interface {
	receive(status ParkingStatus)
}

type slot struct {
	id          int
	numberPlate string
	occupied    bool
}

type ParkingLot struct {
	capacity int
	slots    []slot
	receiver ParkingStatusReceiver
	isFull   bool
}

func (p *ParkingLot) setReceiver(receiver ParkingStatusReceiver) {
	p.receiver = receiver
}

type Car struct {
	numberPlate string
}

type ParkingStatus int

const (
	UNKNOWN_STATUS ParkingStatus = iota
	PARKING_FULL
	PARKING_AVAILABLE
)

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
				p.notifyReceiver(PARKING_FULL)
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
				p.notifyReceiver(PARKING_AVAILABLE)
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

func (p *ParkingLot) notifyReceiver(status ParkingStatus) {
	p.receiver.receive(status)
}
