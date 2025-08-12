package parkinglot

import (
	"testing"
)

func TestCreateLot(t *testing.T) {
	firstLot := &slot{
		id:       1,
		car:      &Car{"KJ-09-AK-123"},
		occupied: false,
	}

	if firstLot.id == 0 {
		t.Error("Struct Lot cannot be created with Id 0")
	}
}

func TestCreateParkingLot(t *testing.T) {
	_, err := NewParkingLot(10)
	if err != nil {
		t.Error("Parking lot is not created")
	}
}

func TestCreateEmptyLotsInParkingLot(t *testing.T) {
	_, err := NewParkingLot(10)
	if err != nil {
		t.Error("Empty lots are not created in ParkingLot")
	}
}

func TestCannotCreateParkingLotWithCapacityLessThanOne(t *testing.T) {
	_, err := NewParkingLot(-1)
	if err == nil {
		t.Error("Parking lot cannot be created with capacity less than 1")
	}
}

func TestParkCar(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)
	car := Car{
		numberPlate: "KJ-09-AK-123",
	}
	parked, _ := parkingLot.park(&car)
	if !parked {
		t.Errorf("Vehicle not parked")
	}
}

func TestCreateParkingLotStruct(t *testing.T) {
	_, err := NewParkingLot((10))
	if err != nil {
		t.Errorf("Parking Lot is not created")
	}

}

func TestCreateCarStruct(t *testing.T) {
	car := &Car{

		numberPlate: "KJ-09-AK-123",
	}

	if car.numberPlate == "" {
		t.Error("Struct Car cannot be created without NumberPlate")
	}
}

func TestParkingLotCreationWithCapacity(t *testing.T) {
	p, _ := NewParkingLot(1)
	if p.capacity != 1 {
		t.Errorf("ParkingLot has not been created with capacity 1")
	}
}

func TestCheckIfLotIsEmptyBeforeParking(t *testing.T) {
	p, _ := NewParkingLot(1)
	car1 := &Car{
		numberPlate: "KK-09-AK-1234",
	}
	car2 := &Car{
		numberPlate: "KK-09-AK-2341",
	}
	p.park(car1)
	result, _ := p.park(car2)

	if result {
		t.Errorf("Cannot park car since no slots are empty")
	}

}

func TestCheckIfParkingLotIsFull(t *testing.T) {
	car1 := &Car{
		numberPlate: "KK-09-AK-1234",
	}
	car2 := &Car{
		numberPlate: "KK-09-AK-2341",
	}
	p, _ := NewParkingLot(1)
	p.park(car1)
	_, err := p.park(car2)

	if err == nil {
		t.Errorf("ParkingLot is full")
	}

}

func TestUnparkCar(t *testing.T) {
	p, _ := NewParkingLot(1)
	car := &Car{
		numberPlate: "KK-09-AK-2341",
	}
	p.park(car)

	result, _ := p.unPark(car)
	if !result {
		t.Errorf("Car not unparked")

	}
}

func TestUnparkCarNotFound(t *testing.T) {
	p, _ := NewParkingLot(1)
	car := &Car{
		numberPlate: "KK-09-AK-2341",
	}
	p.park(car)

	_, err := p.unPark(car)
	if err != nil {
		t.Errorf("Car not found")
	}
}

func TestCarIsParked(t *testing.T) {
	p, _ := NewParkingLot(1)
	car := &Car{
		numberPlate: "KK-09-AK-2341",
	}
	p.park(car)

	result := p.isParked(*car)

	if !result {
		t.Errorf("Car is not parked in the ParkingLot")
	}

}
func TestCheckIfCarAlreadyParked(t *testing.T) {
	p, _ := NewParkingLot(2)
	const numberPlate = "KK-09-AK-1234"
	car1 := &Car{
		numberPlate: numberPlate,
	}
	car2 := &Car{
		numberPlate: numberPlate,
	}
	p.park(car1)
	_, err := p.park(car2)
	if err == nil {
		t.Errorf("Car is already parked")
	}

}

type mockParkingFullReceiver struct {
	status ParkingStatus
}

func (s *mockParkingFullReceiver) receiveFull() {
	s.status = ParkingFull
}

