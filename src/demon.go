package main

import (
	"gopkg.in/mgo.v2/bson"
)

// Demon is Schema of the Demon Collection
type Demon struct {
	Data string
}

// CreateDemon Create new Demon
func CreateDemon(demon Demon) error {
	return Create("demons", demon)
}

// FindOneDemon for getting Single Demon based on condition
func FindOneDemon(condition bson.M) (Demon, error) {
	return FindOne("demons", condition)
}
