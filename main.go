package main

import (
	"fmt"

	"github.com/FXLJA/go-coding-challenge/challenge_01"
	"github.com/FXLJA/go-coding-challenge/challenge_02"
	"github.com/FXLJA/go-coding-challenge/challenge_03"
	"github.com/FXLJA/go-coding-challenge/challenge_04"
)

func main() {

	fmt.Print("Hello World! \n\n")

	Challenge01()
	Challenge02()
	Challenge03()
	Challenge04()
}

func Challenge01() {
	fmt.Println("Challenge 01")

	cylinder := challenge_01.Cylinder{}
	fmt.Print("Tinggi Tabung: ")
	fmt.Scan(&cylinder.Height)

	fmt.Print("Radius Tabung: ")
	fmt.Scan(&cylinder.Radius)

	fmt.Printf("Luas Permukaan Tabung: %.2f\n\n", cylinder.SurfaceArea())
}

func Challenge02() {
	fmt.Println("Challenge 02")

	var score int
	fmt.Print("Score: ")
	fmt.Scan(&score)

	fmt.Printf("Grade: %s\n\n", challenge_02.Grade(score))
}

func Challenge03() {
	fmt.Println("Challenge 03")

	var n int
	fmt.Print("Input: ")
	fmt.Scan(&n)
	fmt.Printf("Total: %d\n\n", challenge_03.Sum(n))
}

func Challenge04() {
	fmt.Println("Challenge 04")

	var height int
	fmt.Print("Tinggi segitiga: ")
	fmt.Scan(&height)

	challenge_04.Triangle(height)
	fmt.Println()
}
