package class

import "gorpg/internal/stats"

type Class struct {
	Name       string
	BaseStats  stats.Stats
	GrowthRate stats.Stats
}

func Warrior() Class {
	return Class{
		Name: "Warrior",
		BaseStats: stats.Stats{
			Strength:     10,
			Agility:      7,
			Intelligence: 4,
			Dexterity:    5,
			Luck:         5,
		},
		GrowthRate: stats.Stats{
			Strength:     3,
			Agility:      2,
			Intelligence: 1,
			Dexterity:    2,
			Luck:         1,
		},
	}
}

func Mage() Class {
	return Class{
		Name: "Mage",
		BaseStats: stats.Stats{
			Strength:     2,
			Agility:      4,
			Intelligence: 12,
			Dexterity:    5,
			Luck:         5,
		},
		GrowthRate: stats.Stats{
			Strength:     1,
			Agility:      1,
			Intelligence: 3,
			Dexterity:    2,
			Luck:         1,
		},
	}
}
