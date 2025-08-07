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
	Parkinglots []*ParkingLot
}

type slot struct {
	id          int
	numberPlate string
	occupied    bool
}

type ParkingLot struct {
	capacity          int
	slots             []slot
	fullReceiver      []ParkingFullReceiver
	availableReceiver ParkingAvailableReceiver
	isFull            bool
}

func (a *Attendant) addParkingLotInAttendent(p *ParkingLot) {
	a.Parkinglots = append(a.Parkinglots, p)
}

func (p *ParkingLot) setParkingAvailableReceiver(parkingAvailableReceiver ParkingAvailableReceiver) {
	p.availableReceiver = parkingAvailableReceiver
}

func (p *ParkingLot) addParkingFullReceiver(r ParkingFullReceiver) {
	p.fullReceiver = append(p.fullReceiver, r)
}

func (p *ParkingLot) setReceiver(receiver ParkingFullReceiver) {
	p.fullReceiver = append(p.fullReceiver, receiver)
}

type Car struct {
	numberPlate string
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
				if p.fullReceiver != nil {
					p.notifyReceiver()

				}

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
				if p.availableReceiver != nil {
					p.availableReceiver.receiveAvailable()
				}

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

func (p *ParkingLot) notifyReceiver() {
	for _, r := range p.fullReceiver {
		r.receiveFull()
	}

}

func (a *Attendant) FindAndPark(car Car) bool {
	for _, p := range a.Parkinglots {
		if !p.isFull {
			result, _ := p.Park(car)
			return result
		}
	}
	return false
}

func (a *Attendant) FindAndUnPark(car Car) bool {
	for _, p := range a.Parkinglots {
		result, _ := p.Unpark(car)
		if result {
			return true
		}
	}
	return false
}
