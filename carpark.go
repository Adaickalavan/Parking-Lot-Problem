package main

import (
	"container/heap"
	"errors"
	"fmt"
	"minheap"
)

// Carpark represents the carpark map, empty slots, and maximum number of slots filled
type Carpark struct {
	Map         map[int]*Car          //Properties of each parked in the carpark
	EmptySlot   minheap.PriorityQueue //Heap containing sorted empty slots in ascending order
	HighestSlot int                   //Highest number of slots filled throughout carpark operation
	MaxSlot     int                   //Maximum number of slots available
}

func (carpark *Carpark) init(maxSlot int) {
	carpark.Map = make(map[int]*Car)            //Setup a map of the carpark
	carpark.EmptySlot = minheap.PriorityQueue{} //Setup an empty heap of empty parking slots
	heap.Init(&carpark.EmptySlot)               //Initialize the heap of empty parking slots
	carpark.MaxSlot = maxSlot                   //Set the maximum number of slots
}

func (carpark *Carpark) insertCar(car *Car) (int, error) {
	//Get next available slot if no empty slots avaiable
	var slotNo int
	if carpark.EmptySlot.Len() == 0 {
		slotNo = carpark.HighestSlot + 1
		carpark.HighestSlot = slotNo
	} else { //Get nearest empty slot if available
		item := heap.Pop(&carpark.EmptySlot)
		slotNo = item.(*minheap.Item).Value
	}
	if slotNo > carpark.MaxSlot {
		return 0, fmt.Errorf("Sorry, parking lot is full")
	}
	//Park the car at the slotNo
	car.slot = slotNo
	carpark.Map[slotNo] = car
	return slotNo, nil
}

func (carpark *Carpark) removeCar(slotNo int) error {
	if _, ok := carpark.Map[slotNo]; ok {
		//Remove car from carpark Map
		delete(carpark.Map, slotNo)
		//Add empty slot to the heap
		heap.Push(&carpark.EmptySlot, &minheap.Item{Value: slotNo})
		return nil
	}
	return errors.New("Car non-existent in carpark")
}

func (carpark *Carpark) getCarsWithColour(colour string) ([]int, []string) {
	var slots []int
	var registrations []string
	for _, v := range carpark.Map {
		if v.colour == colour {
			slots = append(slots, v.slot)
			registrations = append(registrations, v.registration)
		}
	}
	return slots, registrations
}

func (carpark *Carpark) getCarWithRegistrationNo(registration string) (int, error) {
	for _, v := range carpark.Map {
		if v.registration == registration {
			return v.slot, nil
		}
	}
	return 0, fmt.Errorf("Not found")
}
