package parkinglot

import (
	"errors"
	"math"
)

type ParkingType string

type choosenLot func(*Attendant) (*ParkingLot, error)

const (
	ParkAtFirstEmptyLot  ParkingType = "park in the first empty parkinglot"
	ParkInLeastFilledLot ParkingType = "park in the parkinglot with least number of cars"
	ParkInMaxCapacityLot ParkingType = "park in the parkinglot with maximum capacity"
)

type Attendant struct {
	Parkinglots     []*ParkingLot
	parkingStatuses []bool
	lotchoice       choosenLot
}

func (a *Attendant) findLotWithMaxCapacity() (*ParkingLot, error) {
	maxCapacity := 0
	var parkingLotWithMaxCapacity *ParkingLot
	for _, p := range a.Parkinglots {
		if !p.isFullyFilled() && p.capacity > maxCapacity {
			maxCapacity = p.capacity
			parkingLotWithMaxCapacity = p
		}
	}
	if parkingLotWithMaxCapacity == nil {
		return nil, errors.New("parkinglot: findLotWithMaxCapacity: all lots are full")
	}
	return parkingLotWithMaxCapacity, nil
}

func NewAttendant(choice ParkingType, parkingLots ...*ParkingLot) (*Attendant, error) {
	parkinglotSlice := []*ParkingLot{}
	for _, parkingLot := range parkingLots {
		if parkingLot == nil {
			return nil, errors.New("parkinglot: NewAttendant: attendant cannot have nil parkinglot")
		}
	}

	parkinglotSlice = append(parkinglotSlice, parkingLots...)

	statuses := make([]bool, len(parkingLots))
	attendant := Attendant{
		Parkinglots:     parkingLots,
		parkingStatuses: statuses,
	}

	for _, parkinglot := range parkinglotSlice {
		parkinglot.addParkingFullReceiver(&attendant)
		parkinglot.setParkingAvailableReceiver(&attendant)
	}

	var lot choosenLot

	switch choice {

	case ParkInLeastFilledLot:
		lot = (*Attendant).findLeastCarsLot

	case ParkInMaxCapacityLot:
		lot = (*Attendant).findLotWithMaxCapacity

	case ParkAtFirstEmptyLot:
		lot = (*Attendant).firstemptylot

	default:
		return nil, errors.New("unknown parking type")
	}
	attendant.lotchoice = lot

	return &attendant, nil
}

func (a *Attendant) firstemptylot() (*ParkingLot, error) {
	for i, p := range a.Parkinglots {
		if a.parkingStatuses[i] {
			continue
		}
		return p, nil
	}
	return nil, errors.New("parkinglot: firstemptylot: all lots are full")
}

func (a *Attendant) Park(car *Car) error {
	if car == nil {
		return errors.New("parkinglot: park (by attendant): car cannot be nil")
	}

	if a.checkIsCarParked(car) {
		return errors.New("parkinglot: park (by attendant): car already parked")
	}

	lot, err := a.lotchoice(a)

	if err != nil {
		return errors.New("parkinglot: park (by attendant): all parkinglots are full")
	}

	return lot.park(car)

}

func (a *Attendant) UnPark(car *Car) error {
	if !a.checkIsCarParked(car) {
		return errors.New("parkinglot: unpark (by attendant): car is not parked")
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

func (a *Attendant) findLeastCarsLot() (*ParkingLot, error) {
	var targetLot *ParkingLot
	minCars := math.MaxInt
	id := -1

	for i, lot := range a.Parkinglots {

		if a.parkingStatuses[i] {
			continue
		}
		if lot.CarsParkedCount() < minCars {
			minCars = lot.CarsParkedCount()
			targetLot = lot
			id = i
		}

	}
	if id == -1 {
		return nil, errors.New("parkinglot: findLeastCarsLot: all parkinglots are full")
	}
	return targetLot, nil
}

func (a *Attendant) receiveAvailable(p *ParkingLot) {
	for i, l := range a.Parkinglots {
		if l == p {
			a.parkingStatuses[i] = false
			break
		}
	}
}
