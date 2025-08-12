package parkinglot

import "errors"

type Attendant struct {
	Parkinglot  *ParkingLot
	parkingFull bool
}

func NewAttendant(parkingLot *ParkingLot) (*Attendant, error) {

	if parkingLot == nil {
		return nil, errors.New("cannot create attedant with nil parkinglot")
	}
	a := Attendant{
		Parkinglot:  parkingLot,
		parkingFull: false,
	}
	parkingLot.addParkingFullReceiver(&a)

	return &a, nil
}

func (a *Attendant) Park(car *Car) (bool, error) {
	if !a.parkingFull {
		return a.Parkinglot.park(car)
	}

	return false, errors.New("parking lot is full, attendant cannot park the car")
}

func (a *Attendant) UnPark(car *Car) (bool, error) {
	result, err := a.Parkinglot.unPark(car)
	if result {
		a.parkingFull = false
	}
	return result, err
}

func (a *Attendant) receiveFull() {
	a.parkingFull = true
}
