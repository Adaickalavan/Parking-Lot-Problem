package pretty

import "fmt"

type Printer interface {
	print()
}

func PrintStrings(in []string) {
	if len(in) == 0 {
		return
	}
	for i := 0; i < len(in)-1; i++ {
		fmt.Printf("%v, ", in[i])
	}
	fmt.Printf("%v\n", in[len(in)-1])
}

func PrintInts(in []int) {
	if len(in) == 0 {
		return
	}
	for i := 0; i < len(in)-1; i++ {
		fmt.Printf("%v, ", in[i])
	}
	fmt.Printf("%v\n", in[len(in)-1])
}
