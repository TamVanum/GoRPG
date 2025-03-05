package character

import (
	"gorpg/internal/class"
	"gorpg/internal/stats"
)

type Character struct {
	Name  string
	Level int
	Stats *stats.Stats
	Class class.Class
}

func NewCharacter(name string, class class.Class) Character {
	return Character{Name: name, Level: 1, Class: class,
		Stats: &stats.Stats{
			Strength:     class.BaseStats.Strength,
			Agility:      class.BaseStats.Agility,
			Intelligence: class.BaseStats.Intelligence,
			Dexterity:    class.BaseStats.Dexterity,
			Luck:         class.BaseStats.Luck,
		}}
}

func (character *Character) LevelUp() {
	character.Level++
	character.Stats.ApplyGrowth(character.Class.GrowthRate)
}

func (c Character) String() string
