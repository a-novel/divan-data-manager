package divan_data_manager

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func GetBucketsData(username, password, url string) ([]*BucketData, error) {
	if url == "" {
		url = "127.0.0.1:8091"
	}

	var output []*BucketData

	err := getData(username, password, fmt.Sprintf("%s/pools/default/buckets", url), &output)

	return output, err
}

func GetClusterData(username, password, url string) (*ClusterData, error) {
	if url == "" {
		url = "127.0.0.1:8091"
	}

	var output *ClusterData

	err := getData(username, password, fmt.Sprintf("%s/pools/default", url), output)

	return output, err
}

func getData(username, password, url string, ptr interface{}) error {
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
