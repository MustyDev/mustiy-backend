package model

type DanaStore interface {
	All() []Dana
	Save(*Dana) error
	Find(int) *Dana
	Found(int) []Dana
	Update(*Dana) error
	Status(*Dana) error
	Delete(dana *Dana) error
}
