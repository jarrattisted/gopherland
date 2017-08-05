package models

// Landmark struct to base landmarks off of
type Landmark struct {
	Name      string
	Creatures []string
	Produce   map[string]int
}

// // ExploreHuckleberryHill adds mostly Huckleberries to a backpack
// func (b *Backpack) ExploreHuckleberryHill() {
// 	b.Huckleberries += rand.Intn(500)
// 	b.Guffins += rand.Intn(10)
// 	b.GoteaLeaves += rand.Intn(10)
// 	b.Stones += rand.Intn(5)
// 	b.Plasters += rand.Intn(2)
// }
//
// // ExplorePeakPoint adds mostly GoteaLeaves to a backpack
// func (b *Backpack) ExplorePeakPoint() {
// 	b.Huckleberries += rand.Intn(10)
// 	b.Guffins += rand.Intn(50)
// 	b.GoteaLeaves += rand.Intn(500)
// 	b.Stones += rand.Intn(50)
// 	b.Plasters += rand.Intn(10)
// }
//
// // ExploreLengthyLagoon adds mostly Guffins to a backpack
// func (b *Backpack) ExploreLengthyLagoon() {
// 	b.Huckleberries += rand.Intn(50)
// 	b.Guffins += rand.Intn(500)
// 	b.GoteaLeaves += rand.Intn(75)
// 	b.Stones += rand.Intn(100)
// 	b.Plasters += rand.Intn(20)
// }
//
// // ExploreChalkCaves adds mostly Stones to a backpack
// func (b *Backpack) ExploreChalkCaves() {
// 	b.Huckleberries += rand.Intn(10)
// 	b.Guffins += rand.Intn(4)
// 	b.GoteaLeaves += rand.Intn(75)
// 	b.Stones += rand.Intn(500)
// 	b.Plasters += rand.Intn(20)
// }
//
// // ExploreSentryStatue adds mostly Guffins to a backpack
// func (b *Backpack) ExploreSentryStatue() {
// 	b.Huckleberries += rand.Intn(75)
// 	b.Guffins += rand.Intn(500)
// 	b.GoteaLeaves += rand.Intn(20)
// 	b.Stones += rand.Intn(4)
// 	b.Plasters += rand.Intn(120)
// }
