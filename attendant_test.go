package parkinglot

import "testing"

func TestCreateNewAttendant(t *testing.T) {
	parkinglog, _ := NewParkingLot(1)
	_, err := NewAttendant(true, parkinglog)

	if err != nil {
		t.Error("new attendant should be created")
	}
}

func TestAttedantCannotBeCreatedWithNilParkingLot(t *testing.T) {
	_, err := NewAttendant(true, nil)

	const expectedError = "attendant cannot have nil parkinglot"

	if err.Error() != expectedError {
		t.Error("attendant should not create with nil parking lot")
	}
}

func TestParkCarByAttendant(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)
	attendant, _ := NewAttendant(true, parkingLot)

	err := attendant.Park(&car)

	if err != nil {
		t.Errorf("attendant should be able to park the car")
	}
}

func TestAttendantCannotParkWhenParkingFull(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)
	attendant, _ := NewAttendant(true, parkingLot)

	expectedError := "parking lot is full, attendant cannot park the car"

	attendant.Park(&Car{"KK10AA1234"})

	err := attendant.Park(&car)

	if err.Error() != expectedError {
		t.Errorf("attendant cannot park when parking full")
	}
}

func TestUnParkCarByAttendant(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)
	attendant, _ := NewAttendant(true, parkingLot)

	attendant.Park(&car)
	err := attendant.UnPark(&car)

	if err != nil {
		t.Errorf("car should be unparked by attendant")
	}
}

func TestAttendantParkAfterParkingAvailable(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)
	attendant, _ := NewAttendant(true, parkingLot)

	//park the car
	attendant.Park(&car)

	//unpark the car
	attendant.UnPark(&car)

	//should be able to park again
	err := attendant.Park(&car)

	if err != nil {
		t.Error("should be able to park after parking become available")
	}

}

//Multiple parking lots

func TestAttendantCannotParkNilCar(t *testing.T) {
	parkinglot, _ := NewParkingLot(2)
	attendant, _ := NewAttendant(true, parkinglot)

	err := attendant.Park(nil)
	expectedError := "car cannot be nil"

	if err.Error() != expectedError {
		t.Errorf("attendant cannot park nil car")
	}
}

func TestAttendantShouldCheckCarIsParkedAfterUnpark(t *testing.T) {
	parkinglot, _ := NewParkingLot(2)
	attendant, _ := NewAttendant(true, parkinglot)
	car2 := Car{"UM-12-TK-1234"}

	err := attendant.Park(&car)
	if err != nil {
		t.Fatal("car should be parked")
	}

	err = attendant.Park(&car2)
	if err != nil {
		t.Fatal("car2 should be parked")
	}
	err = attendant.UnPark(&car2)
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

func TestAttendantCanManageMultipleParkingLots(t *testing.T) {
	parkinglot1, err1 := NewParkingLot(1)
	if err1 != nil {
		t.Fatalf("failed to create parking lot 1: %v", err1)
	}

	parkinglot2, err2 := NewParkingLot(1)
	if err2 != nil {
		t.Fatalf("failed to create parking lot 2: %v", err2)
	}

	_, err := NewAttendant(true, parkinglot1, parkinglot2)

	if err != nil {
		t.Errorf("expected attendant to be created with multiple parking lots, got error: %v", err)
	}
}

func TestAttendantParkCarInNextParkingLotWithAvailableSlot(t *testing.T) {

	parkinglot1, _ := NewParkingLot(1)
	parkinglot2, _ := NewParkingLot(3)
	attendant, _ := NewAttendant(true, parkinglot1, parkinglot2)

	parkinglot1.park(&car)
	car2 := Car{"1234567"}
	err := attendant.Park(&car2)

	if err != nil {
		t.Fatal("car2 should be parked at parking lot 2")
	}

	if !parkinglot2.slots[0].car.isEqual(&car2) {
		t.Errorf("car2 should be parked at parking lot 2 ")

	}

}

func TestAttendantIsAbleToUnparkAfterParkForMultipleParkingLots(t *testing.T) {
	parkinglot1, _ := NewParkingLot(1)
	parkinglot2, _ := NewParkingLot(2)
	parkinglot3, _ := NewParkingLot(1)
	attendant, _ := NewAttendant(true, parkinglot1, parkinglot2, parkinglot3)

	err := parkinglot1.park(&car)
	if err != nil {
		t.Fatal("car should be parking in parkinglot 1")
	}

	anotherCar := &Car{"1234567"}
	err = parkinglot2.park(anotherCar)
	if err != nil {
		t.Fatal("another should be parking in parkinglot 2")
	}

	err = parkinglot3.park(&Car{"09876533"})
	if err != nil {
		t.Fatal("Car{09876533} should be parking in parkinglot 2")
	}

	err = attendant.UnPark(anotherCar)

	if err != nil {
		t.Error("another car should be unparked from parkinglot 2")
	}

}

func TestAttendantfindLeastCarsLot(t *testing.T) {
	parkinglot1, _ := NewParkingLot(2)
	parkinglot2, _ := NewParkingLot(2)
	parkinglot3, _ := NewParkingLot(2)
	attendant, _ := NewAttendant(true, parkinglot1, parkinglot2, parkinglot3)
	car3 := Car{"KK10AA2369"}
	car2 := Car{"KK10AA2345"}

	attendant.Park(&car)
	attendant.Park(&car2)
	attendant.Park(&car3)

	got := attendant.findLeastCarsLot()

	if got != parkinglot3 {
		t.Errorf("expected lot2 (fewer cars), got %+v", got)
	}
}
