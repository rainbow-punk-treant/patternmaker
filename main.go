package main


import (
	"os"
	"fmt"
	"math/rand"
)

func main() {
	total := ""
	for i := 0;i < 45;i++ {
		spaces := rand.Intn(50)
		for c := 0;c < spaces;c++ {
			total += " "
		}
		v := fmt.Sprint("\033[38;2;0;200;0m"+os.Args[1]+"\n\033[0m")
		total += v
	}		
	fmt.Printf(total)
	

}
