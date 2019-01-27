package builtin

import (
	"io/ioutil"
	"net/http"
	"regexp"

	"gopkg.in/yaml.v2"
)

const DefaultMetafileName = "sobjects.yml"

type MetaFileLoader struct {
	src string
}

func NewMetaFileLoader(src string) *MetaFileLoader {
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
	err = yaml.Unmarshal(body, sobjects)
	return *sobjects, err
}
