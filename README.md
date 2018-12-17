# Parking Lot Problem

See [website](https://adaickalavan.github.io/portfolio/parking_lot_problem/) for information.

## Problem Statement

Mr Zorro owns a multi-storey parking lot that can hold up to `n` vehicles at any given point in time. The parking slots are numbered, beginning at 1 and increases with increasing distance from the entry point in steps of one. Mr Zorro has requested your help to design an automated ticketing system for his parking lot.

When a vehicle enters the parking lot, its vehicle registration number (i.e., number plate) and colour are noted. Then, an available parking slot is allocated. Following are the rules of parking slot ticket issuance:

+ Each customer should be allocated the nearest available parking slot to the entry point.
+ Upon exiting the parking lot, the customer returns the ticket which marks their previously allocated lot as now available.
+ Due to government regulation, the system should provide the ability to determine:
  + Registration numbers of all cars of a particular colour.
  + Slot number in which a car with a given registration number is parked.
  + Slot numbers of all slots where a car of a particular colour is parked.

The ticketing system should be operable via two modes of input, namely, interactive commands and commands from a file. In other words, the ticketing system should be an executable which accepts:

1. Interactive commands from an interactive command prompt shell
2. A filename as an input argument at the command prompt and executes the commands from the given file

Example below includes all the commands which need to be supported.

**Example: File Input**

To run the code so it accepts input from a file:
```
$ bin/parking_lot file_inputs.txt
```
Input (contents of file):
```
create_parking_lot 6
park KA-01-HH-1234 White
park KA-01-HH-9999 White
park KA-01-BB-0001 Black
park KA-01-HH-7777 Red
park KA-01-HH-2701 Blue
park KA-01-HH-3141 Black
leave 4
status
park KA-01-P-333 White
park DL-12-AA-9999 White
registration_numbers_for_cars_with_colour White
slot_numbers_for_cars_with_colour White
slot_number_for_registration_number KA-01-HH-3141
slot_number_for_registration_number MH-04-AY-
```
Output (to STDOUT):
```
Created a parking lot with 6 slots
Allocated slot number: 1
Allocated slot number: 2
Allocated slot number: 3
Allocated slot number: 4
Allocated slot number: 5
Allocated slot number: 6
Slot number 4 is free
Slot No. Registration No Colour
1 KA-01-HH-1234 White
2 KA-01-HH-9999 White
3 KA-01-BB-0001 Black
5 KA-01-HH-2701 Blue
6 KA-01-HH-3141 Black
Allocated slot number: 4
Sorry, parking lot is full
KA-01-HH-1234, KA-01-HH-9999, KA-01-P-333
1, 2, 4
6
Not found
```

**Example: Interactive**

To run the code, launch the shell, and to accept interactive input from the shell:
```
$ bin/parking_lot
```
Assuming a parking lot with `n=6` slots, the following commands should be run in sequence by typing them in at a prompt and should produce output as described below the command. Note that `exit` terminates the process and returns control to the shell.
```
$ create_parking_lot 6
Created a parking lot with 6 slots
$ park KA-01-HH-1234 White
Allocated slot number: 1
$ park KA-01-HH-9999 White
Allocated slot number: 2
$ park KA-01-BB-0001 Black
Allocated slot number: 3
$ park KA-01-HH-7777 Red
Allocated slot number: 4
$ park KA-01-HH-2701 Blue
Allocated slot number: 5
$ park KA-01-HH-3141 Black
Allocated slot number: 6
$ leave 4
Slot number 4 is free
$ status
Slot No. Registration  No Colour
1        KA-01-HH-1234 White
2        KA-01-HH-9999 White
3        KA-01-BB-0001 Black
5        KA-01-HH-2701 Blue
6        KA-01-HH-3141 Black
$ park KA-01-P-333 White
Allocated slot number: 4
$ park DL-12-AA-9999 White
Sorry, parking lot is full
$ registration_numbers_for_cars_with_colour White
KA-01-HH-1234, KA-01-HH-9999, KA-01-P-333
$ slot_numbers_for_cars_with_colour White
1, 2, 4
$ slot_number_for_registration_number KA-01-HH-3141
6
$ slot_number_for_registration_number MH-04-AY-1111
Not found
$ exit
```

## Learning Outcome

At the end of this project, we should be able to:

+ solve the Parking Lot problem with data structures optimized for complexity
+ write comprehensive unit tests in Golang
+ perform functional testing in Golang
+ utilize heap, hash map, and slice, data structures in Go
+ pretty print slices, arrays, and strings, using an interface

## Instructions

1. **Setup Go**
    + Install Go following the instructions [here](https://golang.org/dl/).
    + Set `GOROOT` which is the location of your Go installation. Assuming it is installed at `$HOME/go2.x`, execute:
        ```
        export GOROOT=$HOME/go2.x
        export PATH=$PATH:$GOROOT/bin
        ```
    + Set `GOPATH` environment variable which specifies the location of your Go workspace. It defaults to `$HOME/go` on Unix/Linux.
        ```
        export GOPATH=$HOME/go
        ```
    + Set `GOBIN` path for generation of binary file when `go install` is run.
        ```
        export GOBIN=$GOPATH/bin
        export PATH=$PATH:$GOPATH/bin
        ```
2. **Source code**
    + Git clone the entire project folder into `$GOPATH/src/parking_lot` folder in your computer
        ```
        git clone https://github.com/Adaickalavan/Parking-Lot-Problem.git
        ```
3. **Executable**
    + To create an executable in the `$GOPATH/bin/` directory, execute
        ```
        go install parking_lot
        ```
4. **Unit test and functional test**
    + To run complete test suite, run
        ```
        go test -v parking_lot
        ```
        Here, `-v` is the verbose command flag.
    + To run specific test, run
        ```
        go test -v parking_lot -run xxx
        ```
        Here, `xxx` is the name of test function.
    + Test coverage: 94.1% of statements
5. **Running**
    + Launch interactive user input mode by executing
        ```
        $GOPATH/bin/parking_lot
        ```
    + Launch file input mode by executing
        ```
        $GOPATH/bin/parking_lot.exe $GOPATH/src/parking_lot/inputFile.txt
        ```
        Here, `$GOPATH/src/parking_lot/inputFile.txt` refers to the input file with complete path.

## Project structure

The project structure is as follows:

```txt
project                               # folder containing all project files
├── bin                               # contains executable commands
│   ├── setup                         # contains commands to build/compile the Go code
│   └── parking_lot.exe               # .exe file for parking_lot generated by `go install` command
└── src                               # contains Go source files
    └── parking_lot                   # main folder
        ├── vendor                    # folder containing dependencies
        │   ├── minheap               # dependant package `minheap`  
        │   │   ├── item.go           # element of heap
        │   │   └── priorityQueue.go  # min heap implementation
        │   └── pretty                # dependant package `pretty`  
        │       └── printer.go        # pretty prints array, slice, string
        ├── car.go                    # element of carpark
        ├── carpark.go                # carpark struct and pointer receiver methods
        ├── carpark_test.go           # unit tests of the carpark.go code
        ├── main.go                   # main file of Go code
        ├── main_test.go              # functional test of the main code
        ├── inputFile.txt             # sample input file for testing
        └── inputInteractive.txt      # sample interactive input for testing
```

## Notes on solution

1. **Data structures**
   + A hash map and a min heap was used to solve the parking lot problem.

2. **Complexity**
    + To park a car: O(log(n1)). Here, n1 is the size of the min heap.
    + To remove a car: O(log(n1)). Here, n1 is the size of the min heap.
    + To retrieve a car by colour: O(n2). Here, n2 is the size of the hash map.
    + To retrieve a car by registration number: O(n2). Here, n2 is the size of the hash map.
    + To get status: O(n2). Here, n2 is the size of the hash map.

3. **Assumptions and rationale for choice of data structures to optimize complexity**
    + Car parking and removing operations will be more frequent compared to retrieving car by colour/registration number or status requests.
    + A hash map with `slot number` as `key` is used to store all the cars parked in the carpark. Complexity O(1) of hash map simplifies insertion and removal of cars by slot number.
    + A min heap is used to store *previoulsy-occupied-but-now-empty* slots in ordered sequence with complexity O(log(n1)) for push and pop operations. Here, *empty slots n1 refer only to slots which were previously occupied but is now free*. It does not refer to the total number of free slots in the carpark.

4. **Alternative solutions to reduce complexity at the expense of increased memory**
    + To achieve complexity O(1) in retrieving a car by colour, implement an additional hash map with `colour` as `key` to store all the cars parked in the carpark.
    + To achieve complexity O(1) in retrieving a car by registration number, implement an additional hash map with `registration number` as `key` to store all the cars parked in the carpark.
