package divan_data_manager

import (
	"encoding/json"
	"fmt"
	"github.com/a-novel/divan-data-manager/types"
	"os/exec"
)

func GetBucketsData(username, password, url string) ([]*divan_types.BucketData, error) {
	if url == "" {
		url = "127.0.0.1:8091"
	}

	var output []*divan_types.BucketData

	err := getDataCurl(username, password, fmt.Sprintf("%s/pools/default/buckets", url), &output)

	return output, err
}

func GetClusterData(username, password, url string) (*divan_types.ClusterData, error) {
	if url == "" {
		url = "127.0.0.1:8091"
	}

	var output divan_types.ClusterData

	err := getDataCurl(username, password, fmt.Sprintf("%s/pools/default", url), &output)

	return &output, err
}

func getDataCurl(username, password, url string, ptr interface{}) error {
	cmd := exec.Command(
		"sh", "-c",
		fmt.Sprintf(
			"curl -s -u \"%s\":\"%s\" http://%s",
			username, password, url,
		),
	)

	stdout, err := cmd.Output()
	if err != nil {
		return err
	}

	err = json.Unmarshal(stdout, ptr)
	if err != nil {
		return err
	}

	return nil
}


