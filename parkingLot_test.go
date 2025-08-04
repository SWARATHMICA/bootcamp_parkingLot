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

func TestOwnerStruct(t *testing.T) {

	o := Owner{Name: "abc"}
	if o.Name != "abc" {
		t.Error("Owner struct not created")
	}
}

type dummyObserver struct{}

func (d *dummyObserver) notifyFull()      {}
func (d *dummyObserver) notifyAvailable() {}

func TestSetObserver(t *testing.T) {
	lot, _ := NewParkingLot(1)
	observer := &dummyObserver{}

	lot.setObserver(observer)

	if lot.observer != observer {
		t.Error("Observer was not set correctly in the parking lot")
	}
}

type checkObserver struct {
	fullCalled      bool
	availableCalled bool
}

func (s *checkObserver) notifyFull() {
	s.fullCalled = true
}

func (s *checkObserver) notifyAvailable() {
	s.availableCalled = true
}

func TestObserverMethodsCalled(t *testing.T) {
	s := &checkObserver{}
	s.notifyFull()
	s.notifyAvailable()

	if !s.fullCalled {
		t.Error("Expected notifyFull to be called")
	}
	if !s.availableCalled {
		t.Error("Expected notifyAvailable to be called")
	}
}

func TestNotifyFullWhenLotBecomesFull(t *testing.T) {
	lot, _ := NewParkingLot(1)
	s := &checkObserver{}
	lot.setObserver(s)

	car := Car{numberPlate: "JK-45-WJ-1234"}

	lot.Park(car)

	if !s.fullCalled {
		t.Error("Expected notifyFull to be called when lot becomes full")
	}
}

func TestNotifyAvailableWhenCarIsUnparked(t *testing.T) {
	lot, _ := NewParkingLot(1)
	s := &checkObserver{}
	lot.setObserver(s)

	car := Car{numberPlate: "MH12ZZ9999"}
	lot.Park(car)

	s.fullCalled = false // Clear previous call
	s.availableCalled = false

	lot.Unpark(car)

	if !s.availableCalled {
		t.Error("Expected notifyAvailable to be called when car is unparked")
	}
}

func TestNotifyFullOnReparkingAfterUnpark(t *testing.T) {
	lot, _ := NewParkingLot(1)
	s := &checkObserver{}
	lot.setObserver(s)

	car := Car{numberPlate: "MH12ZZ9999"}
	lot.Park(car)
	lot.Unpark(car)

	s.fullCalled = false

	lot.Park(car)

	if !s.fullCalled {
		t.Error("Expected notifyFull to be called again after re-parking")
	}
}
