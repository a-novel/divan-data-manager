package divan_data_manager

import "github.com/a-novel/divan-data-manager/types"

func FindBucket(name string, buckets []*divan_types.BucketData) *divan_types.BucketData {
	for _, bucket := range buckets {
		if bucket.Name == name {
			return bucket
		}
	}

	return nil
}
