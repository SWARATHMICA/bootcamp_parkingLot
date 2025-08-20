package parkinglot

import "testing"

func TestCreateNewAttendant(t *testing.T) {
	parkinglog, _ := NewParkingLot(1)
	_, err := NewAttendant(ParkAtFirstEmptyLot, parkinglog)

	if err != nil {
		t.Error("new attendant should be created")
	}
}

func TestAttedantCannotBeCreatedWithNilParkingLot(t *testing.T) {
	_, err := NewAttendant(ParkAtFirstEmptyLot, nil)

	const expectedError = "parkinglot:NewAttendant:attendant cannot have nil parkinglot"

	if err.Error() != expectedError {
		t.Error("attendant should not create with nil parking lot")
	}
}

func TestParkCarByAttendant(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)
	attendant, _ := NewAttendant(ParkAtFirstEmptyLot, parkingLot)

	err := attendant.Park(&car)

	if err != nil {
		t.Errorf("attendant should be able to park the car")
	}
}

func TestAttendantCannotParkWhenParkingFull(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)
	attendant, _ := NewAttendant(ParkAtFirstEmptyLot, parkingLot)

	expectedError := "parkinglot: park (by attendant): all parkinglots are full"

	attendant.Park(&Car{"KK10AA1234"})

	err := attendant.Park(&car)

	if err.Error() != expectedError {
		t.Errorf("attendant cannot park when parking full")
	}
}

func TestUnParkCarByAttendant(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)
	attendant, _ := NewAttendant(ParkAtFirstEmptyLot, parkingLot)

	attendant.Park(&car)
	err := attendant.UnPark(&car)

	if err != nil {
		t.Errorf("car should be unparked by attendant")
	}
}

func TestAttendantParkAfterParkingAvailable(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)
	attendant, _ := NewAttendant(ParkAtFirstEmptyLot, parkingLot)

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
	attendant, _ := NewAttendant(ParkAtFirstEmptyLot, parkinglot)

	err := attendant.Park(nil)
	expectedError := "parkinglot:park (by attendant):car cannot be nil"

	if err.Error() != expectedError {
		t.Errorf("attendant cannot park nil car")
	}
}

func TestAttendantShouldCheckCarIsParkedAfterUnpark(t *testing.T) {
	parkinglot, _ := NewParkingLot(2)
	attendant, _ := NewAttendant(ParkAtFirstEmptyLot, parkinglot)
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

	_, err := NewAttendant(ParkAtFirstEmptyLot, parkinglot1, parkinglot2)

	if err != nil {
		t.Errorf("expected attendant to be created with multiple parking lots, got error: %v", err)
	}
}

