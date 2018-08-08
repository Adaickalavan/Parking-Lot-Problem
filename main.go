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
	ii := len(os.Args)
	var scanner *bufio.Scanner
	switch { //Verify if input file or interactive mode is to be used.
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

//operateCarpark reads input queries from console or text file and executes the command
func operateCarpark(carpark *Carpark, scanner *bufio.Scanner) {
	newlineStr := getNewlineStr()
	exit := false
	for !exit && scanner.Scan() {

		input := scanner.Text()
		input = strings.TrimRight(input, newlineStr)
		s := parse(input)

		switch {
		case s[0] == "create_parking_lot" && len(s) == 2: //Initialize carpark
			maxSlot, err := strconv.Atoi(s[1])
			if checkError(err) {
				break
			}
			err = carpark.init(maxSlot)
			if !checkError(err) {
				fmt.Printf("Created parking lot with %v slots\n", maxSlot)
			}

		case s[0] == "park" && len(s) == 3: //Park a new car
			car := Car{
				registration: s[1],
				colour:       s[2],
			}
			slotNo, err := carpark.insertCar(&car)
			if !checkError(err) {
				fmt.Printf("Allocated slot number: %v\n", slotNo)
			}

		case s[0] == "leave" && len(s) == 2: //Remove a parked car
			slotNo, err := strconv.Atoi(s[1])
			if checkError(err) {
				break
			}
			err = carpark.removeCar(slotNo)
			if !checkError(err) {
				fmt.Printf("Slot number %v is free\n", slotNo)
			}

		case s[0] == "registration_numbers_for_cars_with_colour" && len(s) == 2: //Return registration numbers with given car colour
			_, registration, err := carpark.getCarsWithColour(s[1])
			if checkError(err) {
				break
			}
			err = pretty.Printer(registration)
			if err != nil {
				panic(err.Error())
			}

		case s[0] == "slot_numbers_for_cars_with_colour" && len(s) == 2: //Return slot numbers with given car colour
			slots, _, err := carpark.getCarsWithColour(s[1])
			if checkError(err) {
				break
			}
			err = pretty.Printer(slots)
			if err != nil {
				panic(err.Error())
			}

		case s[0] == "slot_number_for_registration_number" && len(s) == 2: //Return slot numbers with given car registration number
			slotNo, err := carpark.getCarWithRegistrationNo(s[1])
			if !checkError(err) {
				fmt.Println(slotNo)
			}

		case s[0] == "status" && len(s) == 1: //Retrieve cars parked in carpark
			carpark.getStatus()

		case s[0] == "exit" && len(s) == 1: //End carpark operation
			exit = true

		default: //Default option
			fmt.Println("Unknown input command")
		}
	}
}

//getNewlineStr identifies operating system and returns newline character used
func getNewlineStr() string {
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
