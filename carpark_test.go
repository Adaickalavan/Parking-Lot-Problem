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
		carpark.maxSlot != wantCarpark.maxSlot {
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
			wantCarpark: &Carpark{Map: values().map0, emptySlot: values().emptySlot0, highestSlot: 0, maxSlot: 12},
		},
		{name: "Carpark already initialized",
			carpark:     &Carpark{Map: values().map0, emptySlot: values().emptySlot0, highestSlot: 8, maxSlot: 10},
			args:        args{maxSlot: 12},
			wantErr:     true,
			wantCarpark: &Carpark{Map: values().map0, emptySlot: values().emptySlot0, highestSlot: 8, maxSlot: 10},
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
			carpark:     &Carpark{Map: values().map1, emptySlot: values().emptySlot0, highestSlot: 1, maxSlot: 10},
			args:        args{car: values().car2},
			want:        2,
			wantErr:     false,
			wantCarpark: &Carpark{Map: values().mapAll, emptySlot: values().emptySlot0, highestSlot: 2, maxSlot: 10},
		},
		{name: "Insert car into a previously occupied but now free slot",
			carpark:     &Carpark{Map: values().map2, emptySlot: values().emptySlot1, highestSlot: 2, maxSlot: 10},
			args:        args{car: values().car1},
			want:        1,
			wantErr:     false,
			wantCarpark: &Carpark{Map: values().mapAll, emptySlot: values().emptySlot0, highestSlot: 2, maxSlot: 10},
		},
		{name: "Insert car beyond maxSlot",
			carpark:     &Carpark{Map: values().mapAll, emptySlot: values().emptySlot0, highestSlot: 2, maxSlot: 2},
			args:        args{car: values().car0},
			want:        0,
			wantErr:     true,
			wantCarpark: &Carpark{Map: values().mapAll, emptySlot: values().emptySlot0, highestSlot: 2, maxSlot: 2},
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
			carpark:     &Carpark{Map: values().mapAll, emptySlot: values().emptySlot0, highestSlot: 2, maxSlot: 10},
			args:        args{slotNo: 1},
			wantErr:     false,
			wantCarpark: &Carpark{Map: values().map2, emptySlot: values().emptySlot1, highestSlot: 2, maxSlot: 10},
		},
		{name: "Remove non-existent car",
			carpark:     &Carpark{Map: values().map1, emptySlot: values().emptySlot2, highestSlot: 2, maxSlot: 10},
			args:        args{slotNo: 2},
			wantErr:     true,
			wantCarpark: &Carpark{Map: values().map1, emptySlot: values().emptySlot2, highestSlot: 2, maxSlot: 10},
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
		{name: "Carpark with car of requested colour",
			carpark: &Carpark{Map: values().map1, emptySlot: values().emptySlot0, highestSlot: 1, maxSlot: 10},
			args:    args{colour: "White"},
			want:    []int{1},
			want1:   []string{"KA-01-HH-1234"},
			wantErr: false,
		},
		{name: "Carpark without car of requested colour",
			carpark: &Carpark{Map: values().map2, emptySlot: values().emptySlot1, highestSlot: 2, maxSlot: 10},
			args:    args{colour: "White"},
			want:    nil,
			want1:   nil,
			wantErr: true,
		},
		{name: "Empty carpark",
			carpark: &Carpark{Map: values().map0, emptySlot: values().emptySlot0, highestSlot: 0, maxSlot: 10},
			args:    args{colour: "White"},
			want:    nil,
			want1:   nil,
			wantErr: true,
		},
		{name: "Uninitialized carpark",
			carpark: &Carpark{},
			args:    args{colour: "White"},
			want:    nil,
			want1:   nil,
			wantErr: true,
		},
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

func TestCarpark_getCarWithRegistrationNo(t *testing.T) {
	type args struct {
		registration string
	}
	tests := []struct {
		name    string
		carpark *Carpark
		args    args
		want    int
		wantErr bool
	}{
		{name: "Carpark with car of requested colour",
			carpark: &Carpark{Map: values().map1, emptySlot: values().emptySlot0, highestSlot: 1, maxSlot: 10},
			args:    args{registration: "KA-01-HH-1234"},
			want:    1,
			wantErr: false,
		},
		{name: "Carpark without car of requested colour",
			carpark: &Carpark{Map: values().map2, emptySlot: values().emptySlot1, highestSlot: 2, maxSlot: 10},
			args:    args{registration: "KA-01-HH-1234"},
			want:    0,
			wantErr: true,
		},
		{name: "Empty carpark",
			carpark: &Carpark{Map: values().map0, emptySlot: values().emptySlot0, highestSlot: 0, maxSlot: 10},
			args:    args{registration: "KA-01-HH-1234"},
			want:    0,
			wantErr: true,
		},
		{name: "Uninitialized carpark",
			carpark: &Carpark{},
			args:    args{registration: "KA-01-HH-1234"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.carpark.getCarWithRegistrationNo(tt.args.registration)
			if (err != nil) != tt.wantErr {
				t.Errorf("Carpark.getCarWithRegistrationNo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Carpark.getCarWithRegistrationNo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarpark_getStatus(t *testing.T) {
	tests := []struct {
		name    string
		carpark *Carpark
		want    []*Car
	}{
		{name: "Carpark not initialized",
			carpark: &Carpark{},
			want:    nil,
		},
		{name: "Empty carpark",
			carpark: &Carpark{Map: values().map0, emptySlot: values().emptySlot0, highestSlot: 0, maxSlot: 10},
			want:    nil,
		},
		{name: "Carpark with cars",
			carpark: &Carpark{Map: values().mapAll, emptySlot: values().emptySlot0, highestSlot: 2, maxSlot: 10},
			want:    []*Car{values().car1, values().car2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.carpark.getStatus(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Carpark.getStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
