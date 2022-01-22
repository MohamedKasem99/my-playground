package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var ()

func get_answer(answer_p *int) {
	fmt.Print("answer: ")
	fmt.Scanf("%d", answer_p)
	fmt.Print("\n")
}
func get_int_env(var_name string) (val int) {
	str_val := os.Getenv(var_name)
	val, err := strconv.Atoi(str_val)
	if err != nil {
		fmt.Printf("Couldn't cast value of %s \"%s\" to int\n", var_name, str_val)
		os.Exit(1)
	}
	return
}
func get_range() (int, int) {
	min := get_int_env("MIN")
	max := get_int_env("MAX")
	return min, max
}

func main() {
	// generate random int between min, max
	var answer int

	rand.Seed(time.Now().Unix())
	min, max := get_range()
	my_guess := rand.Intn(max-min) + min

	fmt.Printf(`Welcome to guess maker

I have chosen a number between %d and %d.
What is it?
`, min, max)

	get_answer(&answer)
	if my_guess != answer {
		fmt.Printf("%d is the wrong answer try again\n", answer)
		for {
			get_answer(&answer)
			if my_guess == answer {
				break
			}
			fmt.Printf("%d is the wrong answer try again\n", answer)
		}
	}
	fmt.Printf(`You guessed correctly. It was %d 
Bye!
`, my_guess)

}
