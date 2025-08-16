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
		parkinglot.addParkingFullReceiver(&attendant)
	}

	return &attendant, nil
}

// TODO fewest elements
func (a *Attendant) Park(car *Car) (bool, error) {
	if car == nil {
		return false, errors.New("car cannot be nil")
	}

	if a.checkIsCarParked(*car) {
		return false, errors.New("attendant: car already parked")
	}

	if !a.parkingFull {
		for _, p := range a.Parkinglot {
			if !p.isFull {
				p.park(car)
				return true, nil
			}

		}
	}

	return false, errors.New("parking lot is full, attendant cannot park the car")
}
func (a *Attendant) UnPark(car *Car) (bool, error) {
	result, err := a.Parkinglot[0].unPark(car)
	if result {
		a.parkingFull = false
	}
	return result, err
}

func (a *Attendant) receiveFull() {
	count := 0
	for _, parkinglot := range a.Parkinglot {
		if parkinglot.isFull {
			count++
		}
	}
	if count == len(a.Parkinglot) {
		a.parkingFull = true
	}
}

func (a *Attendant) checkIsCarParked(car Car) bool {
	for _, parkinglot := range a.Parkinglot {
		if parkinglot.isParked(car) {
			return true
		}
	}
	return false
}
