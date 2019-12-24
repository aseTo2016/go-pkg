package yaml

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// LoadYamlFromFile loads data from yaml
func LoadYamlFromFile(filePath string, result interface{}) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, result)
}
