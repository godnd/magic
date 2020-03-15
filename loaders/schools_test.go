package loaders_test

import (
	"bytes"
	"testing"

	"github.com/inquizarus/godnd/magic/loaders"
	"github.com/stretchr/testify/assert"
)

func TestThatLoadingSchoolsWorksCorrectly(t *testing.T) {
	r := bytes.NewReader([]byte(`[{"name":"abjuration"}]`))
	schools, err := loaders.LoadSchools(r)
	assert.Nil(t, err)
	assert.NotEmpty(t, schools)
	spell := schools[0]
	assert.Equal(t, "abjuration", spell.Name)
}

func TestThatLoadingSchoolsReturnsErrorwhenJSONIsMalformed(t *testing.T) {
	r := bytes.NewReader([]byte(`[{"name""abjuration"}]`))
	_, err := loaders.LoadSchools(r)
	assert.NotNil(t, err)
}

func TestThatLoadingSchoolsReturnsErrorwhenJSONHaveBadStructure(t *testing.T) {
	r := bytes.NewReader([]byte(`{"spells":[{"name":"abjuration"}]}`))
	_, err := loaders.LoadSchools(r)
	assert.NotNil(t, err)
}
