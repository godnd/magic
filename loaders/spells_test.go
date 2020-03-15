package loaders_test

import (
	"bytes"
	"testing"

	"github.com/inquizarus/godnd/magic/loaders"
	"github.com/stretchr/testify/assert"
)

func TestThatLoadingSpellsWorksCorrectly(t *testing.T) {
	r := bytes.NewReader([]byte(`[{"name":"ray of frost"}]`))
	spells, err := loaders.LoadSpells(r)
	assert.Nil(t, err)
	assert.NotEmpty(t, spells)
	spell := spells[0]
	assert.Equal(t, "ray of frost", spell.Name)
}

func TestThatLoadingSpellsReturnsErrorwhenJSONIsMalformed(t *testing.T) {
	r := bytes.NewReader([]byte(`[{"name""ray of frost"}]`))
	_, err := loaders.LoadSpells(r)
	assert.NotNil(t, err)
}

func TestThatLoadingSpellsReturnsErrorwhenJSONHaveBadStructure(t *testing.T) {
	r := bytes.NewReader([]byte(`{"spells":[{"name":"ray of frost"}]}`))
	_, err := loaders.LoadSpells(r)
	assert.NotNil(t, err)
}
