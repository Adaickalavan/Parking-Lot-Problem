package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"pretty"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	// Verify if input file or interactive mode to be used.
	// Identify input file name if any.
	ii := len(os.Args)
	var scanner *bufio.Scanner
	switch {
	case ii > 2:
		log.Fatal("Unknown command line input")
	case ii == 2:
		file, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
	default:
		scanner = bufio.NewScanner(os.Stdin)
	}

	//Create a carpark
	var carpark = &Carpark{}

	//Operate the carpark
	operateCarpark(carpark, scanner)
}

func operateCarpark(carpark *Carpark, scanner *bufio.Scanner) {
	//Read input queries from console or text file
	newlineStr := getNewlineStr()
	exit := false
	for !exit && scanner.Scan() {

		input := scanner.Text()
		input = strings.TrimRight(input, newlineStr)
		s := parse(input)

		switch s[0] {
		case "create_parking_lot": //Initialize carpark
			maxSlot, err := strconv.Atoi(s[1])
			if checkError(err) {
				break
			}
			err = carpark.init(maxSlot)
			if !checkError(err) {
				fmt.Printf("Created parking lot with %v slots\n", maxSlot)
			}

		case "park": //Park a new car
			car := Car{
				registration: s[1],
				colour:       s[2],
			}
			slotNo, err := carpark.insertCar(&car)
			if !checkError(err) {
				fmt.Printf("Allocated slot number: %v\n", slotNo)
			}

		case "leave": //Remove a parked car
			slotNo, err := strconv.Atoi(s[1])
			if checkError(err) {
				break
			}
			err = carpark.removeCar(slotNo)
			if !checkError(err) {
				fmt.Printf("Slot number %v is free\n", slotNo)
			}

		case "registration_numbers_for_cars_with_colour": //Return registration numbers with given car colour
			_, registration, err := carpark.getCarsWithColour(s[1])
			if checkError(err) {
				break
			}
			err = pretty.Printer(registration)
			if err != nil {
				panic(err.Error())
			}

		case "slot_numbers_for_cars_with_colour": //Return slot numbers with given car colour
			slots, _, err := carpark.getCarsWithColour(s[1])
			if checkError(err) {
				break
			}
			err = pretty.Printer(slots)
			if err != nil {
				panic(err.Error())
			}

		case "slot_number_for_registration_number": //Return slot numbers with given car registration number
			slotNo, err := carpark.getCarWithRegistrationNo(s[1])
			if !checkError(err) {
				fmt.Println(slotNo)
			}

		case "status": //Retrieve cars parked in carpark
			carpark.getStatus()

		case "exit": //End carpark operation
			exit = true

		default: //Default option
			fmt.Println("Unknown input command")
		}
	}
}

func getNewlineStr() string {
	//Identify operating system and newline character used
	if runtime.GOOS == "windows" {
		return "\r\n"
	}
	return "\n"
}

func parse(input string) []string {
	s := strings.Split(input, " ")
	return s
}

func checkError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
		return true
	}
	return false
}
