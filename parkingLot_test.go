package parkinglot

import (
	"testing"
)

func TestCreateLot(t *testing.T) {
	firstLot := &slot{
		id:          1,
		numberPlate: "KJ-09-AK-123",
		occupied:    false,
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
	parked, _ := parkingLot.Park(car)
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
	p.Park((*car1))
	result, _ := p.Park(*car2)

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
	p.Park((*car1))
	_, err := p.Park(*car2)

	if err == nil {
		t.Errorf("ParkingLot is full")
	}

}

func TestUnparkCar(t *testing.T) {
	p, _ := NewParkingLot(1)
	car := &Car{
		numberPlate: "KK-09-AK-2341",
	}
	p.Park(*car)

	result, _ := p.Unpark(*car)
	if !result {
		t.Errorf("Car not unparked")

	}
}

func TestUnparkCarNotFound(t *testing.T) {
	p, _ := NewParkingLot(1)
	car := &Car{
		numberPlate: "KK-09-AK-2341",
	}
	p.Park(*car)

	_, err := p.Unpark(*car)
	if err != nil {
		t.Errorf("Car not found")
	}
}

func TestCarIsParked(t *testing.T) {
	p, _ := NewParkingLot(1)
	car := &Car{
		numberPlate: "KK-09-AK-2341",
	}
	p.Park(*car)

	result := p.IsParked(*car)

	if !result {
		t.Errorf("Car is not parked in the ParkingLot")
	}

}
func TestCheckIfCarAlreadyParked(t *testing.T) {
	p, _ := NewParkingLot(2)
	car1 := &Car{
		numberPlate: "KK-09-AK-1234",
	}
	car2 := &Car{
		numberPlate: "KK-09-AK-1234",
	}
	p.Park((*car1))
	_, err := p.Park(*car2)
	if err == nil {
		t.Errorf("Car is already parked")
	}

}

type mockParkingFullReceiver struct {
	status ParkingStatus
}

func (s *mockParkingFullReceiver) receiveFull() {
	s.status = PARKING_FULL
}

var car = Car{numberPlate: "xts"}

// TODO fix name
func TestCheckNotifyCalledOnReceiver(t *testing.T) {
	s := &mockParkingFullReceiver{}
	p, _ := NewParkingLot(1)
	p.setReceiver(s)
	p.Park(car)
	expectedStatus := PARKING_FULL
	if s.status != expectedStatus {
		t.Errorf("Notify not called")
	}
}
func TestMultipleReiversShouldBeNotifiedWhenParkingFull(t *testing.T) {
	p, _ := NewParkingLot(1)
	s := &mockParkingFullReceiver{}
	another := &mockParkingFullReceiver{}

	p.addParkingFullReceiver(s)
	p.addParkingFullReceiver(another)
	p.Park(car)

	expectedStatus := PARKING_FULL

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
	p.Park(car)
	p.Unpark(car)

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

	p.Park(car)
	if parkingFullReceiver.status != PARKING_FULL {
		t.Errorf("Only one person should be notified for parking availability")
	}
	if parkingAvailableReceiver.receiveCalled == true {
		t.Errorf("Parking Available Receiver should not be notified when parking full")
	}

	p.Unpark(car)
	if parkingFullReceiver.status != PARKING_FULL {
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
	p.Park(car)

	if owner.notifiedFull != true {
		t.Errorf("owner should be notified when the parking lot gets full")
	}

	p.Unpark(car)

	if owner.notifiedAvailable != true {
		t.Errorf("owner should be notified when the parking lot is available")
	}

}

func TestParkCarWithAttendent(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)
	attendent := Attendent{}
	attendent.addParkingLotInAttendent(parkingLot)

	attendent.Parkinglots[0].Park(car)

	if parkingLot.isFull == false {
		t.Errorf("parking lot should be full when parked through attendent")
	}
}

func TestUnParkCarWithAttendent(t *testing.T) {
	parkingLot, _ := NewParkingLot(1)
	attendent := Attendent{}
	attendent.addParkingLotInAttendent(parkingLot)

	attendent.Parkinglots[0].Park(car)
	attendent.Parkinglots[0].Unpark(car)

	if parkingLot.isFull {
		t.Errorf("parking lot should  not be full when unparked through attendent")
	}
}
