package divan_types

type KeyspaceData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DatastoreID string `json:"datastore_id"`
	NamespaceID string `json:"namespace_id"`
}