func TestAttendantParkCarInNextParkingLotWithAvailableSlot(t *testing.T) {

	parkinglot1, _ := NewParkingLot(1)
	parkinglot2, _ := NewParkingLot(3)
	attendant, _ := NewAttendant(ParkAtFirstEmptyLot, parkinglot1, parkinglot2)

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
	attendant, _ := NewAttendant(ParkAtFirstEmptyLot, parkinglot1, parkinglot2, parkinglot3)

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
	attendant, _ := NewAttendant(ParkAtFirstEmptyLot, parkinglot1, parkinglot2, parkinglot3)
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

func TestAttendantParksInLotWithFewerCars(t *testing.T) {

	parkinglot1, _ := NewParkingLot(2)
	parkinglot2, _ := NewParkingLot(2)
	parkinglot3, _ := NewParkingLot(2)
	attendant, _ := NewAttendant(ParkInLeastFilledLot, parkinglot1, parkinglot2, parkinglot3)

	car1 := &Car{"KA01AA2345"}
	car2 := &Car{"KA02BB5678"}
	car3 := &Car{"TN10AA3085"}

	attendant.Park(car1)
	attendant.Park(car2)

	err := attendant.Park(car3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if parkinglot3.CarsParkedCount() != 1 {
		t.Errorf("car3 has to be parked in parkinglot3")
	}
}

func TestComplexAttendantParksInAccordingToDistributionAfterUnpark(t *testing.T) {
	lot1, _ := NewParkingLot(2)
	lot2, _ := NewParkingLot(2)
	car1 := &Car{"car1"}
	car2 := &Car{"car2"}
	car3 := &Car{"car3"}
	car4 := &Car{"car4"}

	simpleAttendant, _ := NewAttendant(ParkAtFirstEmptyLot, lot1, lot2)
	complexAttendant, _ := NewAttendant(ParkInLeastFilledLot, lot1, lot2)

	err := simpleAttendant.Park(car1)
	if err != nil {
		t.Fatalf("car1 has to be parked in parkinglot1 first slot")
	}

	err = simpleAttendant.Park(car2)
	if err != nil {
		t.Fatalf("car2 has to be parked in parkinglot1 second slot")
	}

	err = complexAttendant.Park(car3)
	if err != nil {
		t.Fatal("car3 has to be parked in parkinglot2 first slot")
	}
	err = simpleAttendant.UnPark(car1)
	if err != nil {
		t.Fatal("car1 has to be unparked from parkinglot1 first slot")
	}

	err = simpleAttendant.UnPark(car2)
	if err != nil {
		t.Fatal("car2 has to be unparked from parkinglot1 second slot")
	}

	err = complexAttendant.Park(car4)
	if err != nil {
		t.Fatal("car4 has to be parked in parkinglot1 first slot")
	}
	if car4.isEqual(lot1.slots[0].car) == false {
		t.Fatalf("car4 should have been parked in parkinglot1 first slot")
	}
}

func TestAttendantfindLotWithMaxCapacity(t *testing.T) {
	parkinglot1, _ := NewParkingLot(2)
	parkinglot2, _ := NewParkingLot(3)
	attendant, _ := NewAttendant(ParkAtFirstEmptyLot, parkinglot1, parkinglot2)

	lotWithMaxCapacity := attendant.findLotWithMaxCapacity()

	if lotWithMaxCapacity.capacity != 3 {
		t.Errorf("parkinglot2 has maximum capacity of 3")
	}

}

func TestAttendantParksInParkingLotWithMaximumCapacity(t *testing.T) {
	parkinglot1, _ := NewParkingLot(1)
	parkinglot2, _ := NewParkingLot(2)
	attendant, _ := NewAttendant(ParkInMaxCapacityLot, parkinglot1, parkinglot2)

	car1 := &Car{"KA01AA2345"}
	car2 := &Car{"KA02BB5678"}

	attendant.Park(car1)
	attendant.Park(car2)

	if parkinglot2.CarsParkedCount() != 2 {
		t.Errorf("both cars should be parked in park")
	}
}

func TestCombinationOfAllThreeTypesOfAttendants(t *testing.T) {
	parkinglot1, _ := NewParkingLot(1)
	parkinglot2, _ := NewParkingLot(2)
	parkinglot3, _ := NewParkingLot(3)

	attendantSimple, _ := NewAttendant(ParkAtFirstEmptyLot, parkinglot1, parkinglot2, parkinglot3)
	attendantEven, _ := NewAttendant(ParkInLeastFilledLot, parkinglot1, parkinglot2, parkinglot3)
	attendantMax, _ := NewAttendant(ParkInMaxCapacityLot, parkinglot1, parkinglot2, parkinglot3)

	car1 := &Car{"KA01AA2345"}
	car2 := &Car{"KA02BB5678"}
	car3 := &Car{"TN10AA3085"}

	attendantSimple.Park(car1)
	attendantMax.Park(car2)
	attendantEven.Park(car3)

	if !parkinglot1.slots[0].car.isEqual(car1) {
		t.Errorf("car1 should be parked in parkinglot1 by simple attendant")
	}

	if !parkinglot2.slots[0].car.isEqual(car3) {
		t.Errorf("car1 should be parked in parkinglot1 by simple attendant")
	}

	if !parkinglot3.slots[0].car.isEqual(car2) {
		t.Errorf("car1 should be parked in parkinglot1 by simple attendant")
	}

}

func TestAttendantMaxCapacityChoosesOtherLotWhenLargestIsFull(t *testing.T) {
	parkinglot1, _ := NewParkingLot(1)
	parkinglot2, _ := NewParkingLot(2)
	attendant, _ := NewAttendant(ParkInMaxCapacityLot, parkinglot1, parkinglot2)

	car1 := &Car{"KK10AA2345"}
	car2 := &Car{"TN10AA3085"}
	car3 := &Car{"TN10AB2555"}

	err := attendant.Park(car1)
	if err != nil {
		t.Fatalf("car1 should be parked: %v", err)
	}

	attendant.Park(car2)

	if parkinglot2.CarsParkedCount() != 2 {
		t.Fatalf("car2 should be parked in parkinglot2")
	}

	attendant.Park(car3)

	if !parkinglot1.slots[0].car.isEqual(car3) {
		t.Errorf("car3 has to be parked in parkinglot1")
	}
}

func TestAttendantMaxCapacityChoosesFirstLotWhenCapacitiesAreEqual(t *testing.T) {
	parkinglot1, _ := NewParkingLot(3)
	parkinglot2, _ := NewParkingLot(3)
	attendant, _ := NewAttendant(ParkInMaxCapacityLot, parkinglot1, parkinglot2)

	car1 := &Car{"KK10AA2345"}
	car2 := &Car{"TN10AA3085"}

	attendant.Park(car1)
	attendant.Park(car2)

	if !parkinglot1.slots[1].car.isEqual(car2) {
		t.Errorf("car2 should also be parked in lot1 before lot2 is used")
	}
}
