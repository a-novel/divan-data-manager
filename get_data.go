package divan_bucket_manager

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func GetData(username, password, url string) ([]*BucketData, error) {
	if url == "" {
		url = "127.0.0.1:8091"
	}

	cmd := exec.Command(
		"sh", "-c",
		fmt.Sprintf(
			"curl -s -u \"%s\":\"%s\" http://%s/pools/default/buckets",
			username, password, url,
		),
	)

	stdout, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var output []*BucketData

	err = json.Unmarshal(stdout, &output)
	if err != nil {
		return nil,err
	}

	return output, nil
}
