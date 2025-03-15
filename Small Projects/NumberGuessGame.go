package main

import (
	"fmt"
	"math/rand"
	"time"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	var (
		choice int
		guess  int
	)
	Difficulty :=
		`Choose Your Difficulty.
	 1. Easy
	 2. Timer 
	 3. Dynamic
	 4. Hell
	 Enter '0' to Quit`

		//Generate Unique Random Number from 1-100 with the help for time
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100) + 1

	for {
		count := 0
		num := rand.Intn(100) + 1
		fmt.Println(Difficulty)
		fmt.Scan(&choice)

		clearScreen()

		// Quit if 0
		if choice == 0 {
			break
		}

		switch choice {
		case 1:
			fmt.Println("A random number between 1-100 is selected. Place your guesss.")

			for {
				fmt.Scan(&guess)
				if guess == num {
					clearScreen()
					count++
					fmt.Printf("Congratulation on guessing the number %d on tries: %d \n", num, count)
					break
				} else if guess > num {
					fmt.Println("The number is less than the current guess.")
					count++
				} else if guess < num {
					fmt.Println("The number is more than the current guess.")
					count++
				}
			}

		case 2:
			fmt.Println("A random number between 1-100 is selected. Place your guesss within the time limit. ")
			for {

				if time == 0 {
					fmt.Println("Timer has run out. You have lost")
					break
				}
				if guess == num {
					clearScreen()
					count++
					fmt.Printf("Congratulation on guessing the number %d on tries: %d \n", num, count)
					break
				}
			}

		}

	}
	fmt.Println(guess)
	fmt.Println(num)

}
