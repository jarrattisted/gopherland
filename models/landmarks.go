package models

import "math/rand"

// Landmark struct to base landmarks off
type Landmark struct {
	Name      string
	Creatures []string
	// Starts off with certain amount of produce
	Huckleberries [2]int
	Guffins       [2]int
	GoteaLeaves   [2]int
	Stones        [2]int
	Plasters      [2]int
}

// LandmarkItems are the determined amounts from GetRandomItems
type LandmarkItems struct {
	Huckleberries int
	Guffins       int
	GoteaLeaves   int
	Stones        int
	Plasters      int
}

// GetRandomItems decides upon amounts of items based on 2 ints in Landmark struct
func (l *Landmark) GetRandomItems() *LandmarkItems {
	i := &LandmarkItems{
		Huckleberries: l.Huckleberries[0] + rand.Intn(l.Huckleberries[1]-l.Huckleberries[0]),
		Guffins:       l.Guffins[0] + rand.Intn(l.Guffins[1]-l.Guffins[0]),
		GoteaLeaves:   l.GoteaLeaves[0] + rand.Intn(l.GoteaLeaves[1]-l.GoteaLeaves[0]),
		Stones:        l.Stones[0] + rand.Intn(l.Stones[1]-l.Stones[0]),
		Plasters:      l.Plasters[0] + rand.Intn(l.Plasters[1]-l.Plasters[0]),
	}
	return i
}
