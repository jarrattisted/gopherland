package data

import (
	"github.com/jarrattisted/gopherworld/models"
)

// HuckleberryHill is rich with Huckleberries
var HuckleberryHill = &models.Landmark{
	Name: "Huckleberry Hill",
	Creatures: []string{
		"Wabbit",
		"Hucklesucker",
	},
	Huckleberries: [2]int{100, 5000},
	Guffins:       [2]int{0, 30},
	GoteaLeaves:   [2]int{50, 100},
	Stones:        [2]int{0, 20},
	Plasters:      [2]int{0, 5},
}

// PeakPoint is rich with Stones & Plasters
var PeakPoint = &models.Landmark{
	Name: "Peak Point",
	Creatures: []string{
		"Mountalope",
		"Ginoraflopse",
	},
	Huckleberries: [2]int{0, 50},
	Guffins:       [2]int{0, 80},
	GoteaLeaves:   [2]int{0, 20},
	Stones:        [2]int{0, 5000},
	Plasters:      [2]int{10, 500},
}

// LengthyLagoon is rich with Stones, Huckleberries & Guffins
var LengthyLagoon = &models.Landmark{
	Name: "Lengthy Lagoon",
	Creatures: []string{
		"Mountalope",
		"Hucklesuckers",
		"Goonwhale",
	},
	Huckleberries: [2]int{0, 1000},
	Guffins:       [2]int{0, 2000},
	GoteaLeaves:   [2]int{0, 50},
	Stones:        [2]int{0, 3000},
	Plasters:      [2]int{0, 3},
}
