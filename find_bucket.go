package divan_data_manager

func FindBucket(name string, buckets []*BucketData) *BucketData {
	for _, bucket := range buckets {
		if bucket.Name == name {
			return bucket
		}
	}

	return nil
}
