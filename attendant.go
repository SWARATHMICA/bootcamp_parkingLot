package parkinglot

import (
	"errors"
	"math"
)

type ParkingType string

const (
	SimpleParking      ParkingType = "simpleParking"
	EvenParking        ParkingType = "evenParking"
	MaxCapacityParking ParkingType = "maxCapacityParking"
)

type Attendant struct {
	Parkinglots     []*ParkingLot
	parkingStatuses []bool
	choice          ParkingType
}

func (a *Attendant) findLotWithMaxCapacity() *ParkingLot {
	maxCapacity := 0
	var parkingLotWithMaxCapacity *ParkingLot
	for _, p := range a.Parkinglots {
		if !p.isFullyFilled() && p.capacity > maxCapacity {
			maxCapacity = p.capacity
			parkingLotWithMaxCapacity = p
		}
	}
	return parkingLotWithMaxCapacity
}

func NewAttendant(choice ParkingType, parkingLots ...*ParkingLot) (*Attendant, error) {
	parkinglotSlice := []*ParkingLot{}
	for _, parkingLot := range parkingLots {
		if parkingLot == nil {
			return nil, errors.New("parkinglot:NewAttendant:attendant cannot have nil parkinglot")
		}
	}

	parkinglotSlice = append(parkinglotSlice, parkingLots...)

	statuses := make([]bool, len(parkingLots))
	attendant := Attendant{
		Parkinglots:     parkingLots,
		parkingStatuses: statuses,
		choice:          choice,
	}

	for _, parkinglot := range parkinglotSlice {
		parkinglot.addParkingFullReceiver(&attendant)
		parkinglot.setParkingAvailableReceiver(&attendant)
	}

	return &attendant, nil
}

func (a *Attendant) lotBasedOnChoice() *ParkingLot {

	if a.choice == EvenParking {
		return a.findLeastCarsLot()
	}
	if a.choice == MaxCapacityParking {
		return a.findLotWithMaxCapacity()
	}
	return a.firstemptylot()
}

func (a *Attendant) firstemptylot() *ParkingLot {
	for i, p := range a.Parkinglots {
		if a.parkingStatuses[i] {
			continue
		}
		return p
	}
	return nil
}

func (a *Attendant) Park(car *Car) error {
	if car == nil {
		return errors.New("parkinglot:park (by attendant):car cannot be nil")
	}

	if a.checkIsCarParked(car) {
		return errors.New("parkinglot:park (by attendant): car already parked")
	}
	lot := a.lotBasedOnChoice()
	if lot == nil {
		return errors.New("parkinglot: park (by attendant): all parkinglots are full")
	}

	return lot.park(car)

}

func (a *Attendant) UnPark(car *Car) error {
	if !a.checkIsCarParked(car) {
		return errors.New("parkinglot:unpark (by attendant): car is not parked")
	}

	var err error
	for i, parkinglot := range a.Parkinglots {
		if !parkinglot.isParked(car) {
			continue
		}
		err := parkinglot.unPark(car)
		if err != nil {
			return err
		}
		a.parkingStatuses[i] = false
		return nil
	}

	return err
}

func (a *Attendant) receiveFull(p *ParkingLot) {
	for i, l := range a.Parkinglots {
		if l == p {
			a.parkingStatuses[i] = true
			break
		}
	}

}

func (a *Attendant) checkIsCarParked(car *Car) bool {
	for _, parkinglot := range a.Parkinglots {
		if parkinglot.isParked(car) {
			return true
		}
	}
	return false
}

func (a *Attendant) findLeastCarsLot() *ParkingLot {
	var targetLot *ParkingLot
	minCars := math.MaxInt

	for i, lot := range a.Parkinglots {

		if a.parkingStatuses[i] {
			continue
		}
		if lot.CarsParkedCount() < minCars {
			minCars = lot.CarsParkedCount()
			targetLot = lot
		}

	}
	return targetLot
}

func (a *Attendant) receiveAvailable(p *ParkingLot) {
	for i, l := range a.Parkinglots {
		if l == p {
			a.parkingStatuses[i] = false
			break
		}
	}
}
