package divan_types

import (
	"encoding/json"
)

func (bd *BucketData) GetAutocompaction() error {
	if bd.AutoCompactionSettings != nil {
		marshalled, err := json.Marshal(bd.AutoCompactionSettings)
		if err != nil {
			return err
		}

		var output AutoCompactionSettings
		if err := json.Unmarshal(marshalled, &output); err != nil {
			return err
		}

		output.Parse()
		bd.AutoCompactionSettingsD = &output
	}

	return nil
}
