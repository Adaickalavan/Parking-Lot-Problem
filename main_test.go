package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	//Save old settings before rewriting parameters
	oldArgs := os.Args
	oldInputInteractive := inputInteractive
	oldOutput := output
	defer func() {
		os.Args = oldArgs
		inputInteractive = oldInputInteractive
		output = oldOutput
	}()

	//Setup redirection for interactive inputs
	inputInteractive, err := os.Open("inputInteractive.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputInteractive.Close()

	//Create tempFile to save output for comparison
	tmpfile, err := ioutil.TempFile("", "gotOutput")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		tmpfile.Close()
		// os.Remove(tmpfile.Name())
	}()

	//Setup redirection for outputs
	output = tmpfile

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

		})
	}
}
