package divan_data_manager

import (
	"fmt"
	"runtime"
)

func GetCBQPath() (string, error) {
	switch runtime.GOOS {
	case "darwin":
		return "/Applications/Couchbase\\ Server.app/Contents/Resources/couchbase-core/bin/cbq", nil
	case "linux":
		return "/opt/couchbase/bin/cbq", nil
	default:
		return "", fmt.Errorf("unknown operating system %s", runtime.GOOS)
	}
}
