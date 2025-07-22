package main

import "fmt"

func main(){
	
	fmt.Println("Welcome to my quiz game!")

	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)
	//fmt.Println(age)

	fmt.Printf("Hello, %v, welcome to the game! \n", name )


	fmt.Printf("Enter your age: ")
	var age int
    fmt.Scan(&age)
	
	if age >= 10 {
		fmt.Println("you you can play")
	} else {
			fmt.Println("You can not play!")
			
			return
		}

		score := 0 
		num_questions := 2
		

		fmt.Printf("what is better, the rtx 3080 or rtx 3090? ")

		var answer string
		var answer2 string
		fmt.Scan(&answer, &answer2)
		fmt.Println(answer, answer2)

		if answer+ " " + answer2 == "rtx 3090" {
			fmt.Println("Correct!")
			score ++
		} else if answer+ " " + answer2 == " RTX 3090" {
			fmt.Println("Correcte")
		} else {
			fmt.Println("Incorrect!")
		}

		fmt.Printf("how meny cores does the ryzen 9 3500x have")
		var cores uint
		fmt.Scan(&cores)

		if cores == 12 {
			fmt.Println("\n Corret!")
			score ++
		} else{
			fmt.Println(" \n Incorrect!")
		}
		fmt.Printf("your scored %v out of %v. ", score, num_questions)

		percent := (float64(score)/ float64(num_questions)) *100
		fmt.Printf("You Scored: %v%%", percent)
	}