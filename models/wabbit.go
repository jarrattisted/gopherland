package models

import (
	"fmt"
	"math/rand"
)

// Wabbit is a type of animal with a friendly attitude
type Wabbit struct {
	*Animal
	*Creature
}

// GenerateWabbit creates a new Wabbit with default values
func GenerateWabbit() *Wabbit {
	w := &Wabbit{
		Animal: &Animal{
			Name:     GenerateName(),
			HeightCm: rand.Intn(100),
			WeightKg: rand.Intn(100),
			Legs:     4,
			Ears:     2,
			Eyes:     2,
			Tail:     true,
		},
		Creature: &Creature{
			Attitude: GenerateAttitude(),
		},
	}
	return w
}

// All the strings that a Wabbit returns
func (w *Wabbit) String() string {
	s := fmt.Sprintf(`
    🍃 A %v Wabbit named %v appeared. Here's its attributes:
      Height: %vcm
      Weight: %vkg
      Legs: %v
      Ears: %v
      Eyes: %v
      Tail: %v
      Attitude: %v
  `,
		w.Attitude, w.Name, w.HeightCm, w.WeightKg, w.Legs, w.Ears, w.Eyes, w.Tail, w.Attitude,
	)
	return s
}
