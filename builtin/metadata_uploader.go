package builtin

import (
	"encoding/json"
	"io/ioutil"
)

func createMetadata() (map[string]Sobject, error) {
	client := NewSoapClient()
	r, err := client.DescribeGlobal()
	if err != nil {
		return nil, err
	}
	sobjects := map[string]Sobject{}
	for _, sobj := range r.Sobjects {
		r, err := client.DescribeSObject(sobj.Name)
		if err != nil {
			return nil, err
		}
		fields := make([]SobjectField, len(r.Fields))
		for i, f := range r.Fields {
			fields[i] = SobjectField{
				Name:             f.Name,
				Label:            f.Label,
				RelationshipName: f.RelationshipName,
				Type:             string(*f.Type_),
				Custom:           f.Custom,
				ReferenceTo:      f.ReferenceTo,
			}
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

func CreateMetadataFile(filename string) error {
	sobjects, err := createMetadata()
	if err != nil {
		return err
	}
	b, err := json.Marshal(sobjects)
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
	err = json.Unmarshal(b, sobjects)
	return *sobjects, err
}
