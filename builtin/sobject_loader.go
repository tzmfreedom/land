package builtin

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
)

const DefaultMetafileName = "sobjects.json"

type MetaFileLoader struct {
	src string
}

func newMetaFileLoader(src string) *MetaFileLoader {
	return &MetaFileLoader{src: src}
}

func (m *MetaFileLoader) Load() (map[string]Sobject, error) {
	match, err := regexp.MatchString("^http(s)?://", m.src)
	if err != nil {
		return nil, err
	}
	var body []byte
	if match {
		r, err := http.Get(m.src)
		if err != nil {
			return nil, err
		}
		defer r.Body.Close()
		body, err = ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
	} else {
		body, err = ioutil.ReadFile(m.src)
		if err != nil {
			return nil, err
		}
	}
	sobjects := &map[string]Sobject{}
	err = json.Unmarshal(body, sobjects)
	return *sobjects, err
}
