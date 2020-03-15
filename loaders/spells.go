package loaders

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/inquizarus/godnd/magic/models"
)

// LoadSpells retrieves spells from a reader and parse it into
// and a slice of spells
func LoadSpells(r io.Reader) ([]models.Spell, error) {
	var spells []models.Spell
	var err error
	bbytes, err := ioutil.ReadAll(r)
	if nil != err {
		err = fmt.Errorf("LoadSpells %v", err)
	}
	if nil == err {
		err = json.Unmarshal(bbytes, &spells)
		if err != nil {
			err = fmt.Errorf("LoadSpells %v", err)
		}
	}
	return spells, err
}
