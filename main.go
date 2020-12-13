package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	v := ""
	if len(os.Args) > 1 {
		for i := 0; i < 45; i++ {
			num := rand.Intn(250)
			space := ""
			for c := 0; c < num; c++ {
				space += " "
			}
			v = fmt.Sprint("\033[38;2;0;200;0m" + os.Args[1] + "\033[0m\n")
			fmt.Print(space, v)
		}

	} else {
		fmt.Printf("Usage is \033[38;2;200;0;0mpatternmaker \033[38;2;0;200;0mmessage\033[0m")
	}

}
