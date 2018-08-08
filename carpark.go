package main

import (
	"container/heap"
	"errors"
	"fmt"
	"minheap"
	"os"
	"text/tabwriter"
)

//Carpark represents the carpark map, empty slots, and maximum number of slots filled
type Carpark struct {
	Map         map[int]*Car          //Properties of each parked in the carpark
	EmptySlot   minheap.PriorityQueue //Heap containing sorted empty slots in ascending order
	HighestSlot int                   //Highest number of slots filled throughout carpark operation
	MaxSlot     int                   //Maximum number of slots available
}

//Initialize carpark parameters
func (carpark *Carpark) init(maxSlot int) error {
	if err := carpark.initStatus(); err == nil {
		return errors.New("Carpark already initialized")
	}
	carpark.Map = make(map[int]*Car)            //Setup a map of the carpark
	carpark.EmptySlot = minheap.PriorityQueue{} //Setup an empty heap of empty parking slots
	heap.Init(&carpark.EmptySlot)               //Initialize the heap of empty parking slots
	carpark.MaxSlot = maxSlot                   //Set the maximum number of slots
	return nil
}

//Park a car in carpark
func (carpark *Carpark) insertCar(car *Car) (int, error) {
	if err := carpark.initStatus(); err != nil {
		return 0, err
	}
	var slotNo int
	if carpark.EmptySlot.Len() == 0 { //Get next available slot if no empty slots avaiable
		slotNo = carpark.HighestSlot + 1
		carpark.HighestSlot = slotNo
	} else { //Get nearest empty slot if available
		item := heap.Pop(&carpark.EmptySlot)
		slotNo = item.(*minheap.Item).Value
	}
	if slotNo > carpark.MaxSlot {
		return 0, errors.New("Sorry, parking lot is full")
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
		heap.Push(&carpark.EmptySlot, &minheap.Item{Value: slotNo})
		return nil
	}
	return errors.New("Car non-existent in carpark")
}

//Given a car colour, retrieve the car slot and registration numbers
func (carpark *Carpark) getCarsWithColour(colour string) ([]int, []string, error) {
	var slots []int
	var registrations []string
	if err := carpark.initStatus(); err != nil {
		return nil, nil, err
	}
	for _, v := range carpark.Map {
		if v.colour == colour {
			slots = append(slots, v.slot)
			registrations = append(registrations, v.registration)
		}
	}
	return slots, registrations, nil
}

//Given a car registration number, retrieve the car slot number
func (carpark *Carpark) getCarWithRegistrationNo(registration string) (int, error) {
	if err := carpark.initStatus(); err != nil {
		return 0, err
	}
	for _, v := range carpark.Map {
		if v.registration == registration {
			return v.slot, nil
		}
	}
	return 0, errors.New("Not found")
}

//Print details of all the cars parked in the carpark
func (carpark *Carpark) getStatus() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "Slot No.\tRegistration No   \tColour\t")
	for i := 1; i <= carpark.HighestSlot; i++ {
		car, ok := carpark.Map[i]
		if ok {
			s := fmt.Sprintf("%v\t%s\t%s", car.slot, car.registration, car.colour)
			fmt.Fprintln(w, s)
		}
	}
	w.Flush()
}

//Check whether the carpark has been initialized
func (carpark *Carpark) initStatus() error {
	if carpark.Map == nil {
		return errors.New("Carpark not initialized")
	}
	return nil
}
