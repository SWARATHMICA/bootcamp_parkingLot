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
