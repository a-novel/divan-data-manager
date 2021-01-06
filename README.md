# Divan Bucket Manager

Lightweight api to manage bucket information from Go server.

```cgo
go get github.com/a-novel/divan-bucket-manager
```

# Usage

```go
package my_package

import(
	"github.com/a-novel/divan-bucket-manager"
)

func main() {
	clusterUsername := "Administrator"
	clusterPassword := "password"
	clusterURL := "127.0.0.1" // You can also leave it blank if on localhost.
	
	buckets, _ := divan_bucket_manager.GetData(clusterUsername, clusterPassword, clusterURL)
	
	// Find a bucket by name.
	myBucket := divan_bucket_manager.FindBucket("bucket_1", buckets)
	
	if myBucket != nil {
		// Do something.
    } else {
    	// Bucket not found.
    }
}
```

# License

[Licensed under MIT, for A-Novel](https://github.com/a-novel/divan-bucket-manager/blob/master/LICENSE).