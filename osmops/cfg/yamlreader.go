// Convenience functions to read and validate OSM Ops YAML data.
//
package cfg

import (
	v "github.com/go-ozzo/ozzo-validation"
	"gopkg.in/yaml.v2"
)

func fromBytes(yamlData []byte, out v.Validatable) error {
	if err := yaml.Unmarshal(yamlData, out); err != nil {
		return err
	}
	if err := out.Validate(); err != nil {
		return err
	}
	return nil
}

func readOpsConfig(yamlData []byte) (*OpsConfig, error) {
	out := &OpsConfig{}
	err := fromBytes(yamlData, out)
	return out, err
}

func readOsmConnection(yamlData []byte) (*OsmConnection, error) {
	out := &OsmConnection{}
	err := fromBytes(yamlData, out)
	return out, err
}

func readKduNsAction(yamlData []byte) (*KduNsAction, error) {
	out := &KduNsAction{}
	err := fromBytes(yamlData, out)
	return out, err
}
