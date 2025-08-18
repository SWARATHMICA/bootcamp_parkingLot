package parkinglot

import "testing"

func TestCreateNewAttendant(t *testing.T) {
	parkinglog, _ := NewParkingLot(1)
	_, err := NewAttendant(parkinglog)

	if err != nil {
		t.Error("new attendant should be created")
	}
}

func TestAttedantCannotBeCreatedWithNilParkingLot(t *testing.T) {
	_, err := NewAttendant(nil)

	const expectedError = "attendant cannot have nil parkinglot"

	if err.Error() != expectedError {
		t.Error("attendant should not create with nil parking lot")
	}
}

func TestParkCarByAttendant(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)
	attendant, _ := NewAttendant(parkingLot)

	_, err := attendant.Park(&car)

	if err != nil {
		t.Errorf("attendant should be able to park the car")
	}
}

func TestAttendantCannotParkWhenParkingFull(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)
	attendant, _ := NewAttendant(parkingLot)

	expectedError := "parking lot is full, attendant cannot park the car"

	attendant.Park(&Car{"KK10AA1234"})

	_, err := attendant.Park(&car)

	if err.Error() != expectedError {
		t.Errorf("attendant cannot park when parking full")
	}
}

func TestUnParkCarByAttendant(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)
	attendant, _ := NewAttendant(parkingLot)

	attendant.Park(&car)
	_, err := attendant.UnPark(&car)

	if err != nil {
		t.Errorf("car should be unparked by attendant")
	}
}

func TestAttendantParkAfterParkingAvailable(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)
	attendant, _ := NewAttendant(parkingLot)

	//park the car
	attendant.Park(&car)

	//unpark the car
	attendant.UnPark(&car)

	//should be able to park again
	_, err := attendant.Park(&car)

	if err != nil {
		t.Error("should be able to park after parking become available")
	}

}

func TestAttendantReceiveFullNotification(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)
	attendant, _ := NewAttendant(parkingLot)

	attendant.Park(&car)

	if !attendant.parkingFull {
		t.Errorf("attendant should be notified about parking full")
	}
}

//Multiple parking lots

func TestAttendantCannotParkNilCar(t *testing.T) {
	parkinglot, _ := NewParkingLot(2)
	attendant, _ := NewAttendant(parkinglot)

	_, err := attendant.Park(nil)
	expectedError := "car cannot be nil"

	if err.Error() != expectedError {
		t.Errorf("attendant cannot park nil car")
	}
}

func TestAttendantShouldCheckIfCarAlreadyParked(t *testing.T) {
	parkinglot, _ := NewParkingLot(2)
	attendant, _ := NewAttendant(parkinglot)
	attendant.Park(&car)
	isParked := attendant.checkIsCarParked(&car)
	if !isParked {
		t.Errorf("car is already parked")
	}
}

func TestAttendantShouldCheckCarIsParkedAfterUnpark(t *testing.T) {
	parkinglot, _ := NewParkingLot(2)
	attendant, _ := NewAttendant(parkinglot)
	car2 := Car{"UM-12-TK-1234"}

	_, err := attendant.Park(&car)
	if err != nil {
		t.Fatal("car should be parked")
	}

	_, err = attendant.Park(&car2)
	if err != nil {
		t.Fatal("car2 should be parked")
	}
	_, err = attendant.UnPark(&car2)
	if err != nil {
		t.Fatal("car2 should be unparked")
	}

	isCarParked := attendant.checkIsCarParked(&car)
	isCar2Parked := attendant.checkIsCarParked(&car2)

	if isCarParked == false {
		t.Error("car should be parked in parking lot")
	}
	if isCar2Parked == true {
		t.Errorf("car is already parked")
	}
}

func TestAttendantCanAcceptMultipleParkingLot(t *testing.T) {
	parkinglot1, err1 := NewParkingLot(1)
	if err1 != nil {
		t.Fatalf("failed to create parking lot 1: %v", err1)
	}

	parkinglot2, err2 := NewParkingLot(1)
	if err2 != nil {
		t.Fatalf("failed to create parking lot 2: %v", err2)
	}

	_, err := NewAttendant(parkinglot1, parkinglot2)

	if err != nil {
		t.Errorf("expected attendant to be created with multiple parking lots, got error: %v", err)
	}
}
func TestCheckIsCarParkedWhenMultipleParkingLotAvailable(t *testing.T) {
	parkinglot1, _ := NewParkingLot(2)
	parkinglot2, _ := NewParkingLot(3)

	_, err := parkinglot2.park(&car)

	if err != nil {
		t.Fatal("car should be parked in the parking log 2")
	}

	attendant, _ := NewAttendant(parkinglot1, parkinglot2)

	flag := attendant.checkIsCarParked(&car)

	if flag == false {
		t.Error("car should be found in parking lot 2")
	}
}

func TestAttendantParkCarInNextParkingLotWithAvailableSlot(t *testing.T) {
	parkkinglot1, _ := NewParkingLot(1)
	parkkinglot2, _ := NewParkingLot(3)
	attendant, _ := NewAttendant(parkkinglot1, parkkinglot2)

	parkkinglot1.park(&car)

	_, err := attendant.Park(&Car{"1234567"})

	if err != nil {
		t.Error("car should be parked at parking lot 1")
	}

	_, err = attendant.Park(&car)

	const expectedError = "attendant: car already parked"
	if err.Error() != expectedError {
		t.Error("attend should not park the already parked car")
	}
}

func TestAttendantReceiveFullNotificationWhenAllParkingLotIsFull(t *testing.T) {
	parkingLot1, _ := NewParkingLot(1)
	parkingLot2, _ := NewParkingLot(1)
	parkingLot3, _ := NewParkingLot(1)
	attendant, err := NewAttendant(parkingLot1, parkingLot2, parkingLot3)

	if err != nil {
		t.Fatal("attendant was not created with parkinglot1, parkinglot2, parkinglot3")
	}

	attendant.Park(&car)
	attendant.Park(&Car{"123456789"})
	attendant.Park(&Car{"098765432"})

	if !attendant.parkingFull {
		t.Errorf("attendant should be notified about parking full")
	}
}

func TestAttendantIsAbleToUnparkAfterParkForMultipleParkingLots(t *testing.T) {
	parkinglot1, _ := NewParkingLot(1)
	parkinglot2, _ := NewParkingLot(2)
	parkinglot3, _ := NewParkingLot(1)
	attendant, _ := NewAttendant(parkinglot1, parkinglot2, parkinglot3)

	_, err := parkinglot1.park(&car)
	if err != nil {
		t.Fatal("car should be parking in parkinglot 1")
	}

	anotherCar := &Car{"1234567"}
	_, err = parkinglot2.park(anotherCar)
	if err != nil {
		t.Fatal("another should be parking in parkinglot 2")
	}

	_, err = parkinglot3.park(&Car{"09876533"})
	if err != nil {
		t.Fatal("Car{09876533} should be parking in parkinglot 2")
	}

	check, err := attendant.UnPark(anotherCar)

	if err != nil {
		t.Error("another car should be unparked from parkinglot 2")
	}

	if check == false {
		t.Error("flag should be true after another unpark from parkinglot 2")
	}
}
