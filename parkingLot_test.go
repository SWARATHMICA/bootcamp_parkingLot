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
	p, _ := NewParkingLot(1)
	car1 := &Car{
		numberPlate: "KK-09-AK-1234",
	}
	car2 := &Car{
		numberPlate: "KK-09-AK-2341",
	}
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
	p, _ := NewParkingLot(1)
	o := Owner{Name: "abc", ParkingLot: p}
	if o.Name != "abc" {
		t.Error("Owner struct not created")
	}
}

func TestNotifyOwner(t *testing.T) {
	parkinglotwith1Capacity, _ := NewParkingLot(1)
	car1 := &Car{
		numberPlate: "KK-09-AK-1234",
	}
	parkinglotwith1Capacity.Park(*car1)

	owner := Owner{Name: "abc", ParkingLot: parkinglotwith1Capacity}
	message := owner.notify()
	expectedMessage := "ParkingLot is full!!!!"
	if message != expectedMessage {
		t.Errorf("Parking lot is full")
	}

}
func TestNotifyOwnerWhenParkingLotAvailable(t *testing.T) {
	parkinglotwith2Capacity, _ := NewParkingLot(2)
	car1 := &Car{
		numberPlate: "KK-09-AK-1234",
	}
	parkinglotwith2Capacity.Park(*car1)

	owner := Owner{Name: "abc", ParkingLot: parkinglotwith2Capacity}
	message := owner.notify()
	expectedMessage := "Parking lot available"
	if message != expectedMessage {
		t.Errorf("Parking lot should be available")
	}

}
