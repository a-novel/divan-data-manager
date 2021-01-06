package divan_data_manager

import (
	"encoding/json"
	"github.com/a-novel/divan-data-manager/types"
)

func (bd *BucketData) GetAutocompaction() error {
	if bd.AutoCompactionSettings != nil {
		marshalled, err := json.Marshal(bd.AutoCompactionSettings)
		if err != nil {
			return err
		}

		var output divan_types.AutoCompactionSettings
		if err := json.Unmarshal(marshalled, &output); err != nil {
			return err
		}

		output.Parse()
		bd.AutoCompactionSettingsD = &output
	}

	return nil
}
