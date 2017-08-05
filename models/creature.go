package models

import "math/rand"

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
