package models

import (
	"fmt"
	"math/rand"
)

// Creature is an animal other than a Gopher
type Creature struct {
	*Animal
	Attitude string
}

// GenerateAttitude gives creatures an attitude and determines how they behave
// when confronted with a Gopher
func GenerateAttitude() string {
	attitudes := []string{
		"lovely",   // Gives your Gopher many things and maybe even increases spirit
		"friendly", // Takes and give items with Gopher
		"timid",    // Do nothing, runs away
		"mean",     // Take things from a Gopher, and drain spirit
		"nasty",    // Don't wanna come up against this. Will take lots and drain lots of spirit
	}
	return attitudes[rand.Intn(len(attitudes))]
}

/* List of Creatures */

// == Mountalope ==

// Mountalope is a type of animal with Antlers
type Mountalope struct {
	*Animal
	*Creature
	Antlers string
}

// GenerateMountalope creates a new Mountalope with default values
func GenerateMountalope() *Mountalope {
	antlers := [3]string{
		"Enormous",
		"Medium",
		"Small",
	}

	m := &Mountalope{
		Animal: &Animal{
			Name:     GenerateName(),
			HeightCm: rand.Intn(300),
			WeightKg: rand.Intn(200),
			Legs:     4,
			Ears:     2,
			Eyes:     2,
			Tail:     true,
		},
		Creature: &Creature{
			Attitude: GenerateAttitude(),
		},
		Antlers: antlers[rand.Intn(len(antlers))],
	}
	return m
}

// All the strings that a Mountalope returns
func (m *Mountalope) String() string {
	s := fmt.Sprintf(`
    🍃 A %v Mountalope named %v appeared. Here's its attributes:
      Height: %vcm
      Weight: %vkg
      Legs: %v
      Ears: %v
      Eyes: %v
      Tail: %v
      Attitude: %v
			Antlers: %v
  `,
		m.Attitude, m.Name, m.HeightCm, m.WeightKg, m.Legs, m.Ears, m.Eyes, m.Tail, m.Attitude, m.Antlers,
	)
	return s
}

// == Hucklesucker ==

// Hucklesucker is a type of animal
type Hucklesucker struct {
	*Animal
	*Creature
}

// GenerateHucklesucker creates a new Hucklesucker with default values
func GenerateHucklesucker() *Hucklesucker {
	h := &Hucklesucker{
		Animal: &Animal{
			Name:     GenerateName(),
			HeightCm: rand.Intn(30),
			WeightKg: rand.Intn(500),
			Legs:     8,
			Ears:     2,
			Eyes:     2,
			Tail:     true,
		},
		Creature: &Creature{
			Attitude: GenerateAttitude(),
		},
	}
	return h
}

// All the strings that a Hucklesucker returns
func (h *Hucklesucker) String() string {
	s := fmt.Sprintf(`
    🍃 A %v Hucklesucker named %v appeared. Here's its attributes:
      Height: %vcm
      Weight: %vkg
      Legs: %v
      Ears: %v
      Eyes: %v
      Tail: %v
      Attitude: %v
  `,
		h.Attitude, h.Name, h.HeightCm, h.WeightKg, h.Legs, h.Ears, h.Eyes, h.Tail, h.Attitude,
	)
	return s
}

// == Wabbit ==

// Wabbit is a type of animal
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

// == Ginoraflopse ==

// Ginoraflopse is a type of animal
type Ginoraflopse struct {
	*Animal
	*Creature
	WingspanCm int
	Beak       bool
}

// GenerateGinoraflopse creates a new Ginoraflopse with default values
func GenerateGinoraflopse() *Ginoraflopse {
	wingspan := [3]int{
		300,
		700,
		6000,
	}

	g := &Ginoraflopse{
		Animal: &Animal{
			Name:     GenerateName(),
			HeightCm: rand.Intn(3000),
			WeightKg: rand.Intn(5000),
			Legs:     2,
			Ears:     2,
			Eyes:     2,
			Tail:     true,
		},
		Creature: &Creature{
			Attitude: GenerateAttitude(),
		},
		WingspanCm: wingspan[rand.Intn(len(wingspan))],
		Beak:       true,
	}
	return g
}

// All the strings that a Ginoraflopse returns
func (w *Ginoraflopse) String() string {
	s := fmt.Sprintf(`
    🍃 A %v Ginoraflopse named %v appeared. Here's its attributes:
      Height: %vcm
      Weight: %vkg
      Legs: %v
      Ears: %v
      Eyes: %v
      Tail: %v
      Attitude: %v
			WingspanCm: %v
			Beak: %v
  `,
		w.Attitude, w.Name, w.HeightCm, w.WeightKg, w.Legs, w.Ears, w.Eyes, w.Tail,
		w.Attitude, w.WingspanCm, w.Beak,
	)
	return s
}

// == Goonwhale ==

// Goonwhale is a type of animal
type Goonwhale struct {
	*Animal
	*Creature
	SpurtDistance int
}

// GenerateGoonwhale creates a new Goonwhale with default values
func GenerateGoonwhale() *Goonwhale {

	g := &Goonwhale{
		Animal: &Animal{
			Name:     GenerateName(),
			HeightCm: 1000 + rand.Intn(4000),
			WeightKg: 1000 + rand.Intn(9000),
			Legs:     0,
			Ears:     0,
			Eyes:     2,
			Tail:     true,
		},
		Creature: &Creature{
			Attitude: GenerateAttitude(),
		},
		SpurtDistance: rand.Intn(1000),
	}
	return g
}

// All the strings that a Goonwhale returns
func (w *Goonwhale) String() string {
	s := fmt.Sprintf(`
    🍃 A %v Goonwhale named %v appeared. Here's its attributes:
      Height: %vcm
      Weight: %vkg
      Legs: %v
      Ears: %v
      Eyes: %v
      Tail: %v
      Attitude: %v
			SpurtDistance: %v
  `,
		w.Attitude, w.Name, w.HeightCm, w.WeightKg, w.Legs, w.Ears, w.Eyes, w.Tail,
		w.Attitude, w.SpurtDistance,
	)
	return s
}
