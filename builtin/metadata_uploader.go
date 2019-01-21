package builtin

import (
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

func createMetadata(username, password, endpoint string, standardObjects []string) (map[string]Sobject, error) {
	client := NewSoapClient(username, password, endpoint)
	r, err := client.DescribeGlobal()
	if err != nil {
		return nil, err
	}
	sobjects := map[string]Sobject{}
	for _, sobj := range r.Sobjects {
		if !sobj.Custom && !contains(sobj.Name, standardObjects) {
			continue
		}
		r, err := client.DescribeSObject(sobj.Name)
		if err != nil {
			return nil, err
		}
		fields := []SobjectField{}
		for _, f := range r.Fields {
			if string(*f.Type_) == "address" {
				continue
			}
			fields = append(
				fields,
				SobjectField{
					Name:             f.Name,
					Label:            f.Label,
					RelationshipName: f.RelationshipName,
					Type:             string(*f.Type_),
					Custom:           f.Custom,
					ReferenceTo:      f.ReferenceTo,
				},
			)
		}
		sobjects[sobj.Name] = Sobject{
			Name:          sobj.Name,
			Custom:        sobj.Custom,
			CustomSetting: sobj.CustomSetting,
			Label:         sobj.Label,
			Fields:        fields,
		}
	}
	return sobjects, nil
}

func contains(name string, array []string) bool {
	for _, elem := range array {
		if strings.ToLower(elem) == strings.ToLower(name) {
			return true
		}
	}
	return false
}

func CreateMetadataFile(username, password, endpoint, filename string, standardObjects []string) error {
	sobjects, err := createMetadata(username, password, endpoint, standardObjects)
	if err != nil {
		return err
	}
	b, err := yaml.Marshal(sobjects)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, b, 0644)
}

func readMetadataFile(filename string) (map[string]Sobject, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	sobjects := &map[string]Sobject{}
	err = yaml.Unmarshal(b, sobjects)
	return *sobjects, err
}