var car = Car{numberPlate: "AB-12-CD-3456"}

func TestMultipleReceiversShouldBeNotifiedWhenParkingFull(t *testing.T) {
	p, _ := NewParkingLot(1)
	s := &mockParkingFullReceiver{}
	another := &mockParkingFullReceiver{}

	p.addParkingFullReceiver(s)
	p.addParkingFullReceiver(another)
	p.park(&car)

	expectedStatus := ParkingFull

	if s.status != expectedStatus {
		t.Errorf("The status should be changed to parking_full")
	}
	if another.status != expectedStatus {
		t.Errorf("Another person should have status parking_full")
	}

}

type mockParkingAvailableReceiver struct {
	receiveCalled bool
}

func (m *mockParkingAvailableReceiver) receiveAvailable() {
	m.receiveCalled = !m.receiveCalled
}

func TestSingleRecieverNotifiedParkingAvailable(t *testing.T) {
	p, _ := NewParkingLot(1)
	parkingAvailableReceiver := mockParkingAvailableReceiver{}
	p.setParkingAvailableReceiver(&parkingAvailableReceiver)
	p.park(&car)
	p.unPark(&car)

	if !parkingAvailableReceiver.receiveCalled {
		t.Errorf("Receive Function not called")
	}
}

func TestCannotNofifyMultiplePeopleWhenParkingAvailable(t *testing.T) {
	p, _ := NewParkingLot(1)
	parkingFullReceiver := mockParkingFullReceiver{}
	parkingAvailableReceiver := mockParkingAvailableReceiver{}

	p.addParkingFullReceiver(&parkingFullReceiver)

	p.setParkingAvailableReceiver(&parkingAvailableReceiver)

	p.park(&car)
	if parkingFullReceiver.status != ParkingFull {
		t.Errorf("Only one person should be notified for parking availability")
	}
	if parkingAvailableReceiver.receiveCalled == true {
		t.Errorf("Parking Available Receiver should not be notified when parking full")
	}

	p.unPark(&car)
	if parkingFullReceiver.status != ParkingFull {
		t.Errorf("Only one person should be notified for parking availability")
	}

	if !parkingAvailableReceiver.receiveCalled {
		t.Errorf("This receiver has to be Notified when parking is available")
	}

}

type ReceiveBothNotification struct {
	notifiedFull      bool
	notifiedAvailable bool
}

func (o *ReceiveBothNotification) receiveFull() {
	o.notifiedFull = true
}

func (o *ReceiveBothNotification) receiveAvailable() {
	o.notifiedAvailable = true
}

func TestNotifyOwnerWhenFullAndWhenAvailable(t *testing.T) {
	p, _ := NewParkingLot(1)
	owner := ReceiveBothNotification{}
	p.setParkingAvailableReceiver(&owner)
	p.addParkingFullReceiver(&owner)
	p.park(&car)

	if !owner.notifiedFull {
		t.Errorf("owner should be notified when the parking lot gets full")
	}

	p.unPark(&car)

	if !owner.notifiedAvailable {
		t.Errorf("owner should be notified when the parking lot is available")
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

func TestAttendantShouldNotBeCreateWithNilValue(t *testing.T) {

	_, err := NewAttendant(nil)

	expectedError := "cannot create attedant with nil parkinglot"

	if err.Error() != expectedError {
		t.Errorf("attendant should not be creatd with nil parking lot")
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

func TestTwoCarsEqual(t *testing.T) {
	const carNumber = "MH12AA2345"
	car1 := Car{carNumber}
	car2 := Car{carNumber}

	result := car1.isEqual(car2)

	if !result {
		t.Errorf("car1 should be equal to car2")
	}

}

func TestTwoCarsNotEqual(t *testing.T) {
	car1 := Car{"MH12AA2345"}
	car2 := Car{"UP12HH4009"}

	result := car1.isEqual(car2)

	if result {
		t.Errorf("car1 should not be equal to car2")
	}

}

func TestParkAfterUnpark(t *testing.T) {
	parkinglot, _ := NewParkingLot(1)

	parkinglot.park(&car)
	parkinglot.unPark(&car)

	_, err := parkinglot.park(&car)

	if err != nil {
		t.Errorf("car should park after nupark")
	}
}
