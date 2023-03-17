package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

func main() {
	client, err := oss.New("", "", "")
	if err != nil {
		// HandleError(err)
	}

	bucket, err := client.Bucket("app-build")
	if err != nil {
		// HandleError(err)
	}

	// 遍历文件。
	marker := oss.Marker("")
	prefix := oss.Prefix("CI/")
	for {
		lor, err := bucket.ListObjects(marker, prefix)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}

		for _, object := range lor.Objects {
			url, err := bucket.SignURL(object.Key, oss.HTTPGet, 600)
			if err != nil {
				fmt.Printf("err:", err)
			}
			fmt.Printf(object.Key, object.StorageClass)
			fmt.Println(url, object.Type, object.Size)
		}
		if lor.IsTruncated {
			prefix = oss.Prefix(lor.Prefix)
			marker = oss.Marker(lor.NextMarker)
		} else {
			break
		}
	}
}
