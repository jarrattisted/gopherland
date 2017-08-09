package models

// Backpack can carry multiple items
type Backpack struct {
	Huckleberries int
	Guffins       int
	GoteaLeaves   int
	Stones        int
	Plasters      int
}

// AddLandmarkItems takes values from Landmark and adds them to a Gopher's backpack
func (b *Backpack) AddLandmarkItems(li *LandmarkItems) {
	b.Huckleberries += li.Huckleberries
	b.Guffins += li.Guffins
	b.GoteaLeaves += li.GoteaLeaves
	b.Stones += li.Stones
	b.Plasters += li.Plasters
}
