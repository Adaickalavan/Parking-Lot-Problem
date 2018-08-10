package main

import (
	"bytes"
	"log"
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	//Save old settings before rewriting settings
	oldArgs := os.Args
	oldInputInteractive := inputInteractive
	oldOutStream := outStream
	defer func() {
		os.Args = oldArgs
		inputInteractive = oldInputInteractive
		outStream = oldOutStream
	}()

	//Setup redirection for interactive inputs
	inputInteractiveFile, err := os.Open("inputInteractive.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { inputInteractiveFile.Close() }()
	inputInteractive = inputInteractiveFile

	//Setup redirection for outputs
	var gotBuf bytes.Buffer
	outStream = &gotBuf

	//Setup expected output
	wantBuf := bytes.NewBufferString(wantOut()).Bytes()

	tests := []struct {
		name string
		args []string
	}{
		{name: "File input",
			args: []string{"cmd", "inputFile.txt"},
		},
		{name: "Interactive input",
			args: []string{"cmd"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args
			main()
			if !bytes.Equal(gotBuf.Bytes(), wantBuf) {
				t.Errorf("main() = %v, want = %v", gotBuf.String(), string(wantBuf))
			}
		})
		gotBuf.Reset()
	}
}

func wantOut() string {
	out := `Created a parking lot with 6 slots
Allocated slot number: 1
Allocated slot number: 2
Allocated slot number: 3
Allocated slot number: 4
Allocated slot number: 5
Allocated slot number: 6
Slot number 4 is free
Slot No.    Registration No    Colour
1           KA-01-HH-1234      White
2           KA-01-HH-9999      White
3           KA-01-BB-0001      Black
5           KA-01-HH-2701      Blue
6           KA-01-HH-3141      Black
Allocated slot number: 4
Sorry, parking lot is full
KA-01-HH-1234, KA-01-HH-9999, KA-01-P-333
1, 2, 4
6
Not found
Not found
Not found
Unknown input command
`
	return out
}
