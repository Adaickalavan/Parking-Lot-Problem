package main

import (
	"minheap"
	"reflect"
	"testing"
)

//variables act as a struct of all parameters used in testing
type variables struct {
	car0         *Car
	car1         *Car
	car2         *Car
	map0         map[int]*Car
	map1         map[int]*Car
	map2         map[int]*Car
	mapAll       map[int]*Car
	item1        *minheap.Item
	item2        *minheap.Item
	emptySlot0   minheap.PriorityQueue
	emptySlot1   minheap.PriorityQueue
	emptySlot2   minheap.PriorityQueue
	emptySlotAll minheap.PriorityQueue
}

//values() acts a storage of default values and return a 'variables' struct containing default values
func values() variables {
	defaultValues := variables{
		car0:       &Car{registration: "KA-01-HH-2701", colour: "Blue"},
		car1:       &Car{slot: 1, registration: "KA-01-HH-1234", colour: "White"},
		car2:       &Car{slot: 2, registration: "KA-01-HH-7777", colour: "Red"},
		map0:       make(map[int]*Car),
		item1:      &minheap.Item{Value: 1},
		item2:      &minheap.Item{Value: 2},
		emptySlot0: minheap.PriorityQueue{},
	}
	defaultValues.map1 = map[int]*Car{1: defaultValues.car1}
	defaultValues.map2 = map[int]*Car{2: defaultValues.car2}
	defaultValues.mapAll = map[int]*Car{1: defaultValues.car1, 2: defaultValues.car2}
	defaultValues.emptySlot1 = minheap.PriorityQueue{defaultValues.item1}
	defaultValues.emptySlot2 = minheap.PriorityQueue{defaultValues.item2}
	defaultValues.emptySlotAll = minheap.PriorityQueue{defaultValues.item1, defaultValues.item2}

	return defaultValues
}

//Compare two 'Carpark' structs
func compareCarpark(t *testing.T, carpark *Carpark, wantCarpark *Carpark) {
	if !reflect.DeepEqual(carpark.Map, wantCarpark.Map) ||
		!reflect.DeepEqual(carpark.emptySlot, wantCarpark.emptySlot) ||
		carpark.highestSlot != wantCarpark.highestSlot ||
		carpark.MaxSlot != wantCarpark.MaxSlot {
		t.Errorf("gotCarpark = %v, wantCarpark = %v", carpark, wantCarpark)
	}
}

func TestCarpark_init(t *testing.T) {
	type args struct {
		maxSlot int
	}
	tests := []struct {
		name        string
		carpark     *Carpark
		args        args
		wantErr     bool
		wantCarpark *Carpark
	}{
		{name: "Carpark not initialized",
			carpark:     &Carpark{},
			args:        args{maxSlot: 12},
			wantErr:     false,
			wantCarpark: &Carpark{Map: values().map0, emptySlot: values().emptySlot0, highestSlot: 0, MaxSlot: 12},
		},
		{name: "Carpark already initialized",
			carpark:     &Carpark{Map: values().map0, emptySlot: values().emptySlot0, highestSlot: 8, MaxSlot: 10},
			args:        args{maxSlot: 12},
			wantErr:     true,
			wantCarpark: &Carpark{Map: values().map0, emptySlot: values().emptySlot0, highestSlot: 8, MaxSlot: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.carpark.init(tt.args.maxSlot)
			if (err != nil) != tt.wantErr {
				t.Errorf("Carpark.init() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}
			compareCarpark(t, tt.carpark, tt.wantCarpark)
		})
	}
}

func TestCarpark_insertCar(t *testing.T) {
	type args struct {
		car *Car
	}
	tests := []struct {
		name        string
		carpark     *Carpark
		args        args
		want        int
		wantErr     bool
		wantCarpark *Carpark
	}{
		{name: "Carpark not initialized",
			carpark:     &Carpark{},
			args:        args{car: values().car1},
			want:        0,
			wantErr:     true,
			wantCarpark: &Carpark{},
		},
		{name: "Insert car into new slot",
			carpark:     &Carpark{Map: values().map1, emptySlot: values().emptySlot0, highestSlot: 1, MaxSlot: 10},
			args:        args{car: values().car2},
			want:        2,
			wantErr:     false,
			wantCarpark: &Carpark{Map: values().mapAll, emptySlot: values().emptySlot0, highestSlot: 2, MaxSlot: 10},
		},
		{name: "Insert car into a previously occupied but now free slot",
			carpark:     &Carpark{Map: values().map2, emptySlot: values().emptySlot1, highestSlot: 2, MaxSlot: 10},
			args:        args{car: values().car1},
			want:        1,
			wantErr:     false,
			wantCarpark: &Carpark{Map: values().mapAll, emptySlot: values().emptySlot0, highestSlot: 2, MaxSlot: 10},
		},
		{name: "Insert car beyond MaxSlot",
			carpark:     &Carpark{Map: values().mapAll, emptySlot: values().emptySlot0, highestSlot: 2, MaxSlot: 2},
			args:        args{car: values().car0},
			want:        0,
			wantErr:     true,
			wantCarpark: &Carpark{Map: values().mapAll, emptySlot: values().emptySlot0, highestSlot: 2, MaxSlot: 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.carpark.insertCar(tt.args.car)
			if (err != nil) != tt.wantErr {
				t.Errorf("Carpark.insertCar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Carpark.insertCar() = %v, want %v", got, tt.want)
			}
			compareCarpark(t, tt.carpark, tt.wantCarpark)
		})
	}
}

func TestCarpark_removeCar(t *testing.T) {
	type args struct {
		slotNo int
	}
	tests := []struct {
		name        string
		carpark     *Carpark
		args        args
		wantErr     bool
		wantCarpark *Carpark
	}{
		{name: "Carpark not initialized",
			carpark:     &Carpark{},
			args:        args{slotNo: 1},
			wantErr:     true,
			wantCarpark: &Carpark{},
		},
		{name: "Remove car",
			carpark:     &Carpark{Map: values().mapAll, emptySlot: values().emptySlot0, highestSlot: 2, MaxSlot: 10},
			args:        args{slotNo: 1},
			wantErr:     false,
			wantCarpark: &Carpark{Map: values().map2, emptySlot: values().emptySlot1, highestSlot: 2, MaxSlot: 10},
		},
		{name: "Remove non-existent car",
			carpark:     &Carpark{Map: values().map1, emptySlot: values().emptySlot2, highestSlot: 2, MaxSlot: 10},
			args:        args{slotNo: 2},
			wantErr:     true,
			wantCarpark: &Carpark{Map: values().map1, emptySlot: values().emptySlot2, highestSlot: 2, MaxSlot: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.carpark.removeCar(tt.args.slotNo); (err != nil) != tt.wantErr {
				t.Errorf("Carpark.removeCar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			compareCarpark(t, tt.carpark, tt.wantCarpark)
		})
	}
}

func TestCarpark_getCarsWithColour(t *testing.T) {
	type args struct {
		colour string
	}
	tests := []struct {
		name    string
		carpark *Carpark
		args    args
		want    []int
		want1   []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.carpark.getCarsWithColour(tt.args.colour)
			if (err != nil) != tt.wantErr {
				t.Errorf("Carpark.getCarsWithColour() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Carpark.getCarsWithColour() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Carpark.getCarsWithColour() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
