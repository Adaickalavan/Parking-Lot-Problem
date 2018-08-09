package main

import (
	"container/heap"
	"errors"
	"minheap"
)

//Carpark represents the carpark map, empty slots, and maximum number of slots filled
type Carpark struct {
	Map         map[int]*Car          //Properties of each car parked in the carpark
	emptySlot   minheap.PriorityQueue //Heap containing sorted empty slots in ascending order
	highestSlot int                   //Highest number of slots filled throughout carpark operation
	maxSlot     int                   //Maximum number of slots available
}

//Initialize carpark parameters
func (carpark *Carpark) init(maxSlot int) error {
	if err := carpark.initStatus(); err == nil {
		return errors.New("Carpark already initialized")
	}
	carpark.Map = make(map[int]*Car)            //Setup a map of the carpark
	carpark.emptySlot = minheap.PriorityQueue{} //Setup an empty heap of empty parking slots
	heap.Init(&carpark.emptySlot)               //Initialize the heap of empty parking slots
	carpark.maxSlot = maxSlot                   //Set the maximum number of slots
	return nil
}

//Park a car in carpark
func (carpark *Carpark) insertCar(car *Car) (int, error) {
	if err := carpark.initStatus(); err != nil {
		return 0, err
	}
	var slotNo int
	//Check whether all slots are occupied
	if carpark.emptySlot.Len() == 0 {
		if carpark.highestSlot == carpark.maxSlot { //Check whether all slots occupied
			return 0, errors.New("Sorry, parking lot is full")
		}
		//Get next available slot
		slotNo = carpark.highestSlot + 1
		carpark.highestSlot = slotNo
	} else { //Get nearest empty slot which was previously occupied
		item := heap.Pop(&carpark.emptySlot)
		slotNo = item.(*minheap.Item).Value
	}
	//Park the car at the slotNo
	car.slot = slotNo
	carpark.Map[slotNo] = car
	return slotNo, nil
}

//Remove car from carpark
func (carpark *Carpark) removeCar(slotNo int) error {
	if err := carpark.initStatus(); err != nil {
		return err
	}
	if _, ok := carpark.Map[slotNo]; ok {
		//Remove car from carpark Map
		delete(carpark.Map, slotNo)
		//Add empty slot to the heap
		heap.Push(&carpark.emptySlot, &minheap.Item{Value: slotNo})
		return nil
	}
	return errors.New("Car non-existent in carpark")
}

//Given a car colour, retrieve the car slot and registration numbers
func (carpark *Carpark) getCarsWithColour(colour string) ([]int, []string, error) {
	var slots []int
	var registrations []string
	for _, v := range carpark.Map {
		if v.colour == colour {
			slots = append(slots, v.slot)
			registrations = append(registrations, v.registration)
		}
	}
	if slots == nil {
		return nil, nil, errors.New("Not found")
	}
	return slots, registrations, nil
}

//Given a car registration number, retrieve the car slot number
func (carpark *Carpark) getCarWithRegistrationNo(registration string) (int, error) {
	for _, v := range carpark.Map {
		if v.registration == registration {
			return v.slot, nil
		}
	}
	return 0, errors.New("Not found")
}

//Retrieve details of the cars parked in the carpark in order
func (carpark *Carpark) getStatus() []*Car {
	var cars []*Car
	for i := 1; i <= carpark.highestSlot; i++ {
		car, ok := carpark.Map[i]
		if ok {
			cars = append(cars, car)
		}
	}
	return cars
}

//Check whether the carpark has been initialized
func (carpark *Carpark) initStatus() error {
	if carpark.Map == nil {
		return errors.New("Carpark not initialized")
	}
	return nil
}
