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

func FindIndex(name string, indexes []*divan_types.IndexData) *divan_types.IndexData {
	for _, index := range indexes {
		if index.Name == name {
			return index
		}
	}

	return nil
}

func FindNamespace(name string, namespaces []*divan_types.NamespaceData) *divan_types.NamespaceData {
	for _, namespace := range namespaces {
		if namespace.Name == name {
			return namespace
		}
	}

	return nil
}

func FindKeyspace(name string, keyspaces []*divan_types.KeyspaceData) *divan_types.KeyspaceData {
	for _, keyspace := range keyspaces {
		if keyspace.Name == name {
			return keyspace
		}
	}

	return nil
}
