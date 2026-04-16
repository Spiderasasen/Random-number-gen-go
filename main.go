package main

import (
	"math/rand"
	"time"
)

func main() {
	//getting the random number
	var randNum int = random_number() //this is a var in go, var <name> <operation>
	println(randNum)
}

// getting a random number
func random_number() int { //must put the return statment at the end of the function line
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10)
	return num
}
