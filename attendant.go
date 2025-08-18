package parkinglot

import "errors"

type Attendant struct {
	Parkinglot  []*ParkingLot
	parkingFull bool
}

func NewAttendant(parkingLots ...*ParkingLot) (*Attendant, error) {
	parkinglotSlice := []*ParkingLot{}

	for _, parkingLot := range parkingLots {
		if parkingLot == nil {
			return nil, errors.New("attendant cannot have nil parkinglot")
		}
	}

	parkinglotSlice = append(parkinglotSlice, parkingLots...)

	attendant := Attendant{
		Parkinglot:  parkinglotSlice,
		parkingFull: false,
	}

	for _, parkinglot := range parkinglotSlice {
		parkinglot.OnFull(&attendant)
	}

	return &attendant, nil
}

func (a *Attendant) Park(car *Car) error {
	if car == nil {
		return errors.New("car cannot be nil")
	}

	if a.checkIsCarParked(car) {
		return errors.New("attendant: car already parked")
	}

	if !a.parkingFull {
		for _, p := range a.Parkinglot {
			if !p.isFullyFilled() {
				p.park(car)
				return nil
			}

		}
	}

	return errors.New("parking lot is full, attendant cannot park the car")
}
func (a *Attendant) UnPark(car *Car) error {
	if !a.checkIsCarParked(car) {
		return errors.New("attendant/unpark: car is not parked")
	}

	var err error
	for _, parkinglot := range a.Parkinglot {
		if !parkinglot.isParked(car) {
			continue
		}
		err = parkinglot.unPark(car)
		a.parkingFull = false
	}

	return err
}

func (a *Attendant) receiveFull() {
	count := 0
	for _, parkinglot := range a.Parkinglot {
		if parkinglot.isFullyFilled() {
			count++
		}
	}
	if count == len(a.Parkinglot) {
		a.parkingFull = true
	}
}

func (a *Attendant) checkIsCarParked(car *Car) bool {
	for _, parkinglot := range a.Parkinglot {
		if parkinglot.isParked(car) {
			return true
		}
	}
	return false
}
