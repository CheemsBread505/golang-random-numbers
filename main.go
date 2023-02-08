package main

import (
	//imports "fmt" to print to console
	"fmt"
	//imports "math/random" for the random number
	"math/rand"
	//imports "time" for the loop
	"time"
)

func main() {
	for {
		//get's a random number and puts it into a variable
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(999999999999999999)

		//prints out the variable that holds the random number
		fmt.Println(randNum)
		//loops ever 0 seconds
		time.Sleep(0 * time.Second)
	}
}
