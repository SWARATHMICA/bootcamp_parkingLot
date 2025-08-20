package parkinglot

import (
	"errors"
	"math"
)

type ParkingType string

const (
	SimpleParking ParkingType = "simpleParking"
	EvenParking   ParkingType = "evenParking"
)

type Attendant struct {
	Parkinglots     []*ParkingLot
	parkingStatuses []bool
	choice          ParkingType
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

func (a *Attendant) Park(car *Car) error {
	if car == nil {
		return errors.New("parkinglot:park (by attendant):car cannot be nil")
	}

	if a.checkIsCarParked(car) {
		return errors.New("parkinglot:park (by attendant): car already parked")
	}
	if a.choice == EvenParking {
		lot := a.findLeastCarsLot()
		if lot == nil {
			return errors.New("parkinglot:park (choice 2):parking lot is full, attendant cannot park the car")
		}
		return lot.park(car)
	}
	for i, p := range a.Parkinglots {
		if a.parkingStatuses[i] {
			continue
		}
		err := p.park(car)
		if err == nil {
			return nil
		}
		return err
	}

	return errors.New("parkinglot: park (choice 1):parking lot is full, attendant cannot park the car")
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
