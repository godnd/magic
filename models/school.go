package models

// School struct for representing a spellschool
type School struct {
	Name string `json:"name"`
}

// Schools is a list of all DND 5e spellschools
var Schools = []School{
	{Name: "necromancy"},
	{Name: "evocation"},
	{Name: "abjuration"},
	{Name: "illusion"},
	{Name: "divination"},
	{Name: "conjuration"},
	{Name: "transmutation"},
	{Name: "enchantment"},
}
