package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	//getting a random number
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	randNum := rand.Intn(10) + 1

	fmt.Println("The computer chose a number from 1 to 10...")
	fmt.Println("Number selected...")

	fmt.Println(randNum) // delete after the first check

	//creating a variable outside the loop
	victory := false

	//in the loop, we repeat part of the code
	for tryUser := 0; tryUser < 3; tryUser++ {
		// getting data from the user
		fmt.Println("You have", 3-tryUser, "attempts")
		fmt.Print("Enter your number:")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)

		}
		// removing characters from a string
		input = strings.TrimSpace(input)

		// converting a string to a number
		inputAfter, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if inputAfter > randNum {
			fmt.Println("Your value is more than selected...")

		} else if inputAfter < randNum {
			fmt.Println("Your value is less than selected...")
		} else {
			victory = true
			fmt.Println("+++++++++++++++++++++...")
			fmt.Println("Well done, you won!!!...")
			fmt.Println("+++++++++++++++++++++...")
			break
		}

	}
	if !victory {
		fmt.Println("====================================...")
		fmt.Println("You have lost, the attempts are over...")
		fmt.Println("The selected number was: ", randNum)
		fmt.Println("====================================...")
	}

}
