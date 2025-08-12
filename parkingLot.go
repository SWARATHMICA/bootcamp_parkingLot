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

type Attendant struct {
	Parkinglot *ParkingLot
	notified   bool
}

type slot struct {
	id       int
	car      *Car
	occupied bool
}

type ParkingLot struct {
	capacity          int
	slots             []slot
	fullReceiver      []ParkingFullReceiver
	availableReceiver ParkingAvailableReceiver
	isFull            bool
}

func (a *Attendant) addParkingLotInAttendent(p *ParkingLot) {
	a.Parkinglot = p
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

type ParkingStatus int

const (
	//TODO: connvert to mixed case acc to golang convention
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
			id:       i + 1,
			car:      nil,
			occupied: false,
		}
		lots = append(lots, newLot)
	}

	return &ParkingLot{slots: lots, capacity: capacity}, nil
}

func (p *ParkingLot) park(c *Car) (bool, error) {
	if p.isParked(*c) {
		return false, errors.New("Car already parked")
	}
	for i := 0; i < p.capacity; i++ {
		if !p.slots[i].occupied {
			p.slots[i].car = c
			p.slots[i].occupied = true
			if i == p.capacity-1 {
				p.isFull = true
				if p.fullReceiver != nil {
					p.notifyReceiver()

				}

			}
			return true, nil

		}
	}
	return false, errors.New("ParkingLot is full")

}

func (p *ParkingLot) unPark(car *Car) (bool, error) {
	for i := 0; i < p.capacity; i++ {
		if p.isParked(*car) {
			p.slots[i].car = nil
			p.slots[i].occupied = false
			if p.isFull {
				p.isFull = false
				if p.availableReceiver != nil {
					p.availableReceiver.receiveAvailable()
				}

			}
			return true, nil
		}
	}
	return false, errors.New("Car is not found in the parking lot")
}

func (p *ParkingLot) isParked(car Car) bool {
	for _, slot := range p.slots {
		if slot.occupied && slot.car.isEqual(car) {
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

func (a *Attendant) ParkWithAttendant(car *Car) (bool, error) {
	if !a.notified {
		return a.Parkinglot.park(car)
	}

	return false, errors.New("parking lot is full, attendant cannot park the car")
}

func (a *Attendant) UnParkWithAttendant(car *Car) (bool, error) {
	return a.Parkinglot.unPark(car)
}

func (a *Attendant) receiveFull() {
	a.notified = true
}

func NewAttendant(parkingLot *ParkingLot) (*Attendant, error) {

	if parkingLot == nil {
		return nil, errors.New("cannot create attedant with nil parkinglot")
	}
	a := Attendant{
		Parkinglot: parkingLot,
		notified:   false,
	}
	parkingLot.addParkingFullReceiver(&a)

	return &a, nil
}
