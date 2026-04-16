package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//getting the random number
	var randNum int = random_number() //this is a var in go, var <name> <operation>

	var choice int

	for {
		println("Please enter a random number: ")
		_, err := fmt.Scan(&choice) //_ tells to forget, err is the error handler (no try expect here)

		//checks if the number is a number
		if err != nil {
			fmt.Println("Please enter a number")
			continue
		}

		//checking if the user got the number, is too low or too high
		if choice == randNum {
			fmt.Println("You Win!!!")
			break
		} else if choice < randNum { //this is discusting but ok
			fmt.Println("To low")
		} else { //as for this as well
			fmt.Println("To high")
		}
	}
}

// getting a random number
func random_number() int { //must put the return statment at the end of the function line
	rand.Seed(time.Now().UnixNano()) //seed of the random number (should be in main, but im learning)
	num := rand.Intn(10)
	return num
}
