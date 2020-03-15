package converters_test

import (
	"fmt"
	"testing"

	"github.com/inquizarus/godnd/magic/converters"
	"github.com/stretchr/testify/assert"
)

func TestThatKebabStringIsConvertedToRegularCorrectly(t *testing.T) {
	in := "ray-of-frost"
	expected := "ray of frost"
	out := converters.KebabCaseStringToRegular(in)
	assert.Equal(t, expected, out, fmt.Sprintf("Could not convert kebabstring to regular. Got %s but expected %s", out, expected))
}
