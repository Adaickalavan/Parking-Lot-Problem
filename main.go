package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"util"
)

//Create a carpark
var carpark = &Carpark{
	Map:       make(map[int]*Car),
	EmptySlot: util.PriorityQueue{},
}

func main() {
	//Create reader input from console
	reader := bufio.NewReader(os.Stdin)

	//Identify operating system used
	var terminal string
	if runtime.GOOS == "windows" {
		terminal = "\r\n"
	} else {
		terminal = "\n"
	}

	//Initialize carpark
	finish := false
	for !finish {
		fmt.Print("Enter command: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimRight(input, terminal)
		s := parse(input)
		switch s[0] {
		case "create_parking_lot":
			maxSlot, err := strconv.Atoi(s[1])
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			carpark.init(maxSlot)
			fmt.Printf("Created parking lot with %v slots\n", maxSlot)
			finish = true
			
		default: //Default option
			fmt.Println("Your command cannot be understood")
		}
	}

	//Read input queries from console
	finish = false
	for !finish {
		fmt.Print("Enter command: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimRight(input, terminal)
		s := parse(input)
		switch s[0] {
		case "park": //Park a new car
			car := Car{
				registration: s[1],
				colour:       s[2],
			}
			slotNo, err := carpark.insertCar(&car)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("Allocated slot number: %v\n", slotNo)
			}

		case "leave": //Remove a parked car
			slotNo, err := strconv.Atoi(s[1])
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			carpark.removeCar(slotNo)
			fmt.Printf("Slot number %v is free\n", slotNo)

		case "registration_numbers_for_cars_with_colour": //Return registration numbers with given car colour
			_, registration := carpark.getCarsWithColour(s[1])
			prettyPrintStrings(registration)

		case "slot_numbers_for_cars_with_colour": //Return slot numbers with given car colour
			slots, _ := carpark.getCarsWithColour(s[1])
			prettyPrintInts(slots)

		case "slot_number_for_registration_number": //Return slot numbers with given car registration number
			slotNo, err := carpark.getCarWithRegistrationNo(s[1])
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(slotNo)
			}

		case "status":
			fmt.Println("Slot No. Registration No Color")
			for i := 1; i <= carpark.HighestSlot; i++ {
				car, ok := carpark.Map[i]
				if ok {
					fmt.Printf("%v %v %v \n", car.slot, car.registration, car.colour)
				}
			}

		case "done": //End carpark operation
			finish = true

		default: //Default option
			fmt.Println("Your command cannot be understood")
		}
	}
}

func parse(input string) []string {
	s := strings.Split(input, " ")
	return s
}

func prettyPrintStrings(in []string) {
	if len(in) == 0 {
		return
	}
	for i := 0; i < len(in)-1; i++ {
		fmt.Printf("%v, ", in[i])
	}
	fmt.Printf("%v\n", in[len(in)-1])
}

func prettyPrintInts(in []int) {
	if len(in) == 0 {
		return
	}
	for i := 0; i < len(in)-1; i++ {
		fmt.Printf("%v, ", in[i])
	}
	fmt.Printf("%v\n", in[len(in)-1])
}
