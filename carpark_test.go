package main

import (
	"minheap"
	"reflect"
	"testing"
)

var car0 = &Car{registration: "KA-01-HH-2701", colour: "Blue"}
var car1 = &Car{slot: 1, registration: "KA-01-HH-1234", colour: "White"}
var car2 = &Car{slot: 2, registration: "KA-01-HH-7777", colour: "Red"}
var map0 = make(map[int]*Car)
var map1 = map[int]*Car{1: car1}
var map2 = map[int]*Car{2: car2}
var mapAll = map[int]*Car{1: car1, 2: car2}
var item1 = &minheap.Item{Value: 1}
var item2 = &minheap.Item{Value: 2}
var emptySlot0 = minheap.PriorityQueue{}
var emptySlot1 = minheap.PriorityQueue{item1}
var emptySlot2 = minheap.PriorityQueue{item2}
var emptySlotAll = minheap.PriorityQueue{item1, item2}

func compareCarpark(t *testing.T, carpark *Carpark, wantCarpark *Carpark) {
	if !reflect.DeepEqual(carpark.Map, wantCarpark.Map) ||
		!reflect.DeepEqual(carpark.EmptySlot, wantCarpark.EmptySlot) ||
		carpark.HighestSlot != wantCarpark.HighestSlot ||
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
			wantCarpark: &Carpark{Map: map0, EmptySlot: emptySlot0, HighestSlot: 0, MaxSlot: 12},
		},
		{name: "Carpark already initialized",
			carpark:     &Carpark{Map: map0, EmptySlot: emptySlot0, HighestSlot: 8, MaxSlot: 10},
			args:        args{maxSlot: 12},
			wantErr:     true,
			wantCarpark: &Carpark{Map: map0, EmptySlot: emptySlot0, HighestSlot: 8, MaxSlot: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.carpark.init(tt.args.maxSlot)
			if (err != nil) != tt.wantErr {
				t.Errorf("Carpark.init() error = %v, wantErr = %v", err, tt.wantErr)
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
			args:        args{car: car1},
			want:        0,
			wantErr:     true,
			wantCarpark: &Carpark{},
		},
		{name: "Insert car into new slot",
			carpark:     &Carpark{Map: map1, EmptySlot: emptySlot0, HighestSlot: 1, MaxSlot: 10},
			args:        args{car: car2},
			want:        2,
			wantErr:     false,
			wantCarpark: &Carpark{Map: mapAll, EmptySlot: emptySlot0, HighestSlot: 2, MaxSlot: 10},
		},
		{name: "Insert car into a previously occupied but now free slot",
			carpark:     &Carpark{Map: map2, EmptySlot: emptySlot1, HighestSlot: 2, MaxSlot: 10},
			args:        args{car: car1},
			want:        1,
			wantErr:     false,
			wantCarpark: &Carpark{Map: mapAll, EmptySlot: emptySlot0, HighestSlot: 2, MaxSlot: 10},
		},
		{name: "Insert car above MaxSlot",
			carpark:     &Carpark{Map: mapAll, EmptySlot: emptySlot0, HighestSlot: 2, MaxSlot: 2},
			args:        args{car: car0},
			want:        0,
			wantErr:     true,
			wantCarpark: &Carpark{Map: mapAll, EmptySlot: emptySlot0, HighestSlot: 2, MaxSlot: 2},
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
