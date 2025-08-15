package parkinglot

import "errors"

type Attendant struct {
	Parkinglot  []*ParkingLot
	parkingFull bool
}

func NewAttendant(parkingLot ...*ParkingLot) (*Attendant, error) {
	parkinglotSlice := []*ParkingLot{}
	if len(parkingLot) == 0 {

		return &Attendant{Parkinglot: []*ParkingLot{}}, nil

	}
	for _, p := range parkingLot {

		parkinglotSlice = append(parkinglotSlice, p)
	}

	a := Attendant{
		Parkinglot:  parkinglotSlice,
		parkingFull: false,
	}
	for _, p := range parkinglotSlice {
		if p == nil {
			return nil, errors.New("attendant should have be created with nil parking lot")
		}

		p.addParkingFullReceiver(&a)
	}

	return &a, nil
}

// TODO fewest elements
func (a *Attendant) Park(car *Car) (bool, error) {
	if car == nil {
		return false, errors.New("car cannot be nil")
	}
	if !a.parkingFull {
		return a.Parkinglot[0].park(car)
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
	a.parkingFull = true
}

func (a *Attendant) checkCarIsParked(car Car) bool {
	return a.Parkinglot[0].isParked(car)
}
