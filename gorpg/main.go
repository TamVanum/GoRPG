package main

import (
	"fmt"
	"gorpg/internal/character"
	"gorpg/internal/class"
)

func main() {

	firstOne := character.NewCharacter("firstOne", class.Warrior())
	fmt.Println(firstOne.Stats.Strength)
	firstOne.LevelUp()
	fmt.Println(firstOne.Stats.Strength)
}
