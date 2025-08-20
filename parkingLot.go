package parkinglot

import (
	"errors"
)

type ParkingFullReceiver interface {
	receiveFull(*ParkingLot)
}

type ParkingAvailableReceiver interface {
	receiveAvailable(*ParkingLot)
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
	capacity          int
	slots             []slot
	fullReceivers     []ParkingFullReceiver
	availableReceiver ParkingAvailableReceiver
}

func (p *ParkingLot) setParkingAvailableReceiver(parkingAvailableReceiver ParkingAvailableReceiver) {
	p.availableReceiver = parkingAvailableReceiver
}

func (p *ParkingLot) addParkingFullReceiver(r ParkingFullReceiver) {
	p.fullReceivers = append(p.fullReceivers, r)
}

type Car struct {
	numberPlate string
}

func (car1 *Car) isEqual(car2 *Car) bool {
	return car1.numberPlate == car2.numberPlate
}

func NewParkingLot(capacity int) (*ParkingLot, error) {
	if capacity < 1 {
		return nil, errors.New("parkinglot: NewParkingLot: cannot create parking lot with capacity less than 1")
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

func (p *ParkingLot) park(c *Car) error {
	if c == nil {
		return errors.New("parkinglot: park: car cannot be nil")
	}
	if p.isParked(c) {
		return errors.New("parkinglot: park: car already parked")
	}
	for i := 0; i < p.capacity; i++ {
		if p.slots[i].isNotEmpty() {
			continue
		}
		p.slots[i].occupy(c)

		if p.isFullyFilled() {

			p.notifyReceiver()

		}

		return nil

	}
	return errors.New("parkinglot: park: parkingLot is full")

}

func (p *ParkingLot) isFullyFilled() bool {
	for _, slot := range p.slots {
		if !slot.occupied {
			return false
		}
	}
	return true
}

func (p *ParkingLot) unPark(car *Car) error {
	if car == nil {
		return errors.New("parkinglot: unpark: car cannot be nil")
	}
	for i := 0; i < p.capacity; i++ {
		if p.slots[i].isEmpty() {
			continue
		}

		if p.slots[i].car.isEqual(car) {
			p.slots[i].free()

			if p.isFullyFilled() {
				continue
			}
			p.notifyAvailableReciever()
			return nil
		}

	}
	return errors.New("parkinglot: unpark: car is not found in the parking lot")
}

func (p *ParkingLot) notifyAvailableReciever() {
	if p.availableReceiver != nil {
		p.availableReceiver.receiveAvailable(p)
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
	for _, r := range p.fullReceivers {
		r.receiveFull(p)
	}

}

func (lot *ParkingLot) CarsParkedCount() int {
	count := 0
	for _, slot := range lot.slots {
		if slot.car != nil {
			count++
		}
	}
	return count
}
