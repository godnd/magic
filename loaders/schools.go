package loaders

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/inquizarus/godnd/magic/models"
)

// LoadSchools retrieves schools from a reader and parse it into
// and a slice of schools
func LoadSchools(r io.Reader) ([]models.School, error) {
	var schools []models.School
	var err error
	bbytes, err := ioutil.ReadAll(r)
	if nil != err {
		err = fmt.Errorf("LoadSchools %v", err)
	}
	if nil == err {
		err = json.Unmarshal(bbytes, &schools)
		if err != nil {
			err = fmt.Errorf("LoadSchools %v", err)
		}
	}
	return schools, err
}
