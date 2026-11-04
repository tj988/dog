package dog

import (
	"fmt"
)

func Bark() {
	fmt.Println("Woof!")
}

func Barkn(numberOfBarks int) {
	for i := 0; i < numberOfBarks; i++ {
		fmt.Print("Woof! ")
	}
	fmt.Println()
}
