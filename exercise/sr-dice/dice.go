//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRollDice(sides int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(sides) + 1
}

func additionalInfoRoll(totalRolls int, numDice int) string {
	isEven := totalRolls%2 == 0
	isOdd := totalRolls%2 != 0
	if totalRolls == 2 && numDice == 2 {
		return "Snake eyes"
	}
	if totalRolls == 7 {
		return "Lucky 7"
	}
	if isEven {
		return "Even"
	}
	if isOdd {
		return "Odd"
	}
	return ""
}

func main() {

	var numRolls int = 10
	var numDice int = 2
	var numSides int = 6

	for r := 0; r < numRolls; r++ {
		sumRolls := 0
		fmt.Println("Roll:", r+1)
		for d := 0; d < numDice; d++ {
			roll := generateRollDice(numSides)
			sumRolls += roll
			fmt.Println("Dice:", d+1, "Roll:", roll)
		}
		fmt.Println("Rolled:", sumRolls)
		fmt.Println("Additional Info:", additionalInfoRoll(sumRolls, numDice))
	}
}
