package main

// import (
// 	"os"
// 	"testing"
// )

// func Test_main(t *testing.T) {
// 	oldArgs := os.Args
// 	defer func() { os.Args = oldArgs }()

// 	//Setup redirection for interactive inputs
// 	inputInteractive, err := os.Open("inputInteractive.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer inputInteractive.Close()

// 	tests := []struct {
// 		name string
// 		args []string
// 	}{
// 		{name: "File input",
// 			args: []string{"cmd", "inputFile.txt"},
// 		},
// 		{name: "Interactive input",
// 			args: []string{"cmd"},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			os.Args = tt.args
// 			main()
// 		})
// 	}
// }
