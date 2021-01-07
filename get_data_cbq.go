package divan_data_manager

import (
	"encoding/json"
	"fmt"
	divan_types "github.com/a-novel/divan-data-manager/types"
	"os/exec"
)

func GetNamespacesData(username, password, url string) ([]*divan_types.NamespaceData, error) {
	if url == "" {
		url = "127.0.0.1:8091"
	}

	var output struct {
		Results []*divan_types.NamespaceData `json:"results"`
	}

	if err := getDataCBQ(username, password, url, "system:namespaces", &output); err != nil {
		return nil, err
	} else {
		return output.Results, nil
	}
}

func GetKeyspacesData(username, password, url string) ([]*divan_types.KeyspaceData, error) {
	if url == "" {
		url = "127.0.0.1:8091"
	}

	var output struct {
		Results []*divan_types.KeyspaceData `json:"results"`
	}

	if err := getDataCBQ(username, password, url, "system:keyspaces", &output); err != nil {
		return nil, err
	} else {
		return output.Results, nil
	}
}

func GetIndexesData(username, password, url string) ([]*divan_types.IndexData, error) {
	if url == "" {
		url = "127.0.0.1:8091"
	}

	var output struct {
		Results []*divan_types.IndexData `json:"results"`
	}

	if err := getDataCBQ(username, password, url, "system:indexes", &output); err != nil {
		return nil, err
	} else {
		return output.Results, nil
	}
}

func getDataCBQ(username, password, url, selector string, ptr interface{}) error {
	cliPath, err := GetCBQPath()
	if err != nil {
		return err
	}

	cmd := exec.Command(
		"sh", "-c",
		fmt.Sprintf(
			"%s -engine=\"http://%s:%s@%s\" --script=\"SELECT * FROM %s\"",
			cliPath, username, password, url, selector,
		),
	)

	stdout, err := cmd.Output()
	if err != nil {
		return err
	}

	err = json.Unmarshal(stdout, ptr)
	if err != nil {
		return fmt.Errorf("%s\noriginal response : %s", err.Error(), stdout)
	}

	return nil
}
