package stats

import "fmt"

type Stats struct {
	Strength     int
	Agility      int
	Intelligence int
	Dexterity    int
	Luck         int
}

func (stats *Stats) ApplyGrowth(growth Stats) {
	stats.Strength += growth.Strength
	stats.Agility += growth.Agility
	stats.Intelligence += growth.Intelligence
	stats.Dexterity += growth.Dexterity
	stats.Luck += growth.Luck
}

func (s Stats) String() string {
	return fmt.Sprintf(
		"Strength=%d, Agility=%d, Intelligence=%d, Dexterity=%d, Luck=%d",
		s.Strength, s.Agility, s.Intelligence, s.Dexterity, s.Luck,
	)
}
