package main

import (
	"minheap"
	"testing"
)

func TestCarpark_init(t *testing.T) {
	type args struct {
		maxSlot int
	}
	tests := []struct {
		name    string
		carpark *Carpark
		args    args
		wantErr bool
	}{
		{name: "Carpark not initialized",
			carpark: &Carpark{},
			args:    args{maxSlot: 12},
			wantErr: false,
		},
		{name: "Carpark already initialized",
			carpark: &Carpark{Map: make(map[int]*Car), EmptySlot: minheap.PriorityQueue{}, HighestSlot: 10},
			args:    args{maxSlot: 12},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.carpark.init(tt.args.maxSlot); (err != nil) != tt.wantErr {
				t.Errorf("Carpark.init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
