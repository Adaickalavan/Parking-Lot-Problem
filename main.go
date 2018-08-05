package main

import (
	"bufio"
	"fmt"
	"os"
	"pretty"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	// Verify if input file or interactive mode to be used.
	// Identify input file name if any.
	// ii := len(os.Args)
	// switch {
	// case ii > 2:
	// 	log.Fatal("Unknown command line input")
	// case ii == 2:
	// 	file, err := os.Open(os.Args[1])
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	defer file.Close()
	// 	scanner := bufio.NewScanner(file)
	// 	for scanner.Scan() {
	// 		fmt.Println(scanner.Text())
	// 		fmt.Println("kios")
	// 	}

	// 	if err := scanner.Err(); err != nil {
	// 		log.Fatal(err)
	// 	}
	// default:
	// 	fmt.Println("hi")
	// }

	// //Create reader input from console
	// reader := bufio.NewReader(os.Stdin)
	// // input, _ := reader.ReadString('\n')
	// // input = strings.TrimRight(input, "\r\n")
	// // s := parse(input)
	// // fmt.Println(s)

	// return

	//Identify operating system used
	var terminal string
	if runtime.GOOS == "windows" {
		terminal = "\r\n"
	} else {
		terminal = "\n"
	}

	reader := bufio.NewReader(os.Stdin)
	//Create a carpark
	var carpark = &Carpark{}

	//Initialize carpark
	exit := false
	for !exit {

		// file, err := os.Open(os.Args[1])
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// scanner := bufio.NewScanner(file)
		// for scanner.Scan() {
		// 	fmt.Println(scanner.Text())
		// 	fmt.Println("kios")
		// }

		// if err := scanner.Err(); err != nil {
		// 	log.Fatal(err)
		// }

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
			exit = true

		default: //Default option
			fmt.Println("Your command cannot be understood")
		}
	}

	//Read input queries from console
	exit = false
	for !exit {
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
			err = carpark.removeCar(slotNo)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("Slot number %v is free\n", slotNo)
			}

		case "registration_numbers_for_cars_with_colour": //Return registration numbers with given car colour
			_, registration := carpark.getCarsWithColour(s[1])
			pretty.Printer(registration)

		case "slot_numbers_for_cars_with_colour": //Return slot numbers with given car colour
			slots, _ := carpark.getCarsWithColour(s[1])
			pretty.Printer(slots)

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

		case "exit": //End carpark operation
			exit = true

			// default: //Default option
			// 	fmt.Println("Your command cannot be understood")
		}
	}
}

func parse(input string) []string {
	s := strings.Split(input, " ")
	return s
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
