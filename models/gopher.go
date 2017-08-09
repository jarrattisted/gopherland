package models

import (
	"fmt"
	"math/rand"
)

// Gopher is a type of animal with a SqweakPitch & Spirit Level
type Gopher struct {
	*Animal
	*Backpack
	SqweakPitch string
	Spirit      int
}

// GenerateBackpack creates a new backpack with random values inside
func (g *Gopher) GenerateBackpack() {
	if g.Backpack == nil {
		g.Backpack = &Backpack{}
	}
	b := g.Backpack
	b.Huckleberries += rand.Intn(10)
	b.Guffins += rand.Intn(10)
	b.GoteaLeaves += rand.Intn(10)
	b.Stones += rand.Intn(10)
	b.Plasters += rand.Intn(10)
}

// All the strings that a Gopher returns
func (g *Gopher) String() string {
	s := fmt.Sprintf(`
    🍃 Here's %v's attributes:
      Height: %vcm
      Weight: %vkg
      Legs: %v
      Ears: %v
      Eyes: %v
      Tail: %v
      Sqweak Pitch: %v
      Spirit: %v

    🎒 %v has the following in their backpack:
      %v Huckleberries
      %v Guffins
      %v Gotea Leaves
      %v Stones
      %v Plasters
  `,
		g.Name, g.HeightCm, g.WeightKg, g.Legs, g.Ears, g.Eyes, g.Tail, g.SqweakPitch, g.Spirit,
		g.Name, g.Huckleberries, g.Guffins, g.GoteaLeaves, g.Stones, g.Plasters,
	)
	return s
}

// AttributesString reports a Gopher's current attributes
func (g *Gopher) AttributesString() string {
	s := fmt.Sprintf(`
    🍃 Here's %v's attributes:
      Height: %vcm
      Weight: %vkg
      Legs: %v
      Ears: %v
      Eyes: %v
      Tail: %v
      Sqweak Pitch: %v
      Spirit: %v
  `,
		g.Name, g.HeightCm, g.WeightKg, g.Legs, g.Ears, g.Eyes, g.Tail, g.SqweakPitch, g.Spirit,
	)
	return s
}

// BackpackContentsString reports the inventory of a Gopher's backpack
func (g *Gopher) BackpackContentsString() string {
	s := fmt.Sprintf(`
    🎒 %v has the following in their backpack:
      %v Huckleberries
      %v Guffins
      %v Gotea Leaves
      %v Stones
      %v Plasters
  `,
		g.Name, g.Huckleberries, g.Guffins, g.GoteaLeaves, g.Stones, g.Plasters,
	)
	return s
}

// ExploreLandmark gets items from GetRandomItems, comes up with ints with AddLandmarkItems
// and then adds it to the Gopher's backpack
func (g *Gopher) ExploreLandmark(l *Landmark) {
	li := l.GetRandomItems()
	g.Backpack.AddLandmarkItems(li)
}

// GenerateGopher creates a new Gopher with random values
func GenerateGopher(name string) *Gopher {
	sqweaks := []string{
		"High",
		"Medium",
		"Low",
	}

	g := &Gopher{
		Animal: &Animal{
			Name:     name,
			HeightCm: rand.Intn(100),
			WeightKg: rand.Intn(100),
			Legs:     rand.Intn(4),
			Ears:     rand.Intn(2),
			Eyes:     rand.Intn(2),
			Tail:     true,
		},
		SqweakPitch: sqweaks[rand.Intn(len(sqweaks))],
		Spirit:      10,
	}
	g.GenerateBackpack()
	return g
}
