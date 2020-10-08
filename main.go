package main

import (
	"fmt"
	"math/rand"
	"time"
)

type die struct {
	sides int
	value int
}

func (d *die) roll() int {
	rand.Seed(time.Now().UnixNano())

	min := 1
	d.value = rand.Intn((d.sides - min + 1) + min)

	return d.value
}

func (d *die) getPoints() int {
	switch d.value {
	case 1:
		return 1
	case 2:
		return 10
	case 3:
		return 50
	case 4:
		return 100
	case 5:
		return 250
	case 6:
		return 500
	}
	return 0
}

func (d *die) getFish() string {
	switch d.value {
	case 1:
		return "Old Shoe"
	case 2:
		return "Small Fish"
	case 3:
		return "Big Fish"
	case 4:
		return "Huge Fish"
	case 5:
		return "Small Shark"
	case 6:
		return "Big Shark"
	}
	return "Caught Nothing"
}

func info(points int) {
	fmt.Println("\nYou have gathered", points, "points")
	if points >= 2500 {
		fmt.Println("Congratulations! Great Score!")
	} else if points >= 2000 {
		fmt.Println("You did well")
	} else if points >= 1500 {
		fmt.Println("You did good")
	} else if points >= 1000 {
		fmt.Println("You did ok")
	} else {
		fmt.Println("Could have done better")
	}
}

func main() {
	var choice int
	var points int

	d := die{
		sides: 6,
		value: 0,
	}

	fmt.Println("\nFishing Game")

	//start game ask user to select choices on the screen
	for {
		fmt.Println("  1. Fishing Game")
		fmt.Println("  2. Display points")
		fmt.Println("  3. Quit")
		fmt.Println("\nPick your choice:")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			d.roll()
			points += d.getPoints()
			fmt.Println("\nYou caught a", d.getFish())
			break
		case 2:
			fmt.Println("\nYou have", points, "points")
			break
		case 3:
			info(points)
			return
		}
	}
}
