package models

// Spell struct for representing a spell
type Spell struct {
	Name   string `json:"name"`
	School string `json:"school"`
	Ritual bool   `json:"ritual"`
}

// Spells list all the different spells
var Spells = []Spell{}
